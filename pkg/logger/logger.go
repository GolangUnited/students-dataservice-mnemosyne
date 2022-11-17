package logger

var (
	logger ILogger
)

func init() {
	logger = NewZapLogger()
}

type ILogger interface {
	Infos(args ...interface{})
	Info(msg string)
	Infof(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
}

func Info(msg string) {
	logger.Info(msg)
}

func Infos(args ...interface{}) {
	logger.Infos(args...)
}

func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}
