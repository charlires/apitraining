package loggingadapter

type OOLogger interface {
	WithField(string, interface{}) OOLogger
	Info(...interface{})
	Error(...interface{})
}
