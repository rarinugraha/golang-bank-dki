package middleware

import (
	"bytes"
	"io"
	"time"

	"go-bank-dki/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.body.Write(b)
	return rw.ResponseWriter.Write(b)
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		var requestBody []byte
		if c.Request.Body != nil {
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			requestBody = bodyBytes
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		rw := &responseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = rw

		c.Next()

		duration := time.Since(startTime)

		statusCode := c.Writer.Status()
		responseSize := c.Writer.Size()
		responseBody := rw.body.String()

		logger.Logger.Info("Request and Response",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status_code", statusCode),
			zap.Duration("duration", duration),
			zap.ByteString("request_body", requestBody),
			zap.String("response_body", responseBody),
			zap.Int("response_size", responseSize),
		)
	}
}
