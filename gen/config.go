package gen

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

// type templateConfigElem struct {
// 	PackageName      string `json:"packageName"`
// 	ClassRoot        string `json:"classRoot"`
// 	ClassName        string `json:"className"`
// 	ObjectName       string `json:"objectName"`
// 	GenerateDate     string `json:"generateDate"`
// 	TemplateFilePath string `json:"templateFilePath"`
// 	OutputFile       string `json:"outputFile"`
// }

type templateConfigElem struct {
	TemplateFile string            `json:"templateFile"`
	OutputFile   string            `json:"outputFile"`
	RenderData   map[string]string `json:"renderData"`
}

// TemplateConfig 模板配置
type TemplateConfig []templateConfigElem

// ParseConfig 解析配置
func ParseConfig(config []byte) (result TemplateConfig) {
	err := json.Unmarshal(config, &result)
	if err != nil {
		log.Panic("解析配置失败:", err)
	}
	for i := range result {
		if v, ok := result[i].RenderData["generateDate"]; !ok || v == "" {
			result[i].RenderData["generateDate"] = time.Now().Format(time.RFC3339)
		}
	}
	return
}

// ParseConfigFile 解析配置文件
func ParseConfigFile(fileName string) (result TemplateConfig) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panic("读取配置文件失败:", err)
	}
	return ParseConfig(data)
}
