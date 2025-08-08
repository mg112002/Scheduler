package routes

import (
	"net/http"
	"scheduler/handlers"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.GET("/jobs", handlers.ListJobs)
	r.GET("/jobs/:id", handlers.GetJob)
	r.POST("/jobs", handlers.CreateJob)
	r.DELETE("/jobs/:id", handlers.DeleteJob)
}
