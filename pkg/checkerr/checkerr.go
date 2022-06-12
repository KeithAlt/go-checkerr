package checkerr

import "log"

// Check if we have an error; shutdown if so
func checkPanic(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// Check if we have an error; log & shutdown
func checkPanicLog(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, "\n[StackTrace]:\n", err)
	}
}
