package logger

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Info logs info
func Info(msg string) {
	fprintf("[INFO]", "", 0, msg, nil)
}

// Warning displays a warning to stdout
func Warning(err error, msg string) bool {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)

		fprintf("[WARNING]", fn, line, msg, err)
		fmt.Printf("[WARNING] %v\n", msg)

		return true
	}
	return false
}

// Error displays an error to stdout
func Error(err error, msg string) bool {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)

		fprintf("[ERROR]", fn, line, msg, err)
		fmt.Printf("[ERROR] %v\n", msg)

		return true
	}
	return false
}

// Fatal displays an error to stdout and then exits
func Fatal(err error, msg string, session *discordgo.Session) {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)

		fprintf("[FATAL ERROR]", fn, line, msg, err)
		fmt.Printf("[FATAL ERROR] %v\n", msg)

		if session != nil {
			session.Close()
		}

		os.Exit(1)
	}
}

func fprintf(flag string, fn string, line int, msg string, err error) {
	var f *os.File
	if _, err := os.Stat("log.txt"); err == nil {
		f, err = os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			fmt.Println("Error opening file")
			return
		}
	} else {
		f, err = os.Create("log.txt")
		if err != nil {
			fmt.Println("Error creating file")
			return
		}
	}

	w := bufio.NewWriter(f)

	// fmt.Fprintf(w, msg, a)
	if err != nil {
		fmt.Fprintf(w, "%s (%s) %s:%d: %v\n%v\n", flag, time.Now().Format("Mon Jan _2 15:04:05 2006"), fn, line, msg, err)
	} else {
		fmt.Fprintf(w, "%s (%s) %v\n", flag, time.Now().Format("Mon Jan _2 15:04:05 2006"), msg)
	}

	w.Flush()
	f.Close()
}
