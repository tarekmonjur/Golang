package main

import (
	"fmt"
	"io"
	"net/http"
)

// custom writer
type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}

func main() {
	var url string = "http://google.com"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}

	// fmt.Println(resp.Body)

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// resp.Body.Close()

	// io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)

}
