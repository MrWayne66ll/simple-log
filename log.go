package simple_log

var (
	logger *Logger
)

// Sync 刷新日志到磁盘，退出前执行
func Sync() error {
	return logger.logger.Sync()
}

func Info(args ...interface{}) {
	logger.logger.Sugar().Info(args...)
}

func Warn(args ...interface{}) {
	logger.logger.Sugar().Warn(args...)
}

func Error(args ...interface{}) {
	logger.logger.Sugar().Error(args...)
}

func DPanic(args ...interface{}) {
	logger.logger.Sugar().DPanic(args...)
}

func Panic(args ...interface{}) {
	logger.logger.Sugar().Panic(args...)
}

func Fatal(args ...interface{}) {
	logger.logger.Sugar().Fatal(args...)
}

func Debugf(template string, args ...interface{}) {
	logger.logger.Sugar().Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	logger.logger.Sugar().Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	logger.logger.Sugar().Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	logger.logger.Sugar().Errorf(template, args...)
}

func DPanicf(template string, args ...interface{}) {
	logger.logger.Sugar().DPanicf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	logger.logger.Sugar().Panicf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	logger.logger.Sugar().Fatalf(template, args...)
}
