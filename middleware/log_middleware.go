package middleware

import (
	"go-foodease-be/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)
		
		status := c.Writer.Status()

		if c.Request.URL.Path != "/metrics" {
			if status >= 500 {
				logger.Logger.Error("API Server Error",
					zap.String("method", c.Request.Method),
					zap.String("path", c.Request.URL.Path),
					zap.Int("status", status),
					zap.String("client_ip", c.ClientIP()),
					zap.String("user_agent", c.Request.UserAgent()),
					zap.Duration("duration", duration),
					zap.String("error", c.Errors.String()),
				)
			} else if status >= 400 {
				logger.Logger.Error("API Client Error",
					zap.String("method", c.Request.Method),
					zap.String("path", c.Request.URL.Path),
					zap.Int("status", status),
					zap.String("client_ip", c.ClientIP()),
					zap.String("user_agent", c.Request.UserAgent()),
					zap.Duration("duration", duration),
				)
			} else {
				logger.Logger.Info("API Request",
					zap.String("method", c.Request.Method),
					zap.String("path", c.Request.URL.Path),
					zap.Int("status", status),
					zap.String("client_ip", c.ClientIP()),
					zap.String("user_agent", c.Request.UserAgent()),
					zap.Duration("duration", duration),
				)
			}
		}

	}
}