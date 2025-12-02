package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

// Test function name must start with 'Test' to be recognized by 'go test'
func TestUpdateMessage(t *testing.T) {
	msg = "hello world"
	wg.Add(1)
	go updateMessage("hey world")
	wg.Wait()
	if msg != "hey world" {
		t.Errorf("Msg is not updating")
	}
}

func TestPrintMessage(t *testing.T) {
	//Normally os.Stdout is used to print in terminal so we are storing its original method
	OsStd := os.Stdout

	//w is write and r is read a pipe is buidling between both
	r, w, _ := os.Pipe()

	//Now whatever msg it is printing will be save in w
	//Now we are changing os.Stdout format from printing to terminal to writing in w
	os.Stdout = w

	msg = "Hello Test"
	printMessage()

	// it return error and writing is close so we can read it
	_ = w.Close()

	result, _ := io.ReadAll(r)
	var message string = string(result)

	os.Stdout = OsStd

	if !strings.Contains(message, "Hello Test") {
		t.Errorf("Expected %s", message)
	}

}

func TestMainFunction(t *testing.T) {
	//Normally os.Stdout is used to print in terminal so we are storing its original method
	OsStd := os.Stdout

	//w is write and r is read a pipe is buidling between both
	r, w, _ := os.Pipe()

	//Now whatever msg it is printing will be save in w
	//Now we are changing os.Stdout format from printing to terminal to writing in w
	os.Stdout = w

	main()

	// it return error and writing is close so we can read it
	_ = w.Close()

	result, _ := io.ReadAll(r)
	var message string = string(result)

	os.Stdout = OsStd

	if !strings.Contains(message, "Hello, world") || !strings.Contains(message, "Hello, cosmos") || !strings.Contains(message, "Hello, universe") {
		t.Errorf("Expected %s", message)
	}

}
