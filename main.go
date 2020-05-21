package main

import (
	"flag"
	"log"

	"github.com/wangxianzhuo/template-gen/gen"
)

var (
	configFile = flag.String("config", "./config.json", "配置文件")
)

func main() {
	flag.Parse()

	config := gen.ParseConfigFile(*configFile)

	err := gen.ParseTemplateWriteToFile(config)
	if err != nil {
		log.Panic(err)
	}
}
