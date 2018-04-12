package errors

import (
	"fmt"
	"os"
	"runtime"

	"github.com/bwmarrin/discordgo"
)

// Warning displays a warning to stdout
func Warning(err error, msg string) bool {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		fmt.Printf("[WARNING] %s:%d: %v\n", fn, line, msg)

		return true
	}
	return false
}

// Error displays an error to stdout
func Error(err error, msg string) bool {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		fmt.Printf("[ERROR] %s:%d: %v\n", fn, line, msg)

		return true
	}
	return false
}

// Fatal displays an error to stdout and then exits
func Fatal(err error, msg string, session *discordgo.Session) {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		fmt.Printf("[FATAL ERROR] %s:%d: %v\n", fn, line, msg)

		if session != nil {
			session.Close()
		}
		os.Exit(1)
	}
}
