package logger

import "log"

func Printf(format string, args ...any) {
	log.Printf(format+"\n", args...)
}
