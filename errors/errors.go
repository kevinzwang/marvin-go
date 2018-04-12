package errors

import (
	"log"
	"os"
)

var (
	warningLog *log.Logger
	errorLog   *log.Logger
	fatalLog   *log.Logger
)

func init() {
	warningLog = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Llongfile)
	errorLog = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
	fatalLog = log.New(os.Stdout, "FATAL ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
}

// Warning prints a warning
func Warning(msg string) {
	warningLog.Println(msg)
}

// Error prints an error
func Error(msg string) {
	errorLog.Println(msg)
}

// Fatal prints a fatal error
func Fatal(msg string) {
	fatalLog.Fatalln(msg)
}
