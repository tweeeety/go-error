package main

import (
	"fmt"
)

// カスタムErrorの構造体
type MyError struct {
	Msg  string
	Code int
}

// error interfaceを実装
func (err *MyError) Error() string {
	return fmt.Sprintf("ERROR: %d %s", err.Code, err.Msg)
}

// 何かする処理
func doSomething() error {
	return &MyError{Msg: "doSomething is unexpected error", Code: 30001}
}

func main() {
	if err := doSomething(); err != nil {
		fmt.Println(err)
		fmt.Printf("%T\n", err)
	}
}
