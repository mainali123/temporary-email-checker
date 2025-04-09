package email

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// Initialize the logger
var logger *log.Logger

func init() {
	// Create or open the log file
	logFile, err := os.OpenFile("email/email.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}

	// Set up the logger
	logger = log.New(logFile, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// IsTemporary checks if the given email belongs to a temporary email domain
func IsTemporary(email string) bool {
	// Validate email format
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	domain := strings.TrimSpace(parts[1]) // Trim any extra spaces

	// Open the file
	file, err := os.Open("email/emails.txt")
	if err != nil {
		logger.Printf("Error opening file: %s", err)
		return false
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// Skip comments and empty lines
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}
		// Check if the domain matches exactly
		if line == domain {
			return true
		}
	}

	// Handle potential scanner errors
	if err := scanner.Err(); err != nil {
		logger.Printf("Error reading file: %s", err)
		return false
	}

	return false
}

func UpdateEmail() {
	// Define the URL and output file
	url := "https://raw.githubusercontent.com/7c/fakefilter/main/txt/data.txt"
	outputFile := "email/emails.txt"

	// Check if the output file exists
	if _, err := os.Stat(outputFile); err == nil {
		// Generate a backup file name with a timestamp
		timestamp := time.Now().Format("20060102_150405")
		backupFile := "email/emails_backup_" + timestamp + ".txt"

		// Rename the existing file to the backup file
		err := os.Rename(outputFile, backupFile)
		if err != nil {
			logger.Printf("Error creating backup file: %s", err)
			return
		}
		logger.Printf("Backup created: %s", backupFile)
	}

	// Fetch content from the URL
	resp, err := http.Get(url)
	if err != nil {
		logger.Printf("Error fetching content: %s", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Printf("Failed to fetch URL: %s, status code: %d", url, resp.StatusCode)
		return
	}

	// Create the output file
	file, err := os.Create(outputFile)
	if err != nil {
		logger.Printf("Error creating file: %s", err)
		return
	}
	defer file.Close()

	// Write the content to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		logger.Printf("Error writing to file: %s", err)
		return
	}

	logger.Printf("Content successfully saved to %s", outputFile)
}
