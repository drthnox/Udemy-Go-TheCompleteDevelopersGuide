package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
  resp, err := http.Get("http://www.google.com")
  if err != nil {
    fmt.Println("Error: ", err)
    os.Exit(1)
  }

  // version 1
  // bs := []byte{}

  // version 2
  // bs := make([]byte, 99999) // predefined slice size
  // resp.Body.Read(bs)
  // fmt.Println(string(bs))

  // version 3
  io.Copy(os.Stdout, resp.Body)

}