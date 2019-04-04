package soba_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/novln/soba"
)

// Test logger name format.
func TestLogger_IsNameValid(t *testing.T) {

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

// Test logger constructor using context.
func TestLogger_New(t *testing.T) {
	conf := &soba.Config{
		Loggers: map[string]soba.ConfigLogger{},
		Appenders: map[string]soba.ConfigAppender{
			"stdout": {
				Type: "console",
			},
		},
		Root: soba.ConfigLogger{
			Level:    "disabled",
			Additive: false,
			Appenders: []string{
				"stdout",
			},
		},
	}

	ctx := context.Background()

	{
		ctx, err := soba.LoadWithConfig(ctx, conf)
		if err != nil {
			t.Fatal(err)
		}

		logger := soba.New(ctx, "views.greetings")
		if logger.Name() != "views.greetings" {
			t.Fatalf("Unexpected logger name: %s should be %s", logger.Name(), "views.greetings")
		}
		if logger.Level() != soba.NoLevel {
			t.Fatalf("Unexpected logger name: %s should be %s", logger.Level(), soba.NoLevel)
		}
	}

	{
		defer func() {
			thr := recover()
			if thr == nil {
				t.Fatal("A panic was expected")
			}
		}()

		logger := soba.New(ctx, "views.greetings")
		_ = logger
	}
}

// Test logger filters on debug level.
func TestLogger_DebugLevel(t *testing.T) {
	appender := NewTestAppender("foobar")
	logger := soba.NewLogger("foobar", soba.DebugLevel, []soba.Appender{appender})

	logger.Debug("Debug level", soba.Int("idx", 1))
	logger.Info("Info level", soba.Int("idx", 2))
	logger.Warn("Warn level", soba.Int("idx", 3))
	logger.Error("Error level", soba.Int("idx", 4))

	if appender.Size() != 4 {
		t.Fatalf("Unexpected number of entries for appender: %d should be %d", appender.Size(), 4)
	}

	expected1 := fmt.Sprint(
		`{"logger":"foobar","level":"debug","message":"Debug level","idx":1}`,
		"\n",
	)
	expected2 := fmt.Sprint(
		`{"logger":"foobar","level":"info","message":"Info level","idx":2}`,
		"\n",
	)
	expected3 := fmt.Sprint(
		`{"logger":"foobar","level":"warning","message":"Warn level","idx":3}`,
		"\n",
	)
	expected4 := fmt.Sprint(
		`{"logger":"foobar","level":"error","message":"Error level","idx":4}`,
		"\n",
	)

	if appender.Log(0) != expected1 {
		t.Fatalf("Unexpected log message #1: '%s' should be '%s'", appender.Log(0), expected1)
	}
	if appender.Log(1) != expected2 {
		t.Fatalf("Unexpected log message #2: '%s' should be '%s'", appender.Log(1), expected2)
	}
	if appender.Log(2) != expected3 {
		t.Fatalf("Unexpected log message #3: '%s' should be '%s'", appender.Log(2), expected3)
	}
	if appender.Log(3) != expected4 {
		t.Fatalf("Unexpected log message #4: '%s' should be '%s'", appender.Log(3), expected4)
	}
}

// Test logger filters on info level.
func TestLogger_InfoLevel(t *testing.T) {
	appender := NewTestAppender("foobar")
	logger := soba.NewLogger("foobar", soba.InfoLevel, []soba.Appender{appender})

	logger.Debug("Debug level", soba.Int("idx", 1))
	logger.Info("Info level", soba.Int("idx", 2))
	logger.Warn("Warn level", soba.Int("idx", 3))
	logger.Error("Error level", soba.Int("idx", 4))

	if appender.Size() != 3 {
		t.Fatalf("Unexpected number of entries for appender: %d should be %d", appender.Size(), 3)
	}

	expected1 := fmt.Sprint(
		`{"logger":"foobar","level":"info","message":"Info level","idx":2}`,
		"\n",
	)
	expected2 := fmt.Sprint(
		`{"logger":"foobar","level":"warning","message":"Warn level","idx":3}`,
		"\n",
	)
	expected3 := fmt.Sprint(
		`{"logger":"foobar","level":"error","message":"Error level","idx":4}`,
		"\n",
	)

	if appender.Log(0) != expected1 {
		t.Fatalf("Unexpected log message #1: '%s' should be '%s'", appender.Log(0), expected1)
	}
	if appender.Log(1) != expected2 {
		t.Fatalf("Unexpected log message #2: '%s' should be '%s'", appender.Log(1), expected2)
	}
	if appender.Log(2) != expected3 {
		t.Fatalf("Unexpected log message #3: '%s' should be '%s'", appender.Log(2), expected3)
	}
}

// Test logger filters on warning level.
func TestLogger_WarnLevel(t *testing.T) {
	appender := NewTestAppender("foobar")
	logger := soba.NewLogger("foobar", soba.WarnLevel, []soba.Appender{appender})

	logger.Debug("Debug level", soba.Int("idx", 1))
	logger.Info("Info level", soba.Int("idx", 2))
	logger.Warn("Warn level", soba.Int("idx", 3))
	logger.Error("Error level", soba.Int("idx", 4))

	if appender.Size() != 2 {
		t.Fatalf("Unexpected number of entries for appender: %d should be %d", appender.Size(), 4)
	}

	expected1 := fmt.Sprint(
		`{"logger":"foobar","level":"warning","message":"Warn level","idx":3}`,
		"\n",
	)
	expected2 := fmt.Sprint(
		`{"logger":"foobar","level":"error","message":"Error level","idx":4}`,
		"\n",
	)

	if appender.Log(0) != expected1 {
		t.Fatalf("Unexpected log message #1: '%s' should be '%s'", appender.Log(0), expected1)
	}
	if appender.Log(1) != expected2 {
		t.Fatalf("Unexpected log message #2: '%s' should be '%s'", appender.Log(1), expected2)
	}
}

// Test logger filters on error level.
func TestLogger_ErrorLevel(t *testing.T) {
	appender := NewTestAppender("foobar")
	logger := soba.NewLogger("foobar", soba.ErrorLevel, []soba.Appender{appender})

	logger.Debug("Debug level", soba.Int("idx", 1))
	logger.Info("Info level", soba.Int("idx", 2))
	logger.Warn("Warn level", soba.Int("idx", 3))
	logger.Error("Error level", soba.Int("idx", 4))

	if appender.Size() != 1 {
		t.Fatalf("Unexpected number of entries for appender: %d should be %d", appender.Size(), 1)
	}

	expected1 := fmt.Sprint(
		`{"logger":"foobar","level":"error","message":"Error level","idx":4}`,
		"\n",
	)

	if appender.Log(0) != expected1 {
		t.Fatalf("Unexpected log message #1: '%s' should be '%s'", appender.Log(0), expected1)
	}
}

// Test logger filters on no level.
func TestLogger_NoLevel(t *testing.T) {
	appender := NewTestAppender("foobar")
	logger := soba.NewLogger("foobar", soba.NoLevel, []soba.Appender{appender})

	logger.Debug("Debug level", soba.Int("idx", 1))
	logger.Info("Info level", soba.Int("idx", 2))
	logger.Warn("Warn level", soba.Int("idx", 3))
	logger.Error("Error level", soba.Int("idx", 4))

	if appender.Size() != 0 {
		t.Fatalf("Unexpected number of entries for appender: %d should be %d", appender.Size(), 0)
	}
}

// Test logger filters on unknown level should behave as no level.
func TestLogger_UnknownLevel(t *testing.T) {
	appender := NewTestAppender("foobar")
	logger := soba.NewLogger("foobar", soba.UnknownLevel, []soba.Appender{appender})

	logger.Debug("Debug level", soba.Int("idx", 1))
	logger.Info("Info level", soba.Int("idx", 2))
	logger.Warn("Warn level", soba.Int("idx", 3))
	logger.Error("Error level", soba.Int("idx", 4))

	if appender.Size() != 0 {
		t.Fatalf("Unexpected number of entries for appender: %d should be %d", appender.Size(), 0)
	}
}

// Test logger with custom fields.
func TestLogger_WithFields(t *testing.T) {
	appender := NewTestAppender("foobar")
	logger := soba.NewLogger("foobar", soba.InfoLevel, []soba.Appender{appender})

	logger = logger.With(soba.String("module", "xyz"), soba.Bool("shared", true))
	logger.Info("Random message 1", soba.Int("id", 1))
	logger.Info("Random message 2", soba.Int("id", 2))
	logger.Info("Random message 3", soba.Int("id", 3))

	if appender.Size() != 3 {
		t.Fatalf("Unexpected number of entries for appender: %d should be %d", appender.Size(), 3)
	}

	expected1 := fmt.Sprint(
		`{"logger":"foobar","level":"info","message":"Random message 1",`,
		`"module":"xyz","shared":true,"id":1}`,
		"\n",
	)
	expected2 := fmt.Sprint(
		`{"logger":"foobar","level":"info","message":"Random message 2",`,
		`"module":"xyz","shared":true,"id":2}`,
		"\n",
	)
	expected3 := fmt.Sprint(
		`{"logger":"foobar","level":"info","message":"Random message 3",`,
		`"module":"xyz","shared":true,"id":3}`,
		"\n",
	)

	if appender.Log(0) != expected1 {
		t.Fatalf("Unexpected log message #1: '%s' should be '%s'", appender.Log(0), expected1)
	}
	if appender.Log(1) != expected2 {
		t.Fatalf("Unexpected log message #2: '%s' should be '%s'", appender.Log(1), expected2)
	}
	if appender.Log(2) != expected3 {
		t.Fatalf("Unexpected log message #3: '%s' should be '%s'", appender.Log(2), expected3)
	}
}
