package gen

import (
	"io"
	"log"
	"os"
	"text/template"
)

func generate(wr io.Writer, templateFileName string, templateConfigElem templateConfigElem) error {
	t := template.Must(template.ParseFiles(templateFileName))
	err := t.Execute(wr, templateConfigElem)
	if err != nil {
		return err
	}

	return nil
}

// ParseTemplate 向指定writer写入渲染后的模板
func ParseTemplate(wr io.Writer, templateConfigData TemplateConfig) error {
	for _, configDatum := range templateConfigData {
		err := generate(wr, configDatum.TemplateFilePath, configDatum)
		if err != nil {
			log.Panic(err)
		}
	}
	return nil
}

// ParseTemplateWriteToFile 向配置的文件写入渲染后的模板
func ParseTemplateWriteToFile(templateConfig TemplateConfig) error {
	for _, configDatum := range templateConfig {
		f, err := openOutputFile(configDatum.OutputFile)
		if err != nil {
			log.Panic(err)
		}
		err = generate(f, configDatum.TemplateFilePath, configDatum)
		if err != nil {
			log.Panic(err)
		}
	}
	return nil
}

func openOutputFile(fileName string) (io.Writer, error) {
	var f *os.File
	var err error
	if checkFileIsExist(fileName) { //如果文件存在
		f, err = os.OpenFile(fileName, os.O_APPEND, 0666) //打开文件
	} else {
		f, err = os.Create(fileName) //创建文件
	}

	if err != nil {
		return nil, err
	}

	return f, nil
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
