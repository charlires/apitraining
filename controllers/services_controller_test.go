package controllers_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/rhaseven7h/apitraining/controllers"
	. "github.com/smartystreets/goconvey/convey"
)

type MockLogrus struct {
	LogMessagesBuffer *bytes.Buffer
	LogMessages       string
	LogEntry          *logrus.Entry
}

func NewMockLogrus() *MockLogrus {
	return &MockLogrus{
		LogMessagesBuffer: bytes.NewBufferString(""),
		LogMessages:       "",
		LogEntry:          &logrus.Entry{},
	}
}

func (ml *MockLogrus) WithField(key string, value interface{}) *logrus.Entry {
	ml.LogMessages = ml.LogMessages + fmt.Sprintf("key=%#v value=%#v\n", key, value)
	lg := logrus.New()
	lg.Out = ml.LogMessagesBuffer
	ml.LogEntry = &logrus.Entry{
		Message: ml.LogMessages,
		Logger:  lg,
	}
	return ml.LogEntry
}

func (ml *MockLogrus) Info(args ...interface{}) {
	ml.LogMessages = ml.LogMessages + fmt.Sprintf("INFO message=%#v\n", args)
	return
}

func (ml *MockLogrus) Error(args ...interface{}) {
	ml.LogMessages = ml.LogMessages + fmt.Sprintf("ERROR message=%#v", args)
	return
}

func TestServicesController(t *testing.T) {
	Convey("Given a Services Controller Builder", t, func() {
		Convey("When I build a new Services Controller", func() {
			logger := NewMockLogrus()
			sc := controllers.NewServicesController("MyStringID", logger)
			Convey("Then I should get a valid ServicesController", func() {
				So(sc, ShouldNotBeNil)
			})
			Convey("And I should get a ServicesController with the ID I set", func() {
				So(sc.MyServiceID, ShouldEqual, "MyStringID")
			})
			Convey("And I should get a corresponding log", func() {
				So(logger.LogMessagesBuffer.String(), ShouldContainSubstring, "creating services controller")
			})
		})
	})
}
