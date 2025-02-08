package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"text/template"
)

var (
	// 配置和模板文件路径
	configFile = "config.json"
)

// 定义配置结构体，字段名必须与 JSON 匹配
type Config struct {
	Conf      []confData `json:"conf"`
	Templates string     `json:"templates"`
}

// 渲染文件名及渲染信息
type confData struct {
	FileName string  `json:"fileName"`
	TplData  tplData `json:"tplData"`
}

// 模版中需要渲染的字段
type tplData struct {
	StructName string `json:"structName"`
}

// make local
// chmod +x binaryMain
func main() {
	// 1. 读取 JSON 配置
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
		return
	}

	// 2. 使用 ParseGlob 解析整个目录的模板文件
	tmpl, err := template.ParseGlob(config.Templates)
	if err != nil {
		log.Fatalf("读取模板文件失败: %v", err)
		return
	}

	// 3. 渲染模板并写入文件
	for _, value := range config.Conf {
		outputToFile(tmpl, value)
	}
}

// 渲染输出文件
func outputToFile(tmpl *template.Template, value confData) {
	var (
		// 构建文件名
		tplName    = strings.Join([]string{value.FileName, `tpl`}, `.`)
		targetFile = strings.Join([]string{value.FileName, `go`}, `.`)
	)

	outFile, err := os.Create(targetFile)
	if err != nil {
		log.Fatalf("创建输出文件失败: %v", err)

		return
	}

	defer func() {
		_ = outFile.Close()
	}()

	if err = tmpl.ExecuteTemplate(outFile, tplName, value.TplData); err != nil {
		log.Fatalf("模板渲染失败: %v", err)
		return
	}

	log.Println(targetFile, "渲染处理完成,结果已写入")
}

// 读取 JSON 配置文件
func loadConfig() (Config, error) {
	var config Config

	file, err := os.Open(configFile)
	if err != nil {
		return config, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)

	return config, err
}
