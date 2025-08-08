package handlers

import (
	"net/http"
	"scheduler/caching"
	"scheduler/db"
	"scheduler/metrics"
	"scheduler/models"

	"github.com/gin-gonic/gin"
)

func GetJob(c *gin.Context) {
	jobId := c.Param("id")
	job, found := caching.Cache.Get(jobId)
	if !found {
		var job models.Job
		err := db.DB.Where("job_id = ?", jobId).First(&job).Error
		if err != nil {
			metrics.Requests.WithLabelValues("get_job", "failure").Inc()
			c.JSON(http.StatusNotFound, gin.H{"error": "job not found"})
			return
		}

		jobInfo := models.JobInfo{
			JobId:          job.JobId,
			Name:           job.Name,
			CronExpression: job.CronExpression,
			StartDate:      job.StartDate,
			RepeatInterval: job.RepeatInterval,
			CreatedAt:      job.CreatedAt,
			LastRun:        job.LastRun,
		}
		// var mutex sync.Mutex
		// mutex.Lock()
		// caching.Cache.Set(jobId, jobInfo, cache.NoExpiration)
		// mutex.Unlock()
		metrics.CacheSize.Inc()
		metrics.Requests.WithLabelValues("get_job", "success").Inc()
		c.JSON(http.StatusOK, jobInfo)
	} else {
		metrics.Requests.WithLabelValues("get_job", "success").Inc()
		c.JSON(http.StatusOK, job)
	}
}
