package logger

import "go.uber.org/zap"

var sLog *zap.SugaredLogger
func init() {
	log, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	sLog = log.Sugar()
}

func Error(args ...interface{}) {
	sLog.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	sLog.Errorf(template, args...)
}

func Errorw(msg string, kv ...interface{}) {
	sLog.Errorw(msg, kv...)
}

func Infow(msg string, kv ...interface{}) {
	sLog.Infow(msg, kv...)
}
