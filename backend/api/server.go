package api

import (
	"demo-app-go/logger"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
)

// SetupRouter creates and run web server.
//
// Parameters:
// - N/A
//
// Returns:
// - N/A
func SetupRouter() *gin.Engine {
	apiMode := viper.GetString("api.mode")
	ginMode := gin.ReleaseMode
	if apiMode == "debug" {
		ginMode = gin.DebugMode
	}

	gin.SetMode(ginMode)

	router := gin.New()

	// Use logger for request logging
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		logger.Log.Infof(
			"HTTP request | method=%s path=%s status=%d latency=%s ip=%s error=%s",
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.ErrorMessage,
		)
		return ""
	}))

	router.Use(gin.Recovery())

	reg := prometheus.NewRegistry()
	reg.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	reg.MustRegister(collectors.NewGoCollector())

	c := NewController()

	router.GET("/status", c.Status)
	router.GET("/metrics", gin.WrapH(promhttp.HandlerFor(reg, promhttp.HandlerOpts{})))

	v1 := router.Group("/api/v1")
	{
		v1.GET("/hello", c.Hello)
	}

	return router
}
