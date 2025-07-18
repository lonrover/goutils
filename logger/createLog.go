package logger

import (
	"fmt"
	"os"
	"runtime"

	"github.com/lonrover/goutils/common"

	"github.com/sirupsen/logrus"
)

type Log struct {
	// 日志等级
	InfoLevel logrus.Level
	// 日志输出格式
	logOutFormat logrus.Formatter

	// 文件名
	filename string
}

func (log Log) logWithLocation() *logrus.Entry {
	logger := log.settings()
	entry := logger.WithFields(logrus.Fields{
		"file": getCallerFile(),
		"line": getCallerLine(),
	})

	return entry
}

func (log *Log) Info(msg string, arg ...interface{}) {
	entry := log.logWithLocation()
	entry.Info(msg, arg)

}

func (log Log) Warn(msg string, arg ...interface{}) {
	entry := log.logWithLocation()
	entry.Warn(msg, arg)

}

func (log *Log) Error(msg string, arg ...interface{}) {
	entry := log.logWithLocation()
	entry.Error(msg, arg)

}

func (log Log) Debug(msg string, arg ...interface{}) {
	entry := log.logWithLocation()
	entry.Debug(msg, arg)

}

// 获取输出日志的文件信息
func getCallerFile() string {
	_, file, _, ok := runtime.Caller(2)
	if !ok {
		file = "???"
	}
	return file
}

// 定位输出日志的行数
func getCallerLine() int {
	_, _, line, ok := runtime.Caller(2)
	if !ok {
		line = 0
	}
	return line
}

func (_log Log) settings() *logrus.Logger {
	// 创建logrus实例
	logger := logrus.New()

	// 设置日志级别
	logger.SetLevel(_log.InfoLevel)

	// 检查目录是否存在，如果不存在则创建目录
	err := common.CreateFileIfNotExists(_log.filename)

	// 设置输出格式为JSON格式
	logger.SetFormatter(_log.logOutFormat)

	// 打开日志文件，也可以根据日期生成文件名
	logFileName := _log.filename
	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		logger.Fatal("Error opening log file:", err)
	}

	// 设置日志输出到文件
	logger.SetOutput(file)

	return logger
}

func InitLog(filePath string, filename string) *Log {
	// filename := "log/task_log" + time.Now().Format("2006-01-02") + ".json"
	log := Log{
		InfoLevel:    logrus.InfoLevel,
		logOutFormat: &logrus.JSONFormatter{},
		filename:     filePath + filename,
	}

	fmt.Println(log)

	return &log
}

// var Logs = InitLog()
