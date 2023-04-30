package kirk

// Logger represents the common logging methods supported.
type Logger interface {
	Error(err error)
	Panic(err error)
	Debug(message string)
	Info(message string)
	Warn(message string)
}
