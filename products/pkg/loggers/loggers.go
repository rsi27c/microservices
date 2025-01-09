package loggers

import (
	"log"
	"os"
)

var (
	Warn  *log.Logger
	Info  *log.Logger
	Error *log.Logger
)

// creates a new log file or opens if file already exists
func OpenLog() {
	//creates a new log file or opens if file already exists
	file, err := os.OpenFile(os.Getenv("FILE_NAME"), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	//set the log prefix and flags
	Info = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
