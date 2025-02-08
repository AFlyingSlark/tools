package main

import (
	"encoding/json"
	"log"
	"os"
	"text/template"
)

// 定义配置结构体，字段名必须与 JSON 匹配
type Config struct {
	Title   string `json:"title"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Company string `json:"company"`
}

// 读取 JSON 配置文件
func loadConfig(filename string) (Config, error) {
	var config Config

	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	return config, err
}

func main() {
	// 配置和模板文件路径
	configFile := "config.json"
	templateFile := "http.tpl"
	outputFile := "http.txt"

	// 1. 读取 JSON 配置
	config, err := loadConfig(configFile)
	if err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
		return
	}

	// 2. 读取模板文件
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatalf("读取模板文件失败: %v", err)
		return
	}

	// 3. 创建输出文件
	outFile, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("创建输出文件失败: %v", err)
		return
	}
	defer outFile.Close()

	// 4. 渲染模板并写入文件
	err = tmpl.Execute(outFile, config)
	if err != nil {
		log.Fatalf("模板渲染失败: %v", err)
		return
	}

	log.Println("模板处理完成，结果已写入", outputFile)
}
