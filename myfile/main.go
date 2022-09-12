package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fn := os.Args[1]
	f, err := os.Open(fn)
	if err != nil {
		fmt.Println("Error", err)
	}
	io.Copy(os.Stdout, f)
}
