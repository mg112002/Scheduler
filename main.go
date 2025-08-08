package main

import (
	"encoding/json"
	"log"
	"os"
	"scheduler/caching"
	"scheduler/db"
	"scheduler/metrics"
	"scheduler/models"
	"scheduler/routes"
	"scheduler/services"

	"github.com/gin-gonic/gin"
)

func main() {
	//For PROD
	// conf, err := os.Open("manifest/PROD/config.json")
	//For LOCAL
	conf, err := os.Open("manifest/LOCAL/config.json")
	if err != nil {
		log.Fatalf("failed to open config file: %v", err)
	}
	defer conf.Close()

	decodeErr := json.NewDecoder(conf).Decode(&models.Config)
	if decodeErr != nil {
		log.Fatalf("failed to parse config file: %v", decodeErr)
	}

	db.InitDB()

	// Enable below code if caching needed
	caching.InitCache()
	// services.LoadCache()

	metrics.InitMetrics()

	// Initialize and start all job schedulers
	services.InitAndStartSchedulers()

	r := gin.Default()
	routes.SetupRoutes(r)

	log.Println("Scheduler service running on :8000")
	if err := r.Run(":8000"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
