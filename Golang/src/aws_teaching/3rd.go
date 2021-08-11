package main

import (
	"fmt"
	"io"
	_ "log"
	_ "os"
)

// YOUR CODE:
type WriteIt string

func (wi WriteIt) Write(bts []byte) (int, error) {
	fmt.Printf("Wanted: %s\n", bts)
	fmt.Println("Got: ", wi)
	return len([]byte(string(wi))), nil
}
func main() {
	var wrt WriteIt = "Haha you can't log anything"
	logger := New(wrt, "I'm a Prefix: ")
	logger.Printf("Hi there!")
}

// Their Code
type MyLogger struct {
	Prefix string
	W      io.Writer
}

func New(w io.Writer, prefix string) *MyLogger {
	return &MyLogger{Prefix: prefix, W: w}
}
func (ml *MyLogger) Printf(s string) {
	ml.W.Write([]byte(ml.Prefix + s))
}
