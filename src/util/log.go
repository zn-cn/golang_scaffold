package util

import (
	"os"

	"github.com/sirupsen/logrus"
	// "github.com/zbindenren/logrus_mail"
)

// GetLogger 获取 logger
func GetLogger() *logrus.Logger {
	logger := &logrus.Logger{
		Out:       os.Stdout,
		Hooks:     make(logrus.LevelHooks),
		Formatter: new(logrus.JSONFormatter),
		Level:     logrus.WarnLevel,
	}
	// // TODO
	// hook, err := logrus_mail.NewMailAuthHook(
	// 	"logrus_email",
	// 	"smtp.gmail.com",
	// 	587,
	// 	"chenjian158978@gmail.com",
	// 	"271802559@qq.com",
	// 	"chenjian158978@gmail.com",
	// 	"xxxxxxx",
	// )
	// if err == nil {
	// 	logger.Hooks.Add(hook)
	// }
	return logger
}
