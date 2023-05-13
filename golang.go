package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
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
}

func main() {
	// 读取配置文件
	config := ReadConfig("config.yaml")

	files, err := filepath.Glob("sources/articles/*.md")
	if err != nil {
		log.Fatalf("读取 Markdown 文件失败: %v", err)
	}

	// ExtractMarkdown 从 Markdown 文件中提取内容并更新 Config 配置
	for _, path := range files {
		config = ExtractMarkdown(path, config)
	}

	//读取模板html文件
	tmpls := template.Must(template.ParseGlob("templates/*.html"))

	//为模板传入的数据赋值
	data := struct {
		Config
		MdCount int
	}{
		Config: config,
		// 统计 sources/articles 文件夹下的 Markdown 文件数量
		MdCount: CountMarkdownFiles(),
	}
	// 生成 HTML 文件
	CreateHTML(tmpls, data)
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
func ExtractMarkdown(path string, config Config) Config {
	// 读取 Markdown 文件
	mdFile, err := os.ReadFile(path)
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

	configs := [][]string{
		{"title", config.Title},
		{"img", config.Img},
		{"desc", config.Desc},
		{"time", config.Time},
		{"tags", strings.Join(config.Tags, ",")},
	}

	for _, c := range configs {
		if v, ok := mdConfig[c[0]]; ok {
			c[1] = v
		}
	}

	config.Title = configs[0][1]
	config.Img = configs[1][1]
	config.Desc = configs[2][1]
	config.Time = configs[3][1]
	config.Tags = strings.Split(configs[4][1], ",")

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
	mdCount := 0
	err := filepath.Walk("sources/articles", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".md") {
			mdCount++
		}
		return nil
	})
	if err != nil {
		log.Fatalf("遍历目录失败: %v", err)
	}
	return mdCount
}

// 生成 HTML 文件
func CreateHTML(tmpls *template.Template, data interface{}) {
	// 创建输出文件并将模板引擎替换后的结果
	out_home, err := os.Create("pages/home.html")
	if err != nil {
		log.Fatalf("创建home输出文件失败: %v", err)
	}
	out_index, err := os.Create("pages/index.html")
	if err != nil {
		log.Fatalf("创建index输出文件失败: %v", err)
	}
	out_archive, err := os.Create("pages/archive.html")
	if err != nil {
		log.Fatalf("创建index输出文件失败: %v", err)
	}

	err = tmpls.ExecuteTemplate(out_home, "home.html", data)
	if err != nil {
		log.Fatalf("替换home模板中的占位符失败: %v", err)
	}
	err = tmpls.ExecuteTemplate(out_index, "index.html", data)
	if err != nil {
		log.Fatalf("替换index模板中的占位符失败: %v", err)
	}
	err = tmpls.ExecuteTemplate(out_archive, "archive.html", data)
	if err != nil {
		log.Fatalf("替换archive模板中的占位符失败: %v", err)
	}
}
