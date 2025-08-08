package services

import (
	"log"
	"scheduler/caching"
	"scheduler/db"
	"scheduler/metrics"
	"scheduler/models"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

func ExecuteJob(job models.Job) {
	log.Println("Executing job", job.JobId)
	err := db.DB.Model(&job).Where("job_id = ?", job.JobId).Update("last_run", time.Now()).Error
	if err != nil {
		log.Println("Failed to update last_run for job", job.JobId, err)
		metrics.SchedulerFailures.WithLabelValues(job.JobId).Inc()
	} else {
		metrics.SchedulerSuccesses.WithLabelValues(job.JobId).Inc()
	}
	cachedJob, found := caching.Cache.Get(job.JobId)
	if found {
		jobData := cachedJob.(map[string]interface{})
		jobData["last_run"] = time.Now()
		var mutex sync.Mutex
		mutex.Lock()
		caching.Cache.Set(job.JobId, jobData, cache.NoExpiration)
		mutex.Unlock()
	}
	log.Println("Successfully executed job: ", job.JobId)
}
