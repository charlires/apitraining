package loggingadapter

import "github.com/Sirupsen/logrus"

type OOLogger interface {
	WithField(string, interface{}) *logrus.Entry
	Info(...interface{})
	Error(...interface{})
}
