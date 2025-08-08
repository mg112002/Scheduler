package handlers

import (
	"net/http"
	"scheduler/db"
	"scheduler/metrics"
	"scheduler/models"
	"scheduler/services"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateJobRequest struct {
	Name           string `json:"name" binding:"required"`
	CronExpression string `json:"cron_expression"`
	StartDate      string `json:"start_date" binding:"required"`
	RepeatInterval string `json:"repeat_interval"`
	Params         string `json:"params"`
}

func CreateJob(c *gin.Context) {
	var req CreateJobRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.CronExpression == "" && req.RepeatInterval == "" {
		metrics.Requests.WithLabelValues("create_job", "failure").Inc()
		c.JSON(http.StatusBadRequest, gin.H{"error": "cron_expression or repeat_interval is required"})
		return
	}

	jobId := uuid.New().String()
	startTime, err := time.Parse(time.RFC3339, req.StartDate)
	if err != nil {
		metrics.Requests.WithLabelValues("create_job", "failure").Inc()
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_date format"})
		return
	}
	job := models.Job{
		JobId:          jobId,
		Name:           req.Name,
		CronExpression: req.CronExpression,
		StartDate:      startTime,
		RepeatInterval: req.RepeatInterval,
		Params:         req.Params,
		CreatedAt:      time.Now(),
		LastRun:        nil,
	}
	err = db.DB.Create(&job).Error
	if err != nil {
		metrics.Requests.WithLabelValues("create_job", "failure").Inc()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create job"})
		return
	}
	// var mutex sync.Mutex
	// mutex.Lock()
	// caching.Cache.Set(jobId, job, cache.NoExpiration)
	// mutex.Unlock()
	// metrics.CacheSize.Inc()

	go services.ScheduleJob(job)
	metrics.Requests.WithLabelValues("create_job", "success").Inc()
	c.JSON(http.StatusOK, gin.H{"message": "job created successfully", "job_id": jobId})
}
