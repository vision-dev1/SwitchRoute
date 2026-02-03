// Codes by vision
package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type RequestLog struct {
	Timestamp    string `json:"timestamp"`
	Proxy        string `json:"proxy"`
	URL          string `json:"url"`
	Status       string `json:"status"`
	ResponseCode int    `json:"response_code"`
	Error        string `json:"error,omitempty"`
}

type Logger struct {
	logDir string
	mu     sync.Mutex
}

func New(logDir string) (*Logger, error) {
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %w", err)
	}

	return &Logger{
		logDir: logDir,
	}, nil
}

func (l *Logger) Log(proxy, url, status string, responseCode int, err error) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	logEntry := RequestLog{
		Timestamp:    time.Now().Format(time.RFC3339),
		Proxy:        proxy,
		URL:          url,
		Status:       status,
		ResponseCode: responseCode,
	}

	if err != nil {
		logEntry.Error = err.Error()
	}

	logFile := filepath.Join(l.logDir, fmt.Sprintf("requests_%s.json", time.Now().Format("2006-01-02")))
	
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(logEntry); err != nil {
		return fmt.Errorf("failed to write log entry: %w", err)
	}

	l.printSummary(logEntry)
	return nil
}

func (l *Logger) printSummary(log RequestLog) {
	fmt.Printf("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	fmt.Printf("📊 Request Summary\n")
	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	fmt.Printf("Time:     %s\n", log.Timestamp)
	fmt.Printf("Proxy:    %s\n", log.Proxy)
	fmt.Printf("URL:      %s\n", log.URL)
	fmt.Printf("Status:   %s\n", log.Status)
	if log.ResponseCode > 0 {
		fmt.Printf("Code:     %d\n", log.ResponseCode)
	}
	if log.Error != "" {
		fmt.Printf("Error:    %s\n", log.Error)
	}
	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n\n")
}
