package main

import (
	"github.com/sirupsen/logrus"
	"life-check/banner"
	"life-check/flags"
	"life-check/log"
	"life-check/parse"
)

func main() {
	banner.InitBanner("v0.0.1")
	log.InitLog()
	// 解析命令行参数
	params := flags.ParseFlag()
	// 将命令行参数转换为实例
	instance := parse.Params2Instance(params)
	if instance == nil {
		logrus.Error("No instance")
		return
	}
	// 运行实例
	instance.Run()
}
