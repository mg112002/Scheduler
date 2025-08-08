package handlers

import (
	"log"
	"net/http"
	"scheduler/caching"
	"scheduler/db"
	"scheduler/metrics"
	"scheduler/models"

	"github.com/gin-gonic/gin"
)

func ListJobs(c *gin.Context) {
	jobList := make([]models.JobInfo, 0)
	if caching.Cache.ItemCount() > 0 {
		log.Println("Cache is used")
		for _, item := range caching.Cache.Items() {
			job := item.Object.(models.Job)
			jobList = append(jobList, models.JobInfo{
				JobId:          job.JobId,
				Name:           job.Name,
				CronExpression: job.CronExpression,
				StartDate:      job.StartDate,
				RepeatInterval: job.RepeatInterval,
				CreatedAt:      job.CreatedAt,
				LastRun:        job.LastRun,
			})
		}
	} else {
		jobs := make([]models.Job, 0)
		err := db.DB.Find(&jobs).Error
		if err != nil {
			metrics.Requests.WithLabelValues("list_jobs", "failure").Inc()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list jobs"})
			return
		}
		for _, job := range jobs {
			jobList = append(jobList, models.JobInfo{
				JobId:          job.JobId,
				Name:           job.Name,
				CronExpression: job.CronExpression,
				StartDate:      job.StartDate,
				RepeatInterval: job.RepeatInterval,
				CreatedAt:      job.CreatedAt,
				LastRun:        job.LastRun,
			})
		}
	}
	metrics.Requests.WithLabelValues("list_jobs", "success").Inc()
	c.JSON(http.StatusOK, jobList)
}
