package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/lonrover/goutils/common" // 确保这个包有文件操作函数
	"github.com/sirupsen/logrus"
)

var (
	once   sync.Once
	logger *logrus.Logger
)

type Log struct{}

// 获取调用者信息 (文件+行号)
func getCaller(skip int) (file string, line int) {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "???"
		line = 0
	}
	return filepath.Base(file), line // 只保留文件名，减少日志体积
}

// 初始化日志器 (线程安全，仅执行一次)
func Init(filePath, fileName string, enableConsole bool) {
	once.Do(func() {
		logger = logrus.New()
		logger.SetLevel(logrus.InfoLevel)
		// logger.SetFormatter(&logrus.JSONFormatter{})
		logger.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true, // 强制彩色输出
			FullTimestamp: true, // 显示完整时间
			DisableQuote:  true, // 不转义字符串
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				// 简化调用者信息显示
				filename := filepath.Base(f.File)
				return "", fmt.Sprintf("%s:%d", filename, f.Line)
			},
		})

		// 创建多路输出
		var writers []io.Writer

		// 总是启用控制台输出
		if enableConsole {
			writers = append(writers, os.Stdout)
		}

		// 文件输出（可选）
		if filePath != "" && fileName != "" {
			InitDailyRotation(filePath, fileName)
			fullPath := filepath.Join(filePath, fileName)

			if err := common.CreateFileIfNotExists(fullPath); err != nil {
				panic("无法创建日志文件: " + err.Error())
			}

			file, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				panic("无法打开日志文件: " + err.Error())
			}

			writers = append(writers, file)
		}

		// 设置多路输出
		if len(writers) > 0 {
			multiWriter := io.MultiWriter(writers...)
			logger.SetOutput(multiWriter)
		}
	})
}

// 带位置信息的日志条目
func withLocation() *logrus.Entry {
	file, line := getCaller(3) // 跳过3层调用栈
	return logger.WithFields(logrus.Fields{
		"file": file,
		"line": line,
	})
}

// ----------------------
// 日志方法 (线程安全)
// ----------------------
func (Log) Info(msg string, args ...interface{}) {
	withLocation().Info(msg, args)
}

func (Log) Warn(msg string, args ...interface{}) {
	withLocation().Warn(msg, args)
}

func (Log) Error(msg string, args ...interface{}) {
	withLocation().Error(msg, args)
}

func (Log) Debug(msg string, args ...interface{}) {
	withLocation().Debug(msg, args)
}

// 获取日志实例 (推荐使用全局变量)
func GetLogger() Log {
	return Log{}
}
