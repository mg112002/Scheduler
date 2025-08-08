package services

import (
	"log"
	"scheduler/caching"
	"scheduler/db"
	"scheduler/models"
	"sync"

	"github.com/patrickmn/go-cache"
)

func LoadCache() {
	jobs := []models.Job{}
	err := db.DB.Find(&jobs).Error
	if err != nil {
		log.Println("Failed to load jobs from db:", err)
	}
	var mutex sync.Mutex
	for _, job := range jobs {
		mutex.Lock()
		caching.Cache.Set(job.JobId, job, cache.NoExpiration)
		mutex.Unlock()
	}
	log.Println("Loaded", len(jobs), "jobs into cache")
}
