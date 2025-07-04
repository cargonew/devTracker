package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Entry     string    `json:"entry"`
	Tag       string    `json:"tag,omitempty"`
}

const logFile = "log.json"

func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	tag := addCmd.String("tag", "", "Optional tag for the entry")

	if len(os.Args) < 2 {
		fmt.Println("expected 'add' or 'today' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		if addCmd.NArg() < 1 {
			fmt.Println("Usage: devtrack add \"your message\" --tag=optionalTag")
			os.Exit(1)
		}
		entry := strings.Join(addCmd.Args(), " ")
		saveLog(entry, *tag)

	case "today":
		showToday()

	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}
func saveLog(message, tag string) {
	entry := LogEntry{
		Timestamp: time.Now(),
		Entry:     message,
		Tag:       tag,
	}

	var logs []LogEntry
	data, _ := os.ReadFile(logFile)
	if len(data) > 0 {
		json.Unmarshal(data, &logs)
	}

	logs = append(logs, entry)

	updated, _ := json.MarshalIndent(logs, "", "  ")
	os.WriteFile(logFile, updated, 0644)

	fmt.Println("✅ Logged:", message)
}

func showToday() {
	var logs []LogEntry
	data, err := os.ReadFile(logFile)
	if err != nil {
		fmt.Println("No logs found yet.")
		return
	}
	json.Unmarshal(data, &logs)

	today := time.Now().Format("2006-01-02")
	found := false
	for _, log := range logs {
		if log.Timestamp.Format("2006-01-02") == today {
			fmt.Printf("🕒 [%s] %s %s\n", log.Timestamp.Format("15:04"),
				log.Entry, emojiTag(log.Tag))
			found = true
		}
	}
	if !found {
		fmt.Println("No entries for today yet.")
	}
}

func emojiTag(tag string) string {
	if tag == "" {
		return ""
	}
	return fmt.Sprintf("🏷️ %s", tag)
}

