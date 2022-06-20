package logs

func NewLoggerService(loggerAdapter Logger) Logger {
	logger := loggerAdapter

	return logger
}
