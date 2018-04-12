package errors

import (
	"fmt"
	"os"
)

// Warning prints the message along with the error if there is one
func Warning(err error, msg string) bool {
	if err != nil {
		fmt.Println(msg + " - " + err.Error())
		return true
	}
	return false
}

// Fatal prints the message along with the error if there is one, and then stops the program
func Fatal(err error, msg string) {
	if err != nil {
		fmt.Println(msg + " - " + err.Error())
		os.Exit(1)
	}
}
