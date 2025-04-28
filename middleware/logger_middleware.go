package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		// log details
		latency := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		userID, exists := c.Get("userId")

		logEntry := logrus.WithFields(logrus.Fields{
			"status":   status,
			"method":   method,
			"path":     path,
			"latency":  latency,
			"clientIP": clientIP,
		})
		if exists {
			if id, ok := userID.(uint); ok && id > 0 {
				logEntry.WithField("userID", id).Info("Request processed")
				return
			}
		}
		logEntry.Info("Request processed")
	}
}
