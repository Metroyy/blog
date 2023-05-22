package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

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
	Time                   time.Time     `yaml:"time"`
	Tags                   []string      `yaml:"tags"`
	MDPath                 string
}

func main() {
	// 读取配置文件
	var configs []Config
	config := ReadConfig("config.yaml")

	files, err := filepath.Glob("sources/post/*.md")
	if err != nil {
		log.Fatalf("读取 Markdown 文件失败了: %v", err)
	}

	//计算md文件数量
	mdCount := CountMarkdownFiles()

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

	//处理归档页的信息
	arcinfo := ExtArcInfo(uniqueYears)
	//读取模板html文件
	tmpls := template.Must(template.ParseGlob("sources/templates/*.html"))

	//为模板传入的数据赋值
	data := struct {
		ConfigDict     []Config
		MdCount        int
		Archive_Year   []int
		Archive_MDInfo [][]string
	}{
		ConfigDict: configs,
		// 统计 sources/articles 文件夹下的 Markdown 文件数量
		MdCount:        mdCount,
		Archive_Year:   uniqueYears,
		Archive_MDInfo: arcinfo,
	}

	//提取出md文件名，存入files数组中
	for i := 0; i < mdCount; i++ {
		files[i] = ExtMakedownName(files[i])
	}

	//检查年份文件夹，如果没有则创建各个年份的文件夹
	MkdirYears(data.Archive_Year)

	// 生成 HTML 文件
	CreateHTML(tmpls, data, files, yearsList)
	fmt.Println("HTML模板中的占位符替换成功!")
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
	err = yaml.Unmarshal(mdFile, &mdConfig)
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
		{"time", config.Time.Format("2006-01-02")},
		{"tags", strings.Join(config.Tags, ",")},
	}

	for i := range configs {
		if v, ok := mdConfig[configs[i].Field]; ok {
			if configs[i].Field == "time" {
				timeStr := v
				config.Time, _ = time.Parse("2006-01-02", timeStr)
			} else if configs[i].Field == "tags" {
				config.Tags = strings.Split(strings.TrimSpace(v), ",")
			} else {
				configs[i].Value = v
			}
		}
	}

	config.Title = configs[0].Value
	config.Img = configs[1].Value
	config.Desc = configs[2].Value
	config.Tags = strings.Split(strings.TrimSpace(configs[4].Value), ", ")

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

// 生成 HTML 文件
func CreateHTML(tmpls *template.Template, data struct {
	ConfigDict     []Config
	MdCount        int
	Archive_Year   []int
	Archive_MDInfo [][]string
}, files []string, yearsList []int) {
	//创建md对应的HTML文件
	for i := 0; i < data.MdCount; i++ {
		out_md, err := os.Create("sources/articles/" + strconv.Itoa(yearsList[i]) + "/" + files[i] + ".html")
		if err != nil {
			log.Fatalf("创建"+files[i]+"输出文件失败: %v", err)
		}
		errs := tmpls.ExecuteTemplate(out_md, "index.html", data.ConfigDict[i])
		if errs != nil {
			log.Fatalf("替换"+files[i]+"模板中的占位符失败: %v", errs)
		}
	}

	// 创建输出文件并将模板引擎替换后的结果
	out_home, err := os.Create("sources/articles/home.html")
	if err != nil {
		log.Fatalf("创建home输出文件失败: %v", err)
	}

	out_archive, err := os.Create("sources/articles/archive.html")
	if err != nil {
		log.Fatalf("创建index输出文件失败: %v", err)
	}

	errs := tmpls.ExecuteTemplate(out_home, "home.html", data)
	if errs != nil {
		log.Fatalf("替换home模板中的占位符失败: %v", errs)
	}

	err = tmpls.ExecuteTemplate(out_archive, "archive.html", data)
	if err != nil {
		log.Fatalf("替换archive模板中的占位符失败: %v", err)
	}
}

// 处理路径，提取md文件名
func ExtMakedownName(files string) string {
	files = strings.TrimPrefix(strings.TrimSuffix(files, ".md"), "sources\\post\\")
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
		//获取年份
		year := config.Time.Year()
		//判断map中有没有这个年份
		if _, ok := years[year]; !ok {
			//没有，就存储进uniqueYears
			years[year] = struct{}{}
			uniqueYears = append(uniqueYears, year)
		}
		yearsList[i] = year
		i++
	}
	return yearsList, uniqueYears, nil
}

// 检查年份文件夹，如果没有则创建各个年份的文件夹
func MkdirYears(years []int) error {
	const dirPerm = os.ModePerm
	for i := 0; i < len(years); i++ {
		dirExists, err := isDirExist(years[i])
		if err != nil {
			return fmt.Errorf("检查年份文件夹%d失败: %v", years[i], err)
		}
		if !dirExists {
			folder := strconv.Itoa(years[i])
			if err = os.MkdirAll("sources/articles/"+folder, dirPerm); err != nil {
				return fmt.Errorf("创建文件夹%s失败: %v", folder, err)
			}
		}
	}
	return nil
}

// 检查年份文件夹是否重复
func isDirExist(year int) (bool, error) {
	//判断是否有文件
	path := strconv.Itoa(year)
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	//判断是否是目录
	if !info.IsDir() {
		return false, fmt.Errorf("%s 存在但不是目录", path)
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
