package logger

import (
	"fmt"
	"time"
)

func Info(message string) {
	dt := time.Now().UTC()
	fmt.Println(dt.Format("01-02-2006 15:04:05") + " | INFO: " + message)
}

func Warning(message string) {
	dt := time.Now().UTC()
	fmt.Println(dt.Format("01-02-2006 15:04:05") + " | WARNING: " + message)
}

func Error(message string) {
	dt := time.Now().UTC()
	fmt.Println(dt.Format("01-02-2006 15:04:05") + " | ERROR: " + message)
}
