package logger

import (
	"log"
	"os"
)

var (
	Debug    *log.Logger
	Info     *log.Logger
	Warning  *log.Logger
	Error    *log.Logger
	Critical *log.Logger
)

func init() {
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	Critical = log.New(os.Stderr, "CRITICAL: ", log.Ldate|log.Ltime|log.Lshortfile)

	// debug does not get sent to NR
	Debug = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
}
