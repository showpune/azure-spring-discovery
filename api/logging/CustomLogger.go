package logging

import (
	"context"

	"github.com/go-logr/logr"
)

type CustomLogger struct {
	logr logr.Logger
}

func GetAzureLogger(ctx context.Context, annotationsMaps ...map[string]string) *CustomLogger {
	c := &CustomLogger{logr.Logger{}}
	return c
}

func (c *CustomLogger) Info(msg string, keysAndValues ...interface{}) {
	c.logr.V(0).Info(msg, keysAndValues...)
}

func (c *CustomLogger) Debug(msg string, keysAndValues ...interface{}) {
	c.logr.V(1).Info(msg, keysAndValues...)
}

func (c *CustomLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	c.logr.Error(err, msg, keysAndValues...)
}
