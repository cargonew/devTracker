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
//This is teh heart
// Breaks and exit if you do not tag a key command 
//
func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	tag := addCmd.String("tag", "", "Optional tag for the entry")

	if len(os.Args) < 2 {
		fmt.Println("expected 'add' or 'today' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] { //switch case for when you actually enter a valid cmd... it prints your long message 
						//Saves the log to the our Long Entry buddy..
	case "add":
		addCmd.Parse(os.Args[2:])
		if addCmd.NArg() < 1 {
			fmt.Println("Usage: devtrack add \"your message\" --tag=optionalTag")
			os.Exit(1)
		}
		entry := strings.Join(addCmd.Args(), " ")
		saveLog(entry, *tag) //Saving the log 

	case "today":
		showToday() //Shows what you did today

	default:
		fmt.Println("Unknown command:", os.Args[1]) //For handling the silly boys...
	}
}
//Writes the json log file at the exact time you logged the task and the tag  
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

	fmt.Println("Logged:", message) 
	rewardXp(entry)
}

//Shows all logs you logged today ( the day you run the cmd ofcoz ...)	
//If found it  print the time  you logged the task and the opt tag...else No logs
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



//Who doesnt like emojis
func emojiTag(tag string) string {
	if tag == "" {
		return ""
	}
	return fmt.Sprintf("🏷️ %s", tag)
}
//Tasks that are valid for some xp gain ( for now)
 var validXpGain = map[string]int{
	 
			"Learned Go": 100 ,
			"Learned Rust": 120,
			"Did easy leetcode": 20,
			"Did medium leetcode": 40,
			"Did hard leetcode": 70,
			"Learned a new vim motion/Trick": 1000}

//fuzzy finds task woeth some XP again.
//Reads the last log then maps a cubstring that validXpGain() has in one of the keys then the matching xp is Gained
func rewardXp( entry LogEntry){

	for key, xp := range validXpGain{ 

		if strings.Contains(entry.Entry, key){ 
			fmt.Printf("✨ Gained %v XP for: %q\n", key, xp)
			return
		}
	}
	fmt.Println("Definetely a waste of time!")
}





