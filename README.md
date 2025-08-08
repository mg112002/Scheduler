# Scheduler Microservice (Go + Gin)

## Overview
A modular, production-ready scheduler microservice for job scheduling and management. Built with Go, Gin, GORM, and robfig/cron.

## Features
- **Job Scheduling:** Schedule jobs with flexible cron expressions, or with a start date and repeat interval.
- **API Endpoints:**
  - `GET /jobs` — List all jobs (uses cache if available)
  - `GET /jobs/:id` — Get job details (uses cache if available)
  - `POST /jobs` — Create a new job (with schedule, name, params, start date, and repeat interval)
  - `DELETE /jobs/:id` — Stop and delete a job (cancels running jobs and removes from schedule)
- **Database Integration:** SQLite (easy to swap for Postgres/MySQL)
- **Customizable Jobs:** Each job can have custom parameters and schedule.
- **Scalable Design:** Stateless API, DB-backed jobs, extensible scheduler.
- **SOLID Principles:** Modular, interface-driven, testable code.
- **In-memory Cache:** Job data is cached for fast retrieval on list and get endpoints.
- **Prometheus Metrics:** Exposes metrics for cache hits, cache size, scheduler successes, and failures.

## API Usage
### Create a Job (Cron-based)
```
POST /jobs
Content-Type: application/json
{
  "name": "Send Email",
  "schedule": "0 9 * * 1", // every Monday at 9am (cron format)
  "params": "{\"email\":\"user@example.com\"}"
}
```

### Create a Job (Start Date + Repeat Interval)
```
POST /jobs
Content-Type: application/json
{
  "name": "Data Backup",
  "start_date": "2024-06-10T09:00:00Z", // RFC3339 format
  "repeat_interval": "24h", // every 24 hours
  "params": "{\"path\":\"/backup\"}"
}
```

- `schedule` is optional if you provide `start_date` and `repeat_interval`.
- `start_date` is the first time the job will run (RFC3339 format).
- `repeat_interval` is a Go duration string (e.g., "24h", "7d").

### Stop and Delete a Job
```
DELETE /jobs/1
```
- Cancels a running job (if any) and removes it from the schedule and database.

### List Jobs
```
GET /jobs
```
- Returns jobs from cache if available, otherwise loads from DB and populates cache.

### Get Job by ID
```
GET /jobs/1
```
- Returns job from cache if available, otherwise loads from DB and populates cache.

## Metrics
Prometheus metrics are exposed for monitoring:
- `cache_hits`: Number of cache hits
- `cache_size`: Number of jobs in cache
- `scheduler_successes{job_id="..."}`: Number of successful job executions per job
- `scheduler_failures{job_id="..."}`: Number of failed job scheduling attempts per job
- `requests{type, status}`: Number of API requests by type and status

You can scrape these metrics from the `/metrics` endpoint (if enabled in your server setup).

## Setup & Run
1. **Install Go** (>=1.18)
2. **Install dependencies:**
   ```
   go mod tidy
   ```
3. **Run the service:**
   ```
   go run ./main.go
   ```
4. **API available at:** `http://localhost:8080`

## Notes
- **Job Execution:** For POC, jobs just print their name and params. Extend `models/scheduler.go` for real logic.
- **Cron Format:** Use standard cron expressions for scheduling.
- **Start Date + Interval:** Use `start_date` and `repeat_interval` for interval-based jobs.
- **Stopping Jobs:** Use `DELETE /jobs/:id` to stop and delete jobs. Running jobs are cancelled gracefully.
- **Cache:** Job list and get endpoints use in-memory cache for fast access.
- **Metrics:** Prometheus metrics are available for cache and scheduler monitoring.
- **Scalability:** Stateless API, DB-backed jobs, and modular scheduler allow for horizontal scaling and future distributed job execution.

## Extending
- Swap SQLite for Postgres/MySQL by changing GORM driver in `db/db.go`.
- Add authentication, logging, monitoring, or distributed job execution as needed.

---
**Author:** Your Name 
