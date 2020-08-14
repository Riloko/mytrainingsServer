package common

import "log"

// LogFatal ...
func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
