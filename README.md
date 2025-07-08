# DevTracker CLI 

**DevTracker** is a minimalist command-line tool to help developers log their coding activity, track daily progress, and stay consistent — all from the terminal.

> Your own developer journal, streak tracker, and dopamine machine — written in Go.

---

##  Features

-  Log daily coding sessions with `devtrack add`
-  View all your entries for today with `devtrack today`
- Optional tagging for filtering and future stats
-  Simple JSON storage — no database needed
-  weekly xp gained.
---

## Usage

### Add a new log entry:
```bash
devtrack add "Solved binary tree problem" --tag=leetcode
View today's logs:
bash
Copy
Edit
devtrack today
 Installation
1. Clone the repo
bash
Copy
Edit
git clone git@github.com:cargonew/devTracker.git
cd devTracker
2. Build the binary
bash
Copy
Edit
go build -o devtrack
3. (Optional) Move it to your PATH
bash
Copy
Edit
sudo mv devtrack /usr/local/bin
Now you can use devtrack from anywhere!..How cool is that!!

 Coming Soon
 devtrack stats — total logs, streaks, top tags

 Themes, badges, and gamification

 Fuzzy search and log filtering

 Daily journaling mode

 Built With
Go The best programming language ( just kidding haha)

Standard Library (no frameworks)

Terminal love 

 Log Format
All logs are saved in log.json:

json
Copy
Edit
[
  {
    "timestamp": "2025-07-04T21:50:00",
    "entry": "Solved array reversal problem",
    "tag": "leetcode"
  }
]
GOOD LUCK!!!
