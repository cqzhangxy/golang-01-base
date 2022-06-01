package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello", os.Args)

	os.Exit(-1)
}
