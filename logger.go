package swagger

// Logger is the logger interface necessary by this package
type Logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}
