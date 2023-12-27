package parse

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"life-check/check"
	"os"
)

// Params 用于存储命令行参数
type Params struct {
	Url       string
	Timeout   int
	File      string
	ThreadNum int
	OutFile   string
}

// 定义一个key类型，避免与其他context键冲突
type contextKey string

const ParamsKey = contextKey("params")

// Params2Instance 用于将命令行参数转换为实例
func Params2Instance(params Params) *check.LifeCheck {
	outFile, err := os.OpenFile(params.OutFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logrus.Error(errors.New(fmt.Sprintf("打开文件失败:%s", params.OutFile)))
		return nil
	}
	if params.Url != "" {
		return check.NewLifeCheck([]string{params.Url}, params.Timeout, params.ThreadNum, outFile)
	}
	var urls []string
	if params.File != "" {
		urls = parseFile(params.File)
		return check.NewLifeCheck(urls, params.Timeout, params.ThreadNum, outFile)
	}
	return nil
}

// parseFile 用于解析文件
func parseFile(file string) []string {
	//读取文件
	f, err := os.Open(file)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	defer f.Close()
	// 读取文件内容
	var urls []string
	for {
		var url string
		_, err := fmt.Fscanln(f, &url)
		if err != nil {
			if err == io.EOF {
				break
			}
			logrus.Error(err)
			return nil
		}
		urls = append(urls, url)
	}
	return urls
}
