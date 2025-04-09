package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"temp_mail/email"
	"time"

	"github.com/go-co-op/gocron/v2"
)

var logger *log.Logger

func init() {
	// Create or open the log file
	logFile, err := os.OpenFile("main.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}

	// Set up the logger
	logger = log.New(logFile, "MAIN: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	// Log the start of the application
	logger.Println("Application started")

	// Call UpdateEmail once at the start
	logger.Println("Updating email list...")
	email.UpdateEmail()

	router := RegisterRoutes()

	// Initialize the scheduler
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		logger.Fatalf("Failed to initialize scheduler: %s", err)
	}

	// Schedule the UpdateEmail job
	job, err := scheduler.NewJob(
		gocron.DurationJob(7*24*time.Hour),
		gocron.NewTask(email.UpdateEmail),
	)
	if err != nil {
		logger.Fatalf("Failed to schedule job: %s", err)
	}

	logger.Printf("Scheduled job: %v", job)

	// Start the scheduler
	scheduler.Start()
	logger.Println("Scheduler started")

	// Start the HTTP server
	logger.Println("Server is running on port 8080")
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		logger.Fatalf("Failed to start server: %s", err)
	}
}
