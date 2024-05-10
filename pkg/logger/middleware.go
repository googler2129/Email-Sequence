package logger

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Create a custom ResponseWriter to capture the response body
		bodyWriter := &responseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		// Process the request
		requestBody, _ := c.GetRawData()
		c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		requestBodyString := string(requestBody)
		c.Next()

		// Calculate the response time
		duration := time.Since(start)

		// Log the request and response
		if c.Request.URL.Path != "/health" {
			Infof("HTTP Request [%s] %s - %v\nRequest Body: %s\n HTTP Response: %s\n",
				c.Request.Method,
				c.Request.URL.Path,
				duration,
				requestBodyString,
				bodyWriter.body.String(),
			)
		}
	}
}
