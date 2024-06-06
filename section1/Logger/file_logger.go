package Logger

import (
	"log"
	"os"
)

// Logger interface with a Log method
type Logger interface {
	Log(message string)
}

// FileLogger struct
type FileLogger struct {
	Filename string
}

// Log method implementation for FileLogger
func (f *FileLogger) Log(message string) {
	file, err := os.OpenFile(f.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	// Create a new logger that writes to the file
	logger := log.New(file, "", log.LstdFlags)
	logger.Println(message)
}

type ConsoleLogger struct{}

// Log method implementation for ConsoleLogger
func (c *ConsoleLogger) Log(message string) {
	log.Println(message)
}
