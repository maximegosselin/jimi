package main

import (
	"fmt"
	"os"
)

func main() {
	fb, err := app(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(fb)
}
