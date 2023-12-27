package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

func InitLog() {
	// 设置全局日志级别
	logrus.SetLevel(logrus.InfoLevel)

	// 使用TextFormatter，并设置ForceColors为true来强制启用颜色
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true, // 强制启用颜色，即使输出非TTY
		ForceQuote:  true,
	})

	// 设置全局日志输出（默认是stderr）
	logrus.SetOutput(os.Stdout)

	// 记录行号（可选）
	//logrus.SetReportCaller(true)
}
