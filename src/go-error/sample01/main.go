package main

import (
	"errors"
	"log"
)

func errorsNewSsample() error {
	err := errors.New("this is errors.New sample.")
	return err
}

func main() {
	err := errorsNewSsample()
	log.Printf("err: %+v", err)
	log.Printf("err: %T", err)
}
