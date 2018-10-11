package soba_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/novln/soba"
)

// TestAppender is an appender for unit test.
type TestAppender struct {
	name    string
	entries []soba.Entry
}

func (appender *TestAppender) Name() string {
	return appender.name
}

func (appender *TestAppender) Write(entry soba.Entry) {
	appender.entries = append(appender.entries, entry)
}

// NewTestAppender creates a new TestAppender.
func NewTestAppender(name string) *TestAppender {
	return &TestAppender{
		name: name,
	}
}

// Ensure TestAppender implements Appender interface at compile time.
var _ soba.Appender = &TestAppender{}

// Test appender name format.
func TestAppender_IsNameValid(t *testing.T) {

	scenarios := []struct {
		name  string
		valid bool
	}{
		{
			// Scenario #1
			name:  "hello",
			valid: true,
		},
		{
			// Scenario #2
			name:  "hello.world",
			valid: true,
		},
		{
			// Scenario #3
			name:  "Hello.World",
			valid: false,
		},
		{
			// Scenario #4
			name:  "Hello/World",
			valid: false,
		},
		{
			// Scenario #5
			name:  "hello/world",
			valid: false,
		},
		{
			// Scenario #6
			name:  "hello.world0",
			valid: true,
		},
		{
			// Scenario #7
			name:  "hello_world0",
			valid: true,
		},
		{
			// Scenario #8
			name:  "0hello_world0",
			valid: false,
		},
		{
			// Scenario #9
			name:  "_hello_world_",
			valid: false,
		},
		{
			// Scenario #10
			name:  "hello_world_",
			valid: false,
		},
		{
			// Scenario #11
			name:  "hello_world",
			valid: true,
		},
		{
			// Scenario #12
			name:  "xx",
			valid: false,
		},
		{
			// Scenario #13
			name:  "",
			valid: false,
		},
		{
			// Scenario #14
			name:  "foo",
			valid: true,
		},
		{
			// Scenario #15
			name:  "hello-world",
			valid: true,
		},
		{
			// Scenario #16
			name:  "hello-world0",
			valid: true,
		},
	}

	for i, scenario := range scenarios {
		message := fmt.Sprintf("scenario #%d", (i + 1))
		result := soba.IsLoggerNameValid(scenario.name)
		if result != scenario.valid {
			t.Fatalf("Unexpected result for %s: %v should be %v", message, result, scenario.valid)
		}
	}

}

// Test appender write operation.
func TestAppender_Write(t *testing.T) {
	appender := NewTestAppender("foobar")
	before := time.Now()
	entry := soba.NewEntry("foobar.module.asm", soba.InfoLevel, "\\x67")
	defer entry.Flush()
	after := time.Now()

	appender.Write(*entry)

	if len(appender.entries) != 1 {
		t.Fatalf("Unexpected number of entries: %v should be %v", len(appender.entries), 1)
	}
	if appender.entries[0].Name() != "foobar.module.asm" {
		t.Fatalf("Unexpected entry name: %v should be %v", appender.entries[0].Name(), "foobar.module.asm")
	}
	if appender.entries[0].Level() != soba.InfoLevel {
		t.Fatalf("Unexpected entry message: %v should be %v", appender.entries[0].Level(), soba.InfoLevel)
	}
	if appender.entries[0].Message() != "\\x67" {
		t.Fatalf("Unexpected entry message: %v should be %v", appender.entries[0].Message(), "\\x67")
	}
	if appender.entries[0].Unix() < before.Unix() {
		t.Fatalf("Unexpected entry timestamp: %v should be >= %v", appender.entries[0].Unix(), before.Unix())
	}
	if appender.entries[0].Unix() > after.Unix() {
		t.Fatalf("Unexpected entry timestamp: %v should be <= %v", appender.entries[0].Unix(), after.Unix())
	}
}
