package errhandler

import (
	"fmt"
	"os"
)

// Handle prints the message along with the error if there is one
func Handle(err error, msg string) bool {
	if err != nil {
		fmt.Println(msg + " - " + err.Error())
		return true
	}
	return false
}

// HandleFatal prints the message along with the error if there is one, and then stops the program
func HandleFatal(err error, msg string) {
	if err != nil {
		fmt.Println(msg + " - " + err.Error())
		os.Exit(1)
	}
}
