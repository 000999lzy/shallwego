package main

import (
	"bytes"
	"fmt"
	"log"
	"testing"
)

func TestLoggerFunc(t *testing.T) {
	var buf bytes.Buffer

	logger := log.New(&buf, "logger-gogogo", log.Lshortfile)
	logger.Print("hello, log file")

	logger.Output(2, "info")

	fmt.Print(&buf)

}
