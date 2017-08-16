package main

import (
	"fmt"
	"log"
)

func fmtErrorfSsample(str string) error {
	err := fmt.Errorf("this is fmt.Errorf sample. str: %s", str)
	return err
}

func main() {
	err := fmtErrorfSsample("hogehoge")
	log.Printf("err: %+v", err)
	log.Printf("err: %T", err)
}
