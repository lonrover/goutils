package logger

import (
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

func InitDailyRotation(logPath, appName string) {
	// 设置日志路径格式
	logPathPattern := filepath.Join(logPath, appName)

	// 创建轮转器
	rotation, err := rotatelogs.New(
		logPathPattern,
		rotatelogs.WithLinkName(filepath.Join(logPath, appName)), // 当前日志软链接
		rotatelogs.WithRotationTime(24*time.Hour),                // 每天轮转一次
		rotatelogs.WithMaxAge(30*24*time.Hour),                   // 保留30天
		// rotatelogs.WithRotationCount(30),                                // 最多保留30个文件
	)

	if err != nil {
		panic(err)
	}

	logger := logrus.New()
	logger.SetOutput(rotation)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetReportCaller(true)
}
