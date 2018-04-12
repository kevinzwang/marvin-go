package errors

import (
	"log"
	"os"
)

var (
	Warning *log.Logger
	Error   *log.Logger
	Fatal   *log.Logger
)

func init() {
	Warning = log.New(os.Stdout, "[WARNING] ", log.Ldate|log.Ltime|log.Llongfile)
	Error = log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime|log.Llongfile)
	Fatal = log.New(os.Stdout, "[FATAL ERROR] ", log.Ldate|log.Ltime|log.Llongfile)
}
