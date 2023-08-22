package logger

import (
	"github.com/sirupsen/logrus"

	"os"
	"io"
	"fmt"
)

type Logger struct {
	handler *logrus.Logger
}

func NewLogger() *Logger {
	my_logger := logrus.New()
	// 创建用于记录错误信息的文件
	log_file, err := os.OpenFile("./logger/log.txt", os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("warning! log.txt open error:", err)
		return nil
	}

	// 错误信息同时输出到文件和终端
	my_logger.SetOutput(io.MultiWriter(os.Stderr, log_file))
	// 设置输出WarnLevel及以上的信息
	my_logger.SetLevel(logrus.WarnLevel)
	return &Logger{
		handler: my_logger,
	}
}

// 简单封装了一下，能够输出出现的问题的请求IP
func(this *Logger) Warn(ipAddress string, info any) {
	entry := this.handler.WithField("IP", ipAddress)
	entry.Warn(info)
}

func(this *Logger) Error(ipAddress string, info any) {
	entry := this.handler.WithField("IP", ipAddress)
	entry.Error(info)
}

func(this *Logger) Fatal(ipAddress string, info any) {
	entry := this.handler.WithField("IP", ipAddress)
	entry.Fatal(info)
}

func(this *Logger) Panic(ipAddress string, info any) {
	entry := this.handler.WithField("IP", ipAddress)
	entry.Panic(info)
}