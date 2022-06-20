package logs

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}

type ConfigLogger struct {
	Color     bool
	TimeStamp bool
}
