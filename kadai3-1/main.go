package main

import (
	"fmt"
	"os"

	"github.com/sourjp/gopherdojo-studyroom/kadai3-1/typing"
)

func main() {
	n, err := typing.Run()
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	os.Exit(n)
}
