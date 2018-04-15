package logger

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"time"

	"../marvin"
)

// Info logs info
func Info(msg string) {
	fprintf("[INFO]", msg, nil)
}

// Warning displays a warning to stdout
func Warning(err error, msg string) bool {
	if err != nil {

		fprintf("[WARNING]", msg, err)
		fmt.Printf("[WARNING] %v\n", msg)

		return true
	}
	return false
}

// Error displays an error to stdout
func Error(err error, msg string) bool {
	if err != nil {

		fprintf("[ERROR]", msg, err)
		fmt.Printf("[ERROR] %v\n", msg)

		return true
	}
	return false
}

// Fatal displays an error to stdout and then exits
func Fatal(err error, msg string) {
	if err != nil {

		fprintf("[FATAL ERROR]", msg, err)
		fmt.Printf("[FATAL ERROR] %v\n", msg)

		if marvin.Session() != nil {
			marvin.Session().Close()
		}

		os.Exit(1)
	}
}

func fprintf(flag string, msg string, err error) {
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

	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	if err != nil {
		fmt.Fprintf(w, "%s (%s) %v\n\tMessage: %v\n\tTraceback:\n", flag, time.Now().Format("Mon Jan _2 15:04:05 2006"), err, msg)

		count := 2
		for {
			pc, fn, line, ok := runtime.Caller(count)
			if !ok {
				break
			}
			details := runtime.FuncForPC(pc)
			name := ""
			if details != nil {
				name = details.Name()
			}
			fmt.Fprintf(w, "\t\t%s:%d: %s\n", fn, line, name)
			count++
		}
	} else {
		fmt.Fprintf(w, "%s (%s) %v\n", flag, time.Now().Format("Mon Jan _2 15:04:05 2006"), msg)
	}
}
