package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestJsonMarshal(t *testing.T) {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}

	groups := []ColorGroup{
		{1, "reds", []string{"ruby", "maroon"}},
		{2, "blues", []string{"bb", "ll"}},
	}

	b, err := json.Marshal(groups)
	if err != nil {
		fmt.Println("error", err)
	}

	os.Stdout.Write(b)
}

func TestJsonUnmarshal(t *testing.T) {

}
