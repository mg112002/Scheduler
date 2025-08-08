package services

import (
	"log"
	"scheduler/db"
	"scheduler/metrics"
	"scheduler/models"

	"github.com/robfig/cron/v3"
)

var crons = make(map[string]*cron.Cron)

func InitAndStartSchedulers() {
	var jobs []models.Job
	err := db.DB.Find(&jobs).Error
	if err != nil {
		log.Printf("Failed to load jobs from db: %v", err)
		return
	}
	for _, job := range jobs {
		go ScheduleJob(job)
	}
}

func ScheduleJob(job models.Job) {
	if job.CronExpression != "" {
		c := cron.New()
		_, err := c.AddFunc(job.CronExpression, func() {
			log.Printf("Running job %s (cron): %s", job.JobId, job.Name)
			ExecuteJob(job)
		})
		if err == nil {
			crons[job.JobId] = c
			c.Start()
		} else {
			metrics.SchedulerFailures.WithLabelValues(job.JobId).Inc()
			log.Printf("Failed to schedule cron job %s: %v", job.JobId, err)
		}
	} else if job.RepeatInterval != "" {
		expr := "@every " + job.RepeatInterval
		c := cron.New()
		_, err := c.AddFunc(expr, func() {
			log.Printf("Running job %s (interval): %s", job.JobId, job.Name)
			ExecuteJob(job)
		})
		if err == nil {
			crons[job.JobId] = c
			c.Start()
		} else {
			metrics.SchedulerFailures.WithLabelValues(job.JobId).Inc()
			log.Printf("Failed to schedule interval job %s: %v", job.JobId, err)
		}
	}
}
