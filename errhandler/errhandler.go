package errhandler

import (
	"fmt"
)

// Handle prints the message along with the error if there is one
func Handle(err error, msg string) bool {
	if err != nil {
		fmt.Println("%v - %v", msg, err)
		return false
	}
	return true
}
