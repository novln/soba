package soba_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/novln/soba"
	"github.com/novln/soba/encoder/json"
)

// TestAppender is an appender for unit test.
type TestAppender struct {
	name    string
	entries []string
	times   []time.Time
}

func (appender *TestAppender) Name() string {
	return appender.name
}

func (appender *TestAppender) Write(entry *soba.Entry) {
	encoder := json.NewEncoder()
	defer encoder.Close()

	buffer := encoder.Encode(func(encoder soba.Encoder) {
		encoder.AddString("logger", entry.Name())
		encoder.AddStringer("level", entry.Level())
		encoder.AddString("message", entry.Message())
		for _, field := range entry.Fields() {
			field.Write(encoder)
		}
	})

	appender.entries = append(appender.entries, string(buffer))
	appender.times = append(appender.times, time.Unix(entry.Unix(), 0).UTC())
}

func (appender *TestAppender) Clear() {
	appender.entries = []string{}
	appender.times = []time.Time{}
}

func (appender *TestAppender) Size() int {
	return len(appender.entries)
}

func (appender *TestAppender) Log(index int) string {
	return appender.entries[index]
}

func (appender *TestAppender) Time(index int) time.Time {
	return appender.times[index]
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

// Test appender write operation for other tests.
// nolint: gocyclo
func TestAppender_Write(t *testing.T) {
	appender := NewTestAppender("foobar")

	{
		before := time.Now()
		entry := soba.NewEntry("foobar.module.asm", soba.WarnLevel, "Invalid opcode", []soba.Field{
			soba.Binary("opcode", []byte{0x67}),
			soba.String("module", "bootloader"),
		})
		defer entry.Flush()
		after := time.Now()

		appender.Write(entry)

		expected := fmt.Sprint(
			`{"logger":"foobar.module.asm","level":"warning",`,
			`"message":"Invalid opcode","opcode":"Zw==","module":"bootloader"}`,
			"\n",
		)

		if len(appender.entries) != 1 {
			t.Fatalf("Unexpected number of entries: %d should be %d", len(appender.entries), 1)
		}
		if len(appender.times) != len(appender.entries) {
			t.Fatalf("Unexpected number of entries timestamp: %d should be %d",
				len(appender.times), len(appender.entries))
		}
		if appender.entries[0] != expected {
			t.Fatalf("Unexpected entry #1: '%s' should be '%s'", appender.entries[0], expected)
		}
		if appender.Log(0) != appender.entries[0] {
			t.Fatalf("Unexpected entry #1: '%s' should be '%s'", appender.Log(0), appender.entries[0])
		}
		if appender.times[0].Unix() < before.Unix() {
			t.Fatalf("Unexpected entry timestamp: %d should be greater than or equals to %d",
				appender.times[0].Unix(), before.Unix())
		}
		if appender.times[0].Unix() > after.Unix() {
			t.Fatalf("Unexpected entry timestamp: %d should be less than or equals to %d",
				appender.times[0].Unix(), after.Unix())
		}
		if appender.Time(0) != appender.times[0] {
			t.Fatalf("Unexpected entry timestamp: %d should be equals to %d",
				appender.Time(0).Unix(), appender.times[0].Unix())
		}
	}
	{
		before := time.Now()
		entry := soba.NewEntry("foobar.module.asm", soba.DebugLevel, "Jump stack", []soba.Field{
			soba.Uint64("from", 0x23456F34),
			soba.Uint64("to", 0x6723F4AB),
			soba.String("module", "cryptofs"),
		})
		defer entry.Flush()
		after := time.Now()

		appender.Write(entry)

		expected := fmt.Sprint(
			`{"logger":"foobar.module.asm","level":"debug",`,
			`"message":"Jump stack","from":591753012,"to":1730409643,"module":"cryptofs"}`,
			"\n",
		)

		if len(appender.entries) != 2 {
			t.Fatalf("Unexpected number of entries: %d should be %d", len(appender.entries), 2)
		}
		if len(appender.times) != len(appender.entries) {
			t.Fatalf("Unexpected number of entries timestamp: %d should be %d",
				len(appender.times), len(appender.entries))
		}
		if appender.entries[1] != expected {
			t.Fatalf("Unexpected entry #2: '%s' should be '%s'", appender.entries[1], expected)
		}
		if appender.Log(1) != appender.entries[1] {
			t.Fatalf("Unexpected entry #1: '%s' should be '%s'", appender.Log(1), appender.entries[1])
		}
		if appender.times[1].Unix() < before.Unix() {
			t.Fatalf("Unexpected entry timestamp: %d should be greater than or equals to %d",
				appender.times[1].Unix(), before.Unix())
		}
		if appender.times[1].Unix() > after.Unix() {
			t.Fatalf("Unexpected entry timestamp: %d should be less than or equals to %d",
				appender.times[1].Unix(), after.Unix())
		}
		if appender.Time(1) != appender.times[1] {
			t.Fatalf("Unexpected entry timestamp: %d should be equals to %d",
				appender.Time(1).Unix(), appender.times[1].Unix())
		}
	}
}

// Test appender name definition for other tests.
func TestAppender_Name(t *testing.T) {
	appender := NewTestAppender("foobar")
	expected := "foobar"

	if appender.Name() != expected {
		t.Fatalf("Unexpected appender name: %s should be %s", appender.Name(), expected)
	}
}
