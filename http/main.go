package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWriter struct{}

func main() {
	res, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
	}

	lw := logWriter{}
	io.Copy(lw, res.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println("Writing out the values of the Body:", string(bs))
	fmt.Println("Length of bytes:", len(bs))
	return len(bs), nil
}
