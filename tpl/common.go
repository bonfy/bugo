package tpl

import (
	"github.com/russross/blackfriday"
	"path"
	"strings"
)

// 将 markdown 转化成 html
func conver_markdown_to_html(input string) string {
	text := []byte(input)
	return string(blackfriday.Run(text))
}

// 获取 目标地址 (注意 要将 .md 换成 .html)
func get_dist_path(filename string) string {
	suffix := path.Ext(filename)
	// 去除掉后缀
	filenameonly := strings.TrimSuffix(filename, suffix)
	// 构建distpath
	distpath := path.Join("dist", filenameonly + ".html")

	return distpath
}

// 获取 源目录地址
func get_source_path(filename string) string {
	return path.Join("content", filename)
}


// 检查 是否为 md file
func check_md_file (filename string) bool {
	fileSuffix := path.Ext(filename) //获取文件后缀
	// fmt.Println(fileSuffix)
	if strings.ToLower(fileSuffix) == ".md" {
		return true
	}

	return false
}