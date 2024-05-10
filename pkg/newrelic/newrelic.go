package newrelic

import (
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Application struct {
	*newrelic.Application
}

var application *Application

func GetNewrelicApplication() *Application {
	return application
}

func SetNewrelicApplication(newrelic *newrelic.Application) {
	application = &Application{newrelic}
}
