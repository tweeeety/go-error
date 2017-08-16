package main

import (
	"errors"
	"fmt"
)

func errorsNewSsample() error {
	err := errors.New("this is errors.New sample.")
	return err
}

func main() {
	err := errorsNewSsample()

	fmt.Println(err)
	fmt.Printf("%T\n", err)
}
