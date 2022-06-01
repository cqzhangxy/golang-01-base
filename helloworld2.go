package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello,world")
	fmt.Println("hallo jay!")
	fmt.Println("hello", os.Args)

	os.Exit(-1)
}
