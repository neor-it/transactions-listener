package common

import (
	"fmt"
	"log"
)

// CatchPanic is a utility function to catch a panic and log the error.
func CatchPanic() {
	if r := recover(); r != nil {
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("%v", r)
		}

		log.Printf("PANIC - %s", err.Error())
	}
}
