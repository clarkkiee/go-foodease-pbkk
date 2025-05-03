package middleware

import (
	"go-foodease-be/pkg/metrics"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context){
		start := time.Now()

		c.Next()

		duration := time.Since(start).Seconds()
		statusCode := strconv.Itoa(c.Writer.Status())

		metrics.HTTPRequestsTotal.WithLabelValues(c.Request.Method, c.FullPath()).Inc()
		metrics.HTTPRequestDuration.WithLabelValues(c.Request.Method, c.FullPath()).Observe(duration)
		metrics.HTTPResponseStatus.WithLabelValues(statusCode).Inc()
	}
}