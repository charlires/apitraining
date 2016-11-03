package controllers_test

import (
	"fmt"
	"testing"

	"github.com/rhaseven7h/apitraining/controllers"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/rhaseven7h/apitraining/loggingadapter"
)

type MockLogrus struct {
	LogMessages       string
	LogEntry          loggingadapter.OOLogger
}

func NewMockLogrus() *MockLogrus {
	mockLogrus := &MockLogrus{
		LogMessages: "",
	}
	mockLogrus.LogEntry = mockLogrus;
	return mockLogrus
}

func (ml *MockLogrus) WithField(key string, value interface{}) loggingadapter.OOLogger {
	ml.LogMessages = ml.LogMessages + fmt.Sprintf("key=%#v value=%#v\n", key, value)
	return ml.LogEntry
}

func (ml *MockLogrus) Info(args ...interface{}) {
	ml.LogMessages = ml.LogMessages + fmt.Sprintf("INFO message=%v\n", args)
	return
}

func (ml *MockLogrus) Error(args ...interface{}) {
	ml.LogMessages = ml.LogMessages + fmt.Sprintf("ERROR message=%v", args)
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
			Convey("And I should get a info log", func() {
				So(logger.LogMessages, ShouldContainSubstring, "creating services controller")
			})
			Convey("And I should get a withFields log", func() {
				So(logger.LogMessages, ShouldContainSubstring, "key=\"input_id\" value=\"MyStringID\"")
			})
		})
	})
}
