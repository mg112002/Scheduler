package models

import (
	"time"
)

type Job struct {
	ID             int        `gorm:"primaryKey" json:"id"`
	JobId          string     `gorm:"not null" json:"job_id"`
	Name           string     `gorm:"not null" json:"name"`
	CronExpression string     `gorm:"not null" json:"cron_expression"`
	StartDate      time.Time  `gorm:"not null" json:"start_date"`
	RepeatInterval string     `gorm:"not null" json:"repeat_interval"`
	Params         string     `gorm:"not null" json:"params"`
	CreatedAt      time.Time  `gorm:"not null" json:"created_at"`
	LastRun        *time.Time `gorm:"" json:"last_run"`
}

type JobInfo struct {
	JobId          string     `json:"job_id"`
	Name           string     `json:"name"`
	CronExpression string     `json:"cron_expression"`
	StartDate      time.Time  `json:"start_date"`
	RepeatInterval string     `json:"repeat_interval"`
	CreatedAt      time.Time  `json:"created_at"`
	LastRun        *time.Time `json:"last_run"`
}
