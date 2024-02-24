package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

const (
	invalidJSONMsg = "invalid JSON"
	validJSONMsg   = "valid JSON"
)

func main() {
	var jsonData []byte

	stat, _ := os.Stdin.Stat()

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		jsonData, _ = io.ReadAll(os.Stdin)
	} else {
		f, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer f.Close()
		bRead := bufio.NewReader(f)
		for {
			line, _, err := bRead.ReadLine()
			jsonData = append(jsonData, line...)
			if err != nil {
				break
			}
		}
	}
	if json.Valid(jsonData) {
		fmt.Printf("[%s] received at %s\n", validJSONMsg, time.Now())
		os.Exit(0)
	}
	fmt.Printf("[%s] received at %s", invalidJSONMsg, time.Now())
}
