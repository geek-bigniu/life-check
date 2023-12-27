package check

import (
	"github.com/sirupsen/logrus"
	"net"
	"net/url"
	"os"
	"sync"
	"time"
)

type LifeCheck struct {
	Targets    []string
	Timeout    int
	wg         sync.WaitGroup
	threadChan chan int
	OutFile    *os.File
}

func NewLifeCheck(target []string, timeout int, threadNum int, outFile *os.File) *LifeCheck {
	// 检查常见的代理环境变量
	proxies := []string{"HTTP_PROXY", "HTTPS_PROXY", "NO_PROXY"}
	for _, proxy := range proxies {
		value := os.Getenv(proxy)
		if value != "" {
			logrus.Warnf("代理环境下检测可能不准确 - > %s is set to %s\n", proxy, value)
			break
		}
	}
	threadChan := make(chan int, threadNum)
	return &LifeCheck{Targets: target, Timeout: timeout, threadChan: threadChan, OutFile: outFile}
}

func (l *LifeCheck) Run() {
	defer l.OutFile.Close()
	for _, target := range l.Targets {
		l.wg.Add(1)
		l.threadChan <- 1
		go l.Job(target)
	}
	l.wg.Wait()

}
func (l *LifeCheck) Job(jobInfo string) {
	defer func() {
		l.wg.Done()
		<-l.threadChan
	}()
	ok := l.Check(jobInfo)
	if ok {
		logrus.Infof("%s is life", jobInfo)
		l.writeFile(jobInfo)
	} else {
		logrus.Errorf("%s is not life", jobInfo)
	}

}

// writeFile 写出文件
func (l LifeCheck) writeFile(url string) {
	// 写入文件
	_, err := l.OutFile.WriteString(url + "\n")
	if err != nil {
		logrus.Errorf("Failed to write to file: %s\n", err)
	}
}
func (l *LifeCheck) Check(u string) bool {
	parsedURL, err := url.Parse(u)
	if err != nil {
		return false
	}
	var port string
	var address string
	if parsedURL.Port() != "" {
		address = parsedURL.Host
	} else {
		if parsedURL.Scheme == "https" {
			port = "443" // 对于HTTPS，使用443端口
		} else {
			port = "80" // 对于HTTP，使用80端口
		}
		// 拼接域名和端口
		address = net.JoinHostPort(parsedURL.Host, port)
	}

	// 使用tcp连接检查目标是否存活
	conn, err := net.DialTimeout("tcp", address, time.Duration(l.Timeout)*time.Second) // 3 seconds timeout
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
