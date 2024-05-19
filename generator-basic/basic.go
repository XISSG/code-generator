package generator_basic

import (
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

type model struct {
	Name string
	Loop bool
}

func basic() {
	currentPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	//设置输出和执行路径
	srcDir := filepath.Join(currentPath, "templates")
	dstDir := filepath.Join(currentPath, "generate")

	//静态文件
	dst := filepath.Join(dstDir, "static")
	src := filepath.Join(srcDir, "static")
	staticGenerate(src, dst)

	//动态文件
	outputDir := dstDir
	outputFile := path.Join(outputDir, "demo.go")
	templateFile := filepath.Join(srcDir, "demo.tpl")
	dynamicGenerate(outputFile, templateFile)

}

func dynamicGenerate(outputFile string, templateFile string) {
	templ, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Println(err)
		return
	}

	fd, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY, 0644)
	_ = templ.Execute(fd, &map[string]bool{"loop": true})

	fd.Close()
}

func staticGenerate(src string, dst string) {
	sr, err := os.OpenFile(src, os.O_RDONLY, 0644)
	de, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = io.Copy(de, sr)
	if err != nil {
		return
	}
	sr.Close()
	de.Close()
}
