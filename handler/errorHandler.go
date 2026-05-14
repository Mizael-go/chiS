package handler

import (
	"log"
)

// handle error
func ErrorHandler(props error) {
	log.Fatalf("error: %s", props.Error())
}
