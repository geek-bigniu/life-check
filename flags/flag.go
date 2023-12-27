package flags

import (
	"flag"
	"fmt"
	"life-check/parse"
	"os"
)

// 自定义的用法说明
func customUsage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults() // 打印默认的帮助信息
	// 你可以添加任何自定义的帮助信息
	fmt.Fprintf(os.Stderr, "Examples:\n")
	fmt.Fprintf(os.Stderr, "  %s -u https://www.baidu.com -t 3 -n 10 \nor\n  %s -f url.txt -t 3 -n 10\n", os.Args[0], os.Args[0])
}
func ParseFlag() parse.Params {
	// 解析命令行参数
	url := flag.String("u", "", "check url(需要单独检测的url)")
	timeout := flag.Int("t", 3, "connect timeout(连接超时时间) * second")
	file := flag.String("f", "", "input with file(根据文件进行检测)")
	outFile := flag.String("o", "success.txt", "out file(输出文件)")
	threadNum := flag.Int("n", 50, "thread number(线程数量)")
	// 设置自定义的帮助信息
	flag.Usage = customUsage
	flag.Parse()
	if *url == "" && *file == "" {
		flag.Usage()
		os.Exit(2)
	}
	// 存储命令行参数
	params := parse.Params{
		Url:       *url,
		Timeout:   *timeout,
		File:      *file,
		ThreadNum: *threadNum,
		OutFile:   *outFile,
	}
	return params
}
