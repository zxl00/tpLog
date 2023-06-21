package tpLog

import (
	"github.com/sirupsen/logrus"
	"os"
)

type MyFormatter struct{}

var TPLogger *logrus.Logger

func (f *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	msg := entry.Message
	level := entry.Level
	time := entry.Time.Format("2006-01-02 15:04:05")
	return []byte(time + " [" + level.String() + "] " + msg + "\n"), nil
}

func init() {
	// 创建 Logrus 日志实例
	TPLogger = logrus.New()
	// 设置日志级别为 Debug
	TPLogger.SetLevel(logrus.DebugLevel)
	// 输出到标准输出
	TPLogger.SetOutput(os.Stdout)
	// 使用自定义日志格式
	TPLogger.SetFormatter(&MyFormatter{})
}
