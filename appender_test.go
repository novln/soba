package soba_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/novln/soba"
	"github.com/novln/soba/encoder/json"
)

// TestAppender is an appender for test.
type TestAppender struct {
	name    string
	entries []string
	times   []time.Time
}

func (appender *TestAppender) Name() string {
	return appender.name
}

func (TestAppender) Close() error {
	return nil
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

func CloseAppender(t *testing.T, appender soba.Appender) {
	err := appender.Close()
	if err != nil {
		t.Fatalf("Unexpected error: %+v", err)
	}
}

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
	defer CloseAppender(t, appender)

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
	expected := "foobar"
	appender := NewTestAppender(expected)
	defer CloseAppender(t, appender)

	if appender.Name() != expected {
		t.Fatalf("Unexpected appender name: %s should be %s", appender.Name(), expected)
	}
}

// Test console appender constructor with invalid type.
func TestAppender_InvalidNew(t *testing.T) {
	name := "invalid"
	appender, err := soba.NewAppender(name, soba.ConfigAppender{
		Type: "logstash",
	})
	if err == nil {
		CloseAppender(t, appender)
		t.Fatalf(`An error was expected for appender "%s" (invalid type)`, name)
	}
}

// Test console appender constructor.
func TestAppender_ConsoleNew(t *testing.T) {
	name := "console1"
	appender, err := soba.NewAppender(name, soba.ConfigAppender{
		Type: soba.ConsoleAppenderType,
		Path: "/var/log/output.log",
	})
	if err == nil {
		CloseAppender(t, appender)
		t.Fatalf(`An error was expected for appender "%s" (invalid path)`, name)
	}

	name = "console2"
	appender, err = soba.NewAppender(name, soba.ConfigAppender{
		Type:   soba.ConsoleAppenderType,
		Backup: true,
	})
	if err == nil {
		CloseAppender(t, appender)
		t.Fatalf(`An error was expected for appender "%s" (invalid backup)`, name)
	}

	name = "console3"
	appender, err = soba.NewAppender(name, soba.ConfigAppender{
		Type:     soba.ConsoleAppenderType,
		MaxBytes: 60,
	})
	if err == nil {
		CloseAppender(t, appender)
		t.Fatalf(`An error was expected for appender "%s" (invalid max bytes)`, name)
	}

	name = "Console$"
	appender, err = soba.NewAppender(name, soba.ConfigAppender{
		Type: soba.ConsoleAppenderType,
	})
	if err == nil {
		CloseAppender(t, appender)
		t.Fatalf(`An error was expected for appender "%s" (invalid name)`, name)
	}

	name = "console4"
	appender, err = soba.NewAppender(name, soba.ConfigAppender{
		Type: soba.ConsoleAppenderType,
	})
	if err != nil {
		t.Fatalf(`Unexpected error for appender "%s": %+v`, name, err)
	}
	if appender.Name() != name {
		t.Fatalf("Unexpected appender name: %s should be %s", appender.Name(), name)
	}
	CloseAppender(t, appender)
}

// Test console appender write behavior.
func TestAppender_ConsoleWrite(t *testing.T) {
	buffer := &bytes.Buffer{}
	name := "console"

	appender := soba.NewConsoleAppender(name, buffer)
	defer CloseAppender(t, appender)

	entry1 := soba.NewEntry("foobar.module.asm", soba.WarnLevel, "Invalid opcode", []soba.Field{
		soba.Binary("opcode", []byte{0x67}),
		soba.String("module", "bootloader"),
	})
	defer entry1.Flush()

	entry2 := soba.NewEntry("foobar.module.asm", soba.DebugLevel, "Jump stack", []soba.Field{
		soba.Uint64("from", 0x23456F34),
		soba.Uint64("to", 0x6723F4AB),
		soba.String("module", "cryptofs"),
	})
	defer entry2.Flush()

	appender.Write(entry1)
	appender.Write(entry2)

	output := strings.TrimSpace(buffer.String())
	lines := strings.Split(output, "\n")
	if len(lines) != 2 {
		t.Fatalf("Unexpected number of entries: %d should be %d", len(lines), 2)
	}

	checkContains := func(line string, key string) {
		if !strings.Contains(line, key) {
			t.Fatalf("Expect '%s' to be included in: %s", key, line)
		}
	}

	checkContains(lines[0], `"logger":"foobar.module.asm"`)
	checkContains(lines[0], `"level":"warning"`)
	checkContains(lines[0], `"message":"Invalid opcode"`)
	checkContains(lines[0], `"opcode":"Zw=="`)
	checkContains(lines[0], `"module":"bootloader"`)

	checkContains(lines[1], `"logger":"foobar.module.asm"`)
	checkContains(lines[1], `"level":"debug"`)
	checkContains(lines[1], `"message":"Jump stack"`)
	checkContains(lines[1], `"from":591753012`)
	checkContains(lines[1], `"to":1730409643`)
	checkContains(lines[1], `"module":"cryptofs"`)
}

// Test file appender constructor.
func TestAppender_FileNew(t *testing.T) {
	path1 := ""
	path2 := "/output.log"
	path3 := "testdata/logs/file.log"

	name := "file1"
	appender, err := soba.NewAppender(name, soba.ConfigAppender{
		Type:     soba.FileAppenderType,
		Path:     path1,
		Backup:   false,
		MaxBytes: 0,
	})
	if err == nil {
		CloseAppender(t, appender)
		t.Fatalf(`An error was expected for appender "%s" (invalid path)`, name)
	}

	name = "file2"
	appender, err = soba.NewAppender(name, soba.ConfigAppender{
		Type:     soba.FileAppenderType,
		Path:     path2,
		Backup:   false,
		MaxBytes: 0,
	})
	if err == nil {
		CloseAppender(t, appender)
		t.Fatalf(`An error was expected for appender "%s" (cannot create file)`, name)
	}

	name = "file3"
	appender, err = soba.NewAppender(name, soba.ConfigAppender{
		Type:     soba.FileAppenderType,
		Path:     path3,
		Backup:   false,
		MaxBytes: -120,
	})
	if err == nil {
		CloseAppender(t, appender)
		t.Fatalf(`An error was expected for appender "%s" (invalid max bytes)`, name)
	}

	name = "File$"
	appender, err = soba.NewAppender(name, soba.ConfigAppender{
		Type:     soba.FileAppenderType,
		Path:     path3,
		Backup:   false,
		MaxBytes: 0,
	})
	if err == nil {
		CloseAppender(t, appender)
		t.Fatalf(`An error was expected for appender "%s" (invalid name)`, name)
	}

	name = "file4"
	appender, err = soba.NewAppender(name, soba.ConfigAppender{
		Type:     soba.FileAppenderType,
		Path:     path3,
		Backup:   false,
		MaxBytes: 0,
	})
	if err != nil {
		t.Fatalf(`Unexpected error for appender "%s": %+v`, name, err)
	}
	if appender.Name() != name {
		t.Fatalf("Unexpected appender name: %s should be %s", appender.Name(), name)
	}
	CloseAppender(t, appender)
}

// Test file appender write behavior.
// nolint: gocyclo
func TestAppender_FileWrite(t *testing.T) {
	name := "file"
	path1 := "testdata/logs/output.log"
	path2 := "testdata/logs/output.log.1"
	path3 := "testdata/logs/output.log.2"
	path4 := "testdata/logs/output.log-"

	entry1 := soba.NewEntry("foobar.module.asm", soba.WarnLevel, "Invalid opcode", []soba.Field{
		soba.Binary("opcode", []byte{0x67}),
		soba.String("module", "bootloader"),
	})
	defer entry1.Flush()

	entry2 := soba.NewEntry("foobar.module.asm", soba.DebugLevel, "Jump stack", []soba.Field{
		soba.Uint64("from", 0x23456F34),
		soba.Uint64("to", 0x6723F4AB),
		soba.String("module", "cryptofs"),
	})
	defer entry2.Flush()

	deleteFiles := func() {
		err := os.Remove(path1)
		if err != nil && !os.IsNotExist(err) {
			t.Fatalf("Unexpected error: %+v", err)
		}

		err = os.Remove(path2)
		if err != nil && !os.IsNotExist(err) {
			t.Fatalf("Unexpected error: %+v", err)
		}

		err = os.Remove(path3)
		if err != nil && !os.IsNotExist(err) {
			t.Fatalf("Unexpected error: %+v", err)
		}

		err = os.Remove(path4)
		if err != nil && !os.IsNotExist(err) {
			t.Fatalf("Unexpected error: %+v", err)
		}
	}

	getFileContent := func(path string) []byte {
		buffer, err := ioutil.ReadFile(path)
		if err != nil {
			t.Fatalf("Unexpected error: %+v", err)
		}
		if len(buffer) == 0 {
			t.Fatalf("Unexpected empty file: %s", path1)
		}
		return buffer
	}

	getAppender := func(path string, backup bool, maxBytes int64) soba.Appender {
		appender, err := soba.NewFileAppender(name, path, backup, maxBytes)
		if err != nil {
			t.Fatalf("Unexpected error: %+v", err)
		}
		return appender
	}

	checkContains := func(line string, key string) {
		if !strings.Contains(line, key) {
			t.Fatalf("Expect '%s' to be included in: %s", key, line)
		}
	}

	checkFileA := func(path string) {
		buffer := getFileContent(path)
		output := strings.TrimSpace(string(buffer))
		lines := strings.Split(output, "\n")
		if len(lines) != 2 {
			t.Fatalf("Unexpected number of entries: %d should be %d", len(lines), 2)
		}

		checkContains(lines[0], `"logger":"foobar.module.asm"`)
		checkContains(lines[0], `"level":"warning"`)
		checkContains(lines[0], `"message":"Invalid opcode"`)
		checkContains(lines[0], `"opcode":"Zw=="`)
		checkContains(lines[0], `"module":"bootloader"`)

		checkContains(lines[1], `"logger":"foobar.module.asm"`)
		checkContains(lines[1], `"level":"debug"`)
		checkContains(lines[1], `"message":"Jump stack"`)
		checkContains(lines[1], `"from":591753012`)
		checkContains(lines[1], `"to":1730409643`)
		checkContains(lines[1], `"module":"cryptofs"`)
	}

	checkFileB := func(path string) {
		buffer := getFileContent(path)
		output := strings.TrimSpace(string(buffer))
		lines := strings.Split(output, "\n")
		if len(lines) != 1 {
			t.Fatalf("Unexpected number of entries: %d should be %d", len(lines), 1)
		}

		checkContains(lines[0], `"logger":"foobar.module.asm"`)
		checkContains(lines[0], `"level":"warning"`)
		checkContains(lines[0], `"message":"Invalid opcode"`)
		checkContains(lines[0], `"opcode":"Zw=="`)
		checkContains(lines[0], `"module":"bootloader"`)
	}

	checkFileC := func(path string) {
		buffer := getFileContent(path)
		output := strings.TrimSpace(string(buffer))
		lines := strings.Split(output, "\n")
		if len(lines) != 7 {
			t.Fatalf("Unexpected number of entries: %d should be %d", len(lines), 7)
		}

		for i, line := range lines {
			if (i % 2) == 0 {
				checkContains(line, `"logger":"foobar.module.asm"`)
				checkContains(line, `"level":"warning"`)
				checkContains(line, `"message":"Invalid opcode"`)
				checkContains(line, `"opcode":"Zw=="`)
				checkContains(line, `"module":"bootloader"`)
			} else {
				checkContains(line, `"logger":"foobar.module.asm"`)
				checkContains(line, `"level":"debug"`)
				checkContains(line, `"message":"Jump stack"`)
				checkContains(line, `"from":591753012`)
				checkContains(line, `"to":1730409643`)
				checkContains(line, `"module":"cryptofs"`)
			}
		}
	}

	checkFileEmpty := func(path string) {
		buffer, err := ioutil.ReadFile(path)
		if err != nil && !os.IsNotExist(err) {
			t.Fatalf("Unexpected error: %+v", err)
		}
		if len(buffer) != 0 {
			t.Fatalf("An empty file was expected: %s", string(buffer))
		}
	}

	{
		// Test appender with a backup enabled and a limit on file size.
		deleteFiles()

		appender := getAppender(path1, true, 400)
		defer CloseAppender(t, appender)

		appender.Write(entry1)
		appender.Write(entry2)

		checkFileA(path1)
		checkFileEmpty(path2)
		checkFileEmpty(path3)
		checkFileEmpty(path4)

		// Write an entry that execute two file rotation.
		appender.Write(entry1)
		appender.Write(entry2)
		appender.Write(entry1)

		checkFileB(path1)
		checkFileA(path2)
		checkFileA(path3)
		checkFileEmpty(path4)
	}
	{
		// Test appender with a backup disabled and a limit on file size.
		deleteFiles()

		appender := getAppender(path1, false, 400)
		defer CloseAppender(t, appender)

		appender.Write(entry1)
		appender.Write(entry2)

		checkFileA(path1)
		checkFileEmpty(path2)
		checkFileEmpty(path3)
		checkFileEmpty(path4)

		// Write an entry that execute two file rotation.
		appender.Write(entry1)
		appender.Write(entry2)
		appender.Write(entry1)

		checkFileB(path1)
		checkFileA(path4)
		checkFileEmpty(path2)
		checkFileEmpty(path3)
	}
	{
		// Test appender with a backup disabled and no limit on file size.
		deleteFiles()

		appender := getAppender(path1, false, 0)
		defer CloseAppender(t, appender)

		appender.Write(entry1)
		appender.Write(entry2)

		checkFileA(path1)
		checkFileEmpty(path2)
		checkFileEmpty(path3)
		checkFileEmpty(path4)

		// Write an entry that execute no file rotation.
		appender.Write(entry1)
		appender.Write(entry2)
		appender.Write(entry1)
		appender.Write(entry2)
		appender.Write(entry1)

		checkFileC(path1)
		checkFileEmpty(path2)
		checkFileEmpty(path3)
		checkFileEmpty(path4)
	}

	// Cleanup
	deleteFiles()
}
