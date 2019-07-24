package main

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestPipeline(t *testing.T) {
	color := NewPrettyHandler(&ColorOptions{
		EnableColor: false,
	})
	filter := NewFilterHandler(DefaultFilterOptions)

	input := &bytes.Buffer{}
	output := &bytes.Buffer{}

	fmt.Fprintln(input, `{"key":"foo"}`)
	fmt.Fprintln(input, `{"key":"bar"}`)

	pipeline := NewPipeline(input, output, filter, color)
	done := make(chan struct{})
	logs := make(chan string)
	fatal := make(chan string)

	go func() {
		logs <- "Start pipeline go routine"
		err := pipeline.Run()
		if err != nil {
			fatal <- fmt.Sprintf("Unexpected error: %+v", err)
		}
		logs <- "Stop pipeline go routine"
		done <- struct{}{}
	}()

	for {
		select {
		case <-done:
			eol := "\n"
			expected := fmt.Sprint(
				`{`, eol,
				`  "key": "foo"`, eol,
				`}`, eol,
				`{`, eol,
				`  "key": "bar"`, eol,
				`}`, eol,
			)

			if output.String() != expected {
				t.Fatalf("Unexpected json output: '%s' should be '%s'", output.String(), expected)
			}

			return

		case log := <-logs:
			t.Log(log)

		case msg := <-fatal:
			t.Fatal(msg)

		case <-time.After(10 * time.Second):
			t.Fatal("Test has timeout")
		}
	}
}
