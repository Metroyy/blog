package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/russross/blackfriday"
	"gopkg.in/yaml.v2"
)

// Config 结构体定义了网站的基本信息
type Config struct {
	Home_Title             string        `yaml:"home_title"`
	Home_Author            string        `yaml:"home_author"`
	Home_Author_img        string        `yaml:"home_author_img"`
	Home_Self_introduction string        `yaml:"home_self_introduction"`
	Title                  string        `yaml:"title"`
	Img                    string        `yaml:"img"`
	Desc                   string        `yaml:"desc"`
	Text                   template.HTML `yaml:"text"`
	Time                   string        `yaml:"time"`
	Tags                   []string      `yaml:"tags"`
	MDPath                 string
}

func main() {
	//计算md文件数量
	mdCount := CountMarkdownFiles()
	// 读取配置文件
	var configs []Config
	// var archive_mdinfo []Archive_MDInfo
	config := ReadConfig("config.yaml")

	files, err := filepath.Glob("sources/post/*.md")
	if err != nil {
		log.Fatalf("读取 Markdown 文件失败了: %v", err)
	}

	// ExtractMarkdown 从 Markdown 文件中提取内容并更新 Config 配置
	for i := 0; i < len(files); i++ {
		configs = append(configs, ExtractMarkdown(files[i], config, mdCount))
	}

	//处理年份，得到所有md年份和唯一年份
	yearsList, uniqueYears, err := ExtArchiveTime(configs)
	if err != nil {
		log.Fatalf("获取md文档年份失败了: %v", err)
	}

	//拼接home页的跳转路径
	for i := 0; i < len(configs); i++ {
		configs[i].MDPath = strconv.Itoa(yearsList[i]) + "/" + configs[i].MDPath + ".html"
	}

	//将数据按照时间顺序排序
	sortConfigs(configs)

	//计算标签,数量
	uniqueTags, tagsname := CountTags(configs)

	//读取模板html文件
	tmpls := template.Must(template.ParseGlob("sources/templates/*.html"))
	uniquefiles := make([]string, mdCount)
	//提取出md文件名，存入files数组中
	for i := 0; i < mdCount; i++ {
		uniquefiles[i] = ExtMakedownName(files[i])
	}

	//为模板传入的数据赋值
	data := struct {
		ConfigDict     []Config
		MdCount        int
		TagsCount      int
		TagNames       []string
		TagsInfo       map[string]int
		Archive_Year   []int
		Archive_MDInfo [][]string
	}{
		ConfigDict: configs,
		// 统计 sources/articles 文件夹下的 Markdown 文件数量
		MdCount:        mdCount,
		TagsCount:      len(tagsname),
		TagNames:       tagsname,
		TagsInfo:       uniqueTags,
		Archive_Year:   uniqueYears,
		Archive_MDInfo: ExtArcInfo(uniqueYears),
	}
	// fmt.Println(data.Archive_MDInfo)
	//检查年份文件夹，如果没有则创建各个年份的文件夹
	Mkdir(data.Archive_Year)
	// , data.TagsInfo

	// 生成 HTML 文件
	CreateHTML(tmpls, data)
	CreateMdHTML(tmpls, data, uniquefiles, yearsList)

	// //搜索路由
	http.HandleFunc("/Search", func(w http.ResponseWriter, r *http.Request) {
		// 在匿名函数中调用Search函数并传递data结构体的值
		Search(w, r, data.ConfigDict)
	})
	http.HandleFunc("/Gentags", func(w http.ResponseWriter, r *http.Request) {
		// 在匿名函数中调用Search函数并传递data结构体的值
		Gentags(w, r, data, tmpls, uniqueTags)
	})
	log.Fatal(http.ListenAndServe(":4005", nil))
}

// 读取配置文件
func ReadConfig(path string) Config {
	// 读取配置文件
	configFile, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}

	// 解析配置文件
	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}
	return config
}

// ExtractMarkdown 从 Markdown 文件中提取内容并更新 Config 配置
func ExtractMarkdown(mdpath string, config Config, mdCount int) Config {
	// 读取 Markdown 文件
	mdFile, err := os.ReadFile(mdpath)
	if err != nil {
		log.Fatalf("读取 Markdown 文件失败: %v", err)
	}

	// 处理 mdFile 中的内容
	// 解析 Markdown 文件中的关键词
	// 存储从 Markdown 文件中解析的 YAML 前置内容
	mdConfig := make(map[string]string)
	err = yaml.Unmarshal(mdFile, mdConfig)
	if err != nil {
		log.Fatalf("解析 Markdown 文件中的关键词失败: %v", err)
	}

	//处理路径，提取md文件名
	files := ExtMakedownName(mdpath)
	config.MDPath = files

	configs := []struct {
		Field string
		Value string
	}{
		{"title", config.Title},
		{"img", config.Img},
		{"desc", config.Desc},
		{"time", config.Time},
		{"tags", strings.Join(config.Tags, ",")},
	}

	for i := range configs {
		if v, ok := mdConfig[configs[i].Field]; ok {
			if configs[i].Field == "title" {
				config.Title = v
			} else if configs[i].Field == "img" {
				config.Img = v
			} else if configs[i].Field == "desc" {
				config.Desc = v
			} else if configs[i].Field == "time" {
				config.Time = v
			} else if configs[i].Field == "tags" {
				config.Tags = strings.Split(strings.TrimSpace(v), ",")
			} else {
				configs[i].Value = v
			}
		}
	}

	// 获取 Markdown 文件中除头部文件外的内容
	mdContent := string(mdFile)
	start := strings.Index(mdContent, "---") // 找到第一个头部分隔符
	if start != -1 {
		end := strings.Index(mdContent[start+3:], "---") // 找到第二个头部分隔符
		if end != -1 {
			// 从第二个头部分隔符的位置开始截取
			mdContent = mdContent[start+end+6:]
		}
	}

	// 去除头部和尾部的空白字符
	mdContent = strings.TrimSpace(mdContent)

	// 将 Markdown 转换为 HTML 格式
	htmlContent := string(blackfriday.MarkdownCommon([]byte(mdContent)))

	// 将处理后的 HTML 正文赋值给 config
	config.Text = template.HTML(htmlContent)

	return config
}

// 统计 sources/articles 文件夹下的 Markdown 文件数量
func CountMarkdownFiles() int {
	files, err := filepath.Glob("sources/post/*.md")
	if err != nil {
		log.Fatalf("读取 Markdown 文件失败了: %v", err)
	}
	return len(files)
}

// 统计md文档的标签数量，循环所有标签
func CountTags(configs []Config) (uniqueTags map[string]int, tagname []string) {
	uniqueTags = make(map[string]int)
	for i := 0; i < len(configs); i++ {
		for _, tag := range configs[i].Tags {
			tag = strings.TrimSpace(tag)
			if _, ok := uniqueTags[tag]; !ok {
				uniqueTags[tag] = 1
				tagname = append(tagname, tag)
			} else {
				uniqueTags[tag]++
			}
		}
	}
	return uniqueTags, tagname
}

// 生成 HTML 文件,home,archive,tags
func CreateHTML(tmpls *template.Template, data struct {
	ConfigDict     []Config
	MdCount        int
	TagsCount      int
	TagNames       []string
	TagsInfo       map[string]int
	Archive_Year   []int
	Archive_MDInfo [][]string
}) {

	// //创建标签对应的HTML文件
	// for i := 0; i < len(data.TagName); i++ {
	// 	out_tag, err := os.Create("sources/articles/tagpage/" + data.TagName[i] + "/" + data.TagName[i] + "_tagpage.html")
	// 	if err != nil {
	// 		fmt.Printf("创建%s输出文件失败: %v\n", uniquefiles[i], err)
	// 	}
	// 	errs := tmpls.ExecuteTemplate(out_tag, "tag.html", data.TagName[i])
	// 	if errs != nil {
	// 		log.Fatalf("替换"+data.TagName[i]+"标签模板中的占位符失败: %v", errs)
	// 	}
	// }

	// 创建输出文件并将模板引擎替换后的结果
	out_home, err := os.Create("sources/articles/home.html")
	if err != nil {
		log.Fatalf("创建home输出文件失败: %v", err)
	}

	out_archive, err := os.Create("sources/articles/archive.html")
	if err != nil {
		log.Fatalf("创建archive输出文件失败: %v", err)
	}

	out_tags, err := os.Create("sources/articles/tags.html")
	if err != nil {
		log.Fatalf("创建tags输出文件失败: %v", err)
	}

	// out_tag, err := os.Create("sources/articles/tag.html")
	// if err != nil {
	// 	log.Fatalf("创建tag输出文件失败: %v", err)
	// }

	errs := tmpls.ExecuteTemplate(out_home, "home.html", data)
	if errs != nil {
		log.Fatalf("替换home模板中的占位符失败: %v", errs)
	}

	err = tmpls.ExecuteTemplate(out_archive, "archive.html", data)
	if err != nil {
		log.Fatalf("替换archive模板中的占位符失败: %v", err)
	}

	err = tmpls.ExecuteTemplate(out_tags, "tags.html", data)
	if err != nil {
		log.Fatalf("替换tags模板中的占位符失败: %v", err)
	}

	// err = tmpls.ExecuteTemplate(out_tag, "tag.html", data)
	// if err != nil {
	// 	log.Fatalf("替换tag模板中的占位符失败: %v", err)
	// }
}

// 创建md对应的HTML文件
func CreateMdHTML(tmpls *template.Template, data struct {
	ConfigDict     []Config
	MdCount        int
	TagsCount      int
	TagNames       []string
	TagsInfo       map[string]int
	Archive_Year   []int
	Archive_MDInfo [][]string
}, uniquefiles []string, yearsList []int) {

	for i := 0; i < data.MdCount; i++ {
		out_md, err := os.Create("sources/articles/" + strconv.Itoa(yearsList[i]) + "/" + uniquefiles[i] + ".html")
		if err != nil {
			log.Fatalf("创建"+uniquefiles[i]+"输出文件失败: %v", err)
		}
		errs := tmpls.ExecuteTemplate(out_md, "index.html", data.ConfigDict[i])
		if errs != nil {
			log.Fatalf("替换"+uniquefiles[i]+"模板中的占位符失败: %v", errs)
		}
	}
	fmt.Println("执行完成")
}

// 处理路径，提取md文件名
func ExtMakedownName(files string) string {
	files = strings.TrimPrefix(strings.TrimSuffix(files, ".md"), "sources/post/")
	return files
}

// 处理md文档中的time字段，提取年，取出所有年和唯一年
func ExtArchiveTime(configs []Config) (yearsList []int, uniqueYears []int, err error) {

	years := make(map[int]struct{})
	if len(configs) == 0 {
		return nil, nil, nil
	}
	yearsList = make([]int, len(configs))
	i := 0
	for _, config := range configs {
		if len(config.Time) >= 4 {
			//获取年份
			year, err := strconv.Atoi(config.Time[:4])
			if err != nil {
				log.Fatalf("获取年份失败: %v", err)
			}
			if year > 0 {
				//判断map中有没有这个年份
				if _, ok := years[year]; !ok {
					//没有，就存储进uniqueYears
					years[year] = struct{}{}
					uniqueYears = append(uniqueYears, year)
				}
				yearsList[i] = year
				i++
			}
		}
	}
	return yearsList, uniqueYears, nil
}

// 检查文件夹，如果没有则创建各个年份和标签的的文件夹
func Mkdir(years []int) error {
	// , tags map[string]int
	//年份文件夹
	const dirPerm = os.ModePerm
	for i := 0; i < len(years); i++ {
		dirExists, err := isDirExist(years[i])
		if err != nil {
			return fmt.Errorf("检查年份文件夹%d失败: %v", years[i], err)
		}
		if !dirExists {
			folder := strconv.Itoa(years[i])
			if err = os.MkdirAll("sources/articles/"+folder, dirPerm); err != nil {
				return fmt.Errorf("创建年份文件夹%s失败: %v", folder, err)
			}
		}
	}

	// //标签文件夹
	// //创建md对应的HTML文件
	// for k := range tags {
	// 	tag := k
	// 	//tag去过重，直接检查是否有文件夹，没有就创建
	// 	dirExists, err := isDirExist(tag)
	// 	if err != nil {
	// 		return fmt.Errorf("检查标签文件夹%s失败: %v", tag, err)
	// 	}
	// 	if !dirExists {
	// 		if err = os.MkdirAll("sources/articles/tagpage/"+tag, dirPerm); err != nil {
	// 			return fmt.Errorf("创建标签文件夹%s失败: %v", tag, err)
	// 		}
	// 	}
	// }
	return nil
}

// 检查年份和标签文件夹是否重复
func isDirExist(input interface{}) (bool, error) {
	var info fs.FileInfo
	var err error
	//判断输入类型
	switch in := input.(type) {
	case int:
		//判断年份
		path := strconv.Itoa(in)
		info, err = os.Stat("sources/articles/" + path)
	case string:
		//判断标签
		path := in
		info, err = os.Stat("sources/articles/tagpage/" + path)
	default:
		return false, fmt.Errorf("不支持的输入类型:%T", input)
	}

	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	//判断是否是目录
	if !info.IsDir() {
		return false, fmt.Errorf("存在但不是目录")
	}
	return true, nil
}

// 处理归档页的信息
func ExtArcInfo(uniqueYears []int) [][]string {
	arcinfo := make([][]string, len(uniqueYears))

	files, err := filepath.Glob("sources/articles/*/*.html")
	if err != nil {
		log.Fatalf("读取 Markdown 文件失败了: %v", err)
	}
	filename := make([]string, len(files))
	for i := 0; i < len(files); i++ {
		filename[i] = strings.TrimSuffix(filepath.Base(files[i]), ".html")
	}

	for i := 0; i < len(uniqueYears); i++ {
		arcinfo[i] = []string{strconv.Itoa(uniqueYears[i])}
		for j := 0; j < len(filename); j++ {
			if strings.Contains(files[j], strconv.Itoa(uniqueYears[i])) {
				arcinfo[i] = append(arcinfo[i], filename[j])
			}
		}
	}
	return arcinfo
}

// 搜索实现
func Search(w http.ResponseWriter, r *http.Request, data []Config) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// 针对预检请求进行处理
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	input := strings.TrimSpace(r.FormValue("input"))
	var response []string
	for _, item := range data {
		if strings.Contains(string(item.Title), input) && input != "" {
			response = append(response, item.MDPath)
		}
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 设置正确的响应头
	w.Header().Set("Content-Type", "application/json")

	// 返回 JSON 字符串
	w.Write(jsonResponse)
}

// tag页面动态生成
func Gentags(w http.ResponseWriter, r *http.Request, data struct {
	ConfigDict     []Config
	MdCount        int
	TagsCount      int
	TagNames       []string
	TagsInfo       map[string]int
	Archive_Year   []int
	Archive_MDInfo [][]string
}, tmpls *template.Template, uniqueTags map[string]int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// 针对预检请求进行处理
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	type TemplateData struct {
		Configs  []Config
		Tagname  string
		TagCount int
	}

	var configs []Config
	tagname := strings.TrimSpace(r.FormValue("tagname"))
	for i := 0; i < len(data.ConfigDict); i++ {
		for j := 0; j < len(data.ConfigDict[i].Tags); j++ {
			if strings.Contains(data.ConfigDict[i].Tags[j], tagname) {
				configs = append(configs, data.ConfigDict[i])
			}
		}
	}

	var tagcount int
	for key, value := range uniqueTags {
		if tagname == key {
			tagcount = value
		}
	}

	datas := TemplateData{
		Configs:  configs,
		Tagname:  tagname,
		TagCount: tagcount,
	}
	//创建输出文件并将模板引擎替换后的结果
	out_tag, err := os.Create("sources/articles/tag.html")
	if err != nil {
		log.Fatalf("创建home输出文件失败: %v", err)
	}
	errs := tmpls.ExecuteTemplate(out_tag, "tag.html", datas)
	if errs != nil {
		log.Fatalf("替换home模板中的占位符失败: %v", errs)
	}
	jsonResponse, err := json.Marshal(datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 设置正确的响应头
	w.Header().Set("Content-Type", "application/json")

	// 返回 JSON 字符串
	w.Write(jsonResponse)
}

// 数据按时间排序
func sortConfigs(configs []Config) {
	sort.Slice(configs, func(i, j int) bool {
		num1, err := strconv.Atoi(strings.ReplaceAll(configs[i].Time, "-", ""))
		if err != nil {
			fmt.Println("无法将字符串转换为整数")
		}
		num2, err := strconv.Atoi(strings.ReplaceAll(configs[j].Time, "-", ""))
		if err != nil {
			fmt.Println("无法将字符串转换为整数")
		}
		return num2 < num1
	})
}
