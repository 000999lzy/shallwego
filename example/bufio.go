package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func bufioFunc() {
	// scanLines()
	// scanWords()
	writerFunc()

}

func writerFunc() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "hello,")
	fmt.Fprint(w, "world!\n")
	w.Flush()

	var buf bytes.Buffer
	w2 := bufio.NewWriter(&buf)

	data := "This is a buffered writer example.\n"
	n, err := w2.ReadFrom(strings.NewReader(data))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading data:", err)
		return
	}

	if err = w2.Flush(); err != nil {
		fmt.Fprintln(os.Stderr, "Error flushing buffer:", err)
		return
	}

	fmt.Printf("Wrote %d bytes to buffer:\n%s", n, buf.String())
}

func scanWords() {
	const input = "This is a sample input\nwith multiple lines.\n	Each line has words."
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)

	count := 0
	for scanner.Scan() {
		count++
		word := scanner.Text()
		fmt.Printf("Word %d: %s\n", count, word)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}

	fmt.Printf("Total words: %d\n", count)
}

func scanLines() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("You entered:", line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}
