package main

import (
	"fmt"
)

func fmtErrorfSsample(str string) error {
	err := fmt.Errorf("this is fmt.Errorf sample. str: %s", str)
	return err
}

func main() {
	err := fmtErrorfSsample("hogehoge")
	fmt.Println(err)
	fmt.Printf("%T\n", err)
}
