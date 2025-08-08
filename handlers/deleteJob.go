package handlers

import (
	"net/http"
	"scheduler/db"
	"scheduler/metrics"
	"scheduler/models"

	"github.com/gin-gonic/gin"
)

func DeleteJob(c *gin.Context) {
	jobId := c.Param("job_id")
	err := db.DB.Model(&models.JobInfo{}).Where("job_id = ?", jobId).Delete(&models.JobInfo{}).Error
	if err != nil {
		metrics.Requests.WithLabelValues("delete_job", "failure").Inc()
		c.JSON(http.StatusNotFound, gin.H{"error": "job not found"})
		return
	}

	// caching.Cache.Delete(jobId)
	metrics.Requests.WithLabelValues("delete_job", "success").Inc()
	metrics.CacheSize.Dec()
	c.JSON(http.StatusOK, gin.H{"message": "Job deleted"})
}
