package soba_test

import (
	"fmt"
	"strings"
	"testing"

	random "github.com/Pallinder/go-randomdata"

	"github.com/novln/soba"
)

// NewHandler creates a default handler for test and benchmark.
func NewHandler() (soba.Handler, error) {
	return soba.Create(&soba.Config{
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
	})
}

// Benchmark allocation of new Logger from Handler.
func BenchmarkHandler_NewLogger(b *testing.B) {
	handler, err := NewHandler()
	if err != nil {
		b.Fatal(err)
	}

	list := []string{}

	number := random.Number(100, 500)
	for i := 0; i < number; i++ {
		levels := []string{}
		for y := 0; y < 10; y++ {
			if (y % 2) == 0 {
				levels = append(levels, strings.ToLower(random.Noun()))
			} else {
				levels = append(levels, strings.ToLower(random.Adjective()))
			}
		}
		for y := 0; y < 10; y++ {
			if random.Number(0, 2) == 0 {
				name := strings.Join(levels[0:y], ".")
				if soba.IsLoggerNameValid(name) {
					list = append(list, name)
				}
			}
		}
	}

	max := len(list)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {

		// Create a local logger and update it's value from handler to prevent the
		// compiler to eliminate function execution.
		var l soba.Logger

		i := 0
		for pb.Next() {
			i = (i + 1) % max
			l = handler.New(list[i])
		}

		// Store logger instance in a global variable so the compiler cannot eliminate the benchmark.
		// It create a race conditions but it's okay since it's only a benchmark and not a test.
		gl = l

	})
}

// Test hierarchy of loggers.
// nolint: gocyclo
func TestHandler_NewLogger(t *testing.T) {
	apiAppender := NewTestAppender("api-log")
	defer CloseAppender(t, apiAppender)

	dbAppender := NewTestAppender("db-log")
	defer CloseAppender(t, dbAppender)

	authAppender := NewTestAppender("auth-log")
	defer CloseAppender(t, authAppender)

	stdoutAppender := NewTestAppender("stdout")
	defer CloseAppender(t, stdoutAppender)

	err := soba.RegisterAppenders(apiAppender, dbAppender, authAppender, stdoutAppender)
	if err != nil {
		t.Fatalf("Unexpected error: %+v", err)
	}

	handler, err := soba.Create(&soba.Config{
		Root: soba.ConfigLogger{
			Level:    "info",
			Additive: false,
			Appenders: []string{
				"stdout",
			},
		},
		Appenders: map[string]soba.ConfigAppender{
			"console-log": {
				Type: soba.ConsoleAppenderType,
			},
			"file-log": {
				Type: soba.FileAppenderType,
				Path: "/var/lib/soba/file.log",
			},
		},
		Loggers: map[string]soba.ConfigLogger{
			"components.auth": {
				Level:    "debug",
				Additive: true,
				Appenders: []string{
					"auth-log",
				},
			},
			"repositories": {
				Level:    "info",
				Additive: false,
				Appenders: []string{
					"db-log",
				},
			},
			"views": {
				Level:    "warning",
				Additive: false,
				Appenders: []string{
					"api-log",
				},
			},
			"pkg.cache": {
				Level:    "debug",
				Additive: false,
				Appenders: []string{
					"stdout",
				},
			},
			"pkg.core": {
				Level:    "debug",
				Additive: true,
				Appenders: []string{
					"stdout",
				},
			},
			"pkg.proxy": {
				Level:     "debug",
				Additive:  true,
				Appenders: []string{},
			},
		},
	})
	if err != nil {
		t.Fatalf("Unexpected error: %+v", err)
	}

	//
	// Testing repositories loggers
	//
	{

		handler.New("repositories.users").
			Info("User created", soba.String("id", "01CV5FN4JF1STZMYDJWMGQR68W"))
		handler.New("repositories.users").
			Debug("Transaction executed", soba.Bool("commit", true))
		handler.New("repositories.users").
			Info("User created", soba.String("id", "01CV5FP35H929DBQ3KJ6JRST3W"))
		handler.New("repositories.users").
			Debug("Transaction executed", soba.Bool("commit", true))
		handler.New("repositories.organizations").
			Info("Organization created", soba.String("id", "01CV5H2TRFXDHVSTCK5BJVQ1TK"))
		handler.New("repositories.organizations").
			Debug("Transaction executed", soba.Bool("commit", true))

		if dbAppender.Size() != 3 {
			t.Fatalf("Unexpected number of entries for db appender: %d should be %d", dbAppender.Size(), 3)
		}
		if stdoutAppender.Size() != 0 {
			t.Fatalf("Unexpected number of entries for stdout appender: %d should be %d", stdoutAppender.Size(), 0)
		}
		if apiAppender.Size() != 0 {
			t.Fatalf("Unexpected number of entries for api appender: %d should be %d", apiAppender.Size(), 0)
		}
		if authAppender.Size() != 0 {
			t.Fatalf("Unexpected number of entries for auth appender: %d should be %d", authAppender.Size(), 0)
		}

		expected1 := fmt.Sprint(
			`{"logger":"repositories.users","level":"info",`,
			`"message":"User created","id":"01CV5FN4JF1STZMYDJWMGQR68W"}`,
			"\n",
		)
		expected2 := fmt.Sprint(
			`{"logger":"repositories.users","level":"info",`,
			`"message":"User created","id":"01CV5FP35H929DBQ3KJ6JRST3W"}`,
			"\n",
		)
		expected3 := fmt.Sprint(
			`{"logger":"repositories.organizations","level":"info",`,
			`"message":"Organization created","id":"01CV5H2TRFXDHVSTCK5BJVQ1TK"}`,
			"\n",
		)

		if dbAppender.Log(0) != expected1 {
			t.Fatalf("Unexpected log message #1: '%s' should be '%s'", dbAppender.Log(0), expected1)
		}
		if dbAppender.Log(1) != expected2 {
			t.Fatalf("Unexpected log message #2: '%s' should be '%s'", dbAppender.Log(1), expected2)
		}
		if dbAppender.Log(2) != expected3 {
			t.Fatalf("Unexpected log message #3: '%s' should be '%s'", dbAppender.Log(2), expected3)
		}

		dbAppender.Clear()
		stdoutAppender.Clear()

	}

	//
	// Testing auth logger
	//
	{

		handler.New("components.auth").Info("User authentication has failed",
			soba.String("id", "01CZBCXAM0S9VCMWSCAP0K5DQF"), soba.String("status", "failure"))
		handler.New("components.auth").Debug("Increase authentication attempts",
			soba.String("id", "01CZBCXAM0S9VCMWSCAP0K5DQF"), soba.Int("counter", 1))
		handler.New("components.auth").Info("User authentication has failed",
			soba.String("id", "01CZBCXAM0S9VCMWSCAP0K5DQF"), soba.String("status", "failure"))
		handler.New("components.auth").Debug("Increase authentication attempts",
			soba.String("id", "01CZBCXAM0S9VCMWSCAP0K5DQF"), soba.Int("counter", 2))
		handler.New("components.auth").Info("User authentication has failed",
			soba.String("id", "01CZBCXAM0S9VCMWSCAP0K5DQF"), soba.String("status", "failure"))
		handler.New("components.auth").Debug("Increase authentication attempts",
			soba.String("id", "01CZBCXAM0S9VCMWSCAP0K5DQF"), soba.Int("counter", 3))
		handler.New("components.auth").Info("User authentication has succeeded",
			soba.String("id", "01CZBCXAM0S9VCMWSCAP0K5DQF"), soba.String("status", "success"))
		handler.New("components.auth").Debug("Reset authentication attempts",
			soba.String("id", "01CZBCXAM0S9VCMWSCAP0K5DQF"), soba.Int("counter", 0))
		handler.New("components.auth").Info("User authentication has succeeded",
			soba.String("id", "01CZBD1326W8KJWE37WXZ2M6K0"), soba.String("status", "success"))

		if authAppender.Size() != 9 {
			t.Fatalf("Unexpected number of entries for auth appender: %d should be %d", authAppender.Size(), 9)
		}
		if stdoutAppender.Size() != 9 {
			t.Fatalf("Unexpected number of entries for stdout appender: %d should be %d", stdoutAppender.Size(), 9)
		}
		if apiAppender.Size() != 0 {
			t.Fatalf("Unexpected number of entries for api appender: %d should be %d", apiAppender.Size(), 0)
		}
		if dbAppender.Size() != 0 {
			t.Fatalf("Unexpected number of entries for db appender: %d should be %d", dbAppender.Size(), 0)
		}

		expected1 := fmt.Sprint(
			`{"logger":"components.auth","level":"info","message":"User authentication has failed",`,
			`"id":"01CZBCXAM0S9VCMWSCAP0K5DQF","status":"failure"}`,
			"\n",
		)
		expected2 := fmt.Sprint(
			`{"logger":"components.auth","level":"debug","message":"Increase authentication attempts",`,
			`"id":"01CZBCXAM0S9VCMWSCAP0K5DQF","counter":1}`,
			"\n",
		)
		expected3 := fmt.Sprint(
			`{"logger":"components.auth","level":"debug","message":"Increase authentication attempts",`,
			`"id":"01CZBCXAM0S9VCMWSCAP0K5DQF","counter":2}`,
			"\n",
		)
		expected4 := fmt.Sprint(
			`{"logger":"components.auth","level":"debug","message":"Increase authentication attempts",`,
			`"id":"01CZBCXAM0S9VCMWSCAP0K5DQF","counter":3}`,
			"\n",
		)
		expected5 := fmt.Sprint(
			`{"logger":"components.auth","level":"info","message":"User authentication has succeeded",`,
			`"id":"01CZBCXAM0S9VCMWSCAP0K5DQF","status":"success"}`,
			"\n",
		)
		expected6 := fmt.Sprint(
			`{"logger":"components.auth","level":"debug","message":"Reset authentication attempts",`,
			`"id":"01CZBCXAM0S9VCMWSCAP0K5DQF","counter":0}`,
			"\n",
		)
		expected7 := fmt.Sprint(
			`{"logger":"components.auth","level":"info","message":"User authentication has succeeded",`,
			`"id":"01CZBD1326W8KJWE37WXZ2M6K0","status":"success"}`,
			"\n",
		)

		if authAppender.Log(0) != expected1 {
			t.Fatalf("Unexpected log message #1: '%s' should be '%s'", authAppender.Log(0), expected1)
		}
		if stdoutAppender.Log(0) != expected1 {
			t.Fatalf("Unexpected log message #1: '%s' should be '%s'", stdoutAppender.Log(0), expected1)
		}
		if authAppender.Log(1) != expected2 {
			t.Fatalf("Unexpected log message #2: '%s' should be '%s'", authAppender.Log(1), expected2)
		}
		if stdoutAppender.Log(1) != expected2 {
			t.Fatalf("Unexpected log message #2: '%s' should be '%s'", stdoutAppender.Log(1), expected2)
		}
		if authAppender.Log(2) != expected1 {
			t.Fatalf("Unexpected log message #3: '%s' should be '%s'", authAppender.Log(2), expected1)
		}
		if stdoutAppender.Log(2) != expected1 {
			t.Fatalf("Unexpected log message #3: '%s' should be '%s'", stdoutAppender.Log(2), expected1)
		}
		if authAppender.Log(3) != expected3 {
			t.Fatalf("Unexpected log message #4: '%s' should be '%s'", authAppender.Log(3), expected3)
		}
		if stdoutAppender.Log(3) != expected3 {
			t.Fatalf("Unexpected log message #4: '%s' should be '%s'", stdoutAppender.Log(3), expected3)
		}
		if authAppender.Log(4) != expected1 {
			t.Fatalf("Unexpected log message #5: '%s' should be '%s'", authAppender.Log(4), expected1)
		}
		if stdoutAppender.Log(4) != expected1 {
			t.Fatalf("Unexpected log message #5: '%s' should be '%s'", stdoutAppender.Log(4), expected1)
		}
		if authAppender.Log(5) != expected4 {
			t.Fatalf("Unexpected log message #6: '%s' should be '%s'", authAppender.Log(5), expected4)
		}
		if stdoutAppender.Log(5) != expected4 {
			t.Fatalf("Unexpected log message #6: '%s' should be '%s'", stdoutAppender.Log(5), expected4)
		}
		if authAppender.Log(6) != expected5 {
			t.Fatalf("Unexpected log message #7: '%s' should be '%s'", authAppender.Log(6), expected5)
		}
		if stdoutAppender.Log(6) != expected5 {
			t.Fatalf("Unexpected log message #7: '%s' should be '%s'", stdoutAppender.Log(6), expected5)
		}
		if authAppender.Log(7) != expected6 {
			t.Fatalf("Unexpected log message #8: '%s' should be '%s'", authAppender.Log(7), expected6)
		}
		if stdoutAppender.Log(7) != expected6 {
			t.Fatalf("Unexpected log message #8: '%s' should be '%s'", stdoutAppender.Log(7), expected6)
		}
		if authAppender.Log(8) != expected7 {
			t.Fatalf("Unexpected log message #9: '%s' should be '%s'", authAppender.Log(8), expected7)
		}
		if stdoutAppender.Log(8) != expected7 {
			t.Fatalf("Unexpected log message #9: '%s' should be '%s'", stdoutAppender.Log(8), expected7)
		}

		authAppender.Clear()
		stdoutAppender.Clear()

	}

	//
	// Testing views logger
	//
	{
		handler.New("views.users").Info("User created",
			soba.String("id", "01CV5FN4JF1STZMYDJWMGQR68W"),
			soba.String("action", "create"), soba.String("status", "success"),
		)
		handler.New("views.users").Info("User created",
			soba.String("id", "01CV5FP35H929DBQ3KJ6JRST3W"),
			soba.String("action", "create"), soba.String("status", "success"),
		)
		handler.New("views.organizations").Info("Organization created",
			soba.String("id", "01CV5H2TRFXDHVSTCK5BJVQ1TK"),
			soba.String("action", "create"), soba.String("status", "success"),
		)
		handler.New("views.users").Warn("User authentication has failed",
			soba.String("id", "01CZBCXAM0S9VCMWSCAP0K5DQF"),
			soba.String("action", "login"), soba.String("status", "failure"),
		)
		handler.New("views.users").Warn("User authentication has failed",
			soba.String("id", "01CZBCXAM0S9VCMWSCAP0K5DQF"),
			soba.String("action", "login"), soba.String("status", "failure"),
		)
		handler.New("views.users").Warn("User authentication has failed",
			soba.String("id", "01CZBCXAM0S9VCMWSCAP0K5DQF"),
			soba.String("action", "login"), soba.String("status", "failure"),
		)
		handler.New("views.users").Info("User authentication has succeeded",
			soba.String("id", "01CZBCXAM0S9VCMWSCAP0K5DQF"),
			soba.String("action", "login"), soba.String("status", "success"),
		)

		if apiAppender.Size() != 3 {
			t.Fatalf("Unexpected number of entries for api appender: %d should be %d", apiAppender.Size(), 3)
		}
		if stdoutAppender.Size() != 0 {
			t.Fatalf("Unexpected number of entries for stdout appender: %d should be %d", stdoutAppender.Size(), 0)
		}
		if authAppender.Size() != 0 {
			t.Fatalf("Unexpected number of entries for auth appender: %d should be %d", authAppender.Size(), 0)
		}
		if dbAppender.Size() != 0 {
			t.Fatalf("Unexpected number of entries for db appender: %d should be %d", dbAppender.Size(), 0)
		}

		expected1 := fmt.Sprint(
			`{"logger":"views.users","level":"warning","message":"User authentication has failed",`,
			`"id":"01CZBCXAM0S9VCMWSCAP0K5DQF","action":"login","status":"failure"}`,
			"\n",
		)
		expected2 := fmt.Sprint(
			`{"logger":"views.users","level":"warning","message":"User authentication has failed",`,
			`"id":"01CZBCXAM0S9VCMWSCAP0K5DQF","action":"login","status":"failure"}`,
			"\n",
		)
		expected3 := fmt.Sprint(
			`{"logger":"views.users","level":"warning","message":"User authentication has failed",`,
			`"id":"01CZBCXAM0S9VCMWSCAP0K5DQF","action":"login","status":"failure"}`,
			"\n",
		)

		if apiAppender.Log(0) != expected1 {
			t.Fatalf("Unexpected log message #1: '%s' should be '%s'", apiAppender.Log(0), expected1)
		}
		if apiAppender.Log(1) != expected2 {
			t.Fatalf("Unexpected log message #2: '%s' should be '%s'", apiAppender.Log(1), expected2)
		}
		if apiAppender.Log(2) != expected3 {
			t.Fatalf("Unexpected log message #2: '%s' should be '%s'", apiAppender.Log(2), expected3)
		}

		apiAppender.Clear()
		stdoutAppender.Clear()

	}

	//
	// Testing pkg.cache logger
	//
	{

		handler.New("pkg.cache").Debug("Cache miss", soba.String("key", "foobar"), soba.Bool("hit", false))
		handler.New("pkg.cache").Debug("Cache hit", soba.String("key", "foobar"), soba.Bool("hit", true))

		if stdoutAppender.Size() != 2 {
			t.Fatalf("Unexpected number of entries for stdout appender: %d should be %d", stdoutAppender.Size(), 2)
		}
		if apiAppender.Size() != 0 {
			t.Fatalf("Unexpected number of entries for api appender: %d should be %d", apiAppender.Size(), 0)
		}
		if authAppender.Size() != 0 {
			t.Fatalf("Unexpected number of entries for auth appender: %d should be %d", authAppender.Size(), 0)
		}
		if dbAppender.Size() != 0 {
			t.Fatalf("Unexpected number of entries for db appender: %d should be %d", dbAppender.Size(), 0)
		}

		expected1 := fmt.Sprint(
			`{"logger":"pkg.cache","level":"debug","message":"Cache miss",`,
			`"key":"foobar","hit":false}`,
			"\n",
		)
		expected2 := fmt.Sprint(
			`{"logger":"pkg.cache","level":"debug","message":"Cache hit",`,
			`"key":"foobar","hit":true}`,
			"\n",
		)

		if stdoutAppender.Log(0) != expected1 {
			t.Fatalf("Unexpected log message #1: '%s' should be '%s'", stdoutAppender.Log(0), expected1)
		}
		if stdoutAppender.Log(1) != expected2 {
			t.Fatalf("Unexpected log message #1: '%s' should be '%s'", stdoutAppender.Log(1), expected2)
		}

		stdoutAppender.Clear()

	}

	//
	// Testing pkg.core logger
	//
	{

		handler.New("pkg.core").Debug("Account closed", soba.String("id", "01D7JJEQASQXNZP5C6DFT1J7RN"))
		handler.New("pkg.core").Debug("Account closed", soba.String("id", "01D7JJEQASQXNZP5C6DFT1J7RN"))

		if stdoutAppender.Size() != 2 {
			t.Fatalf("Unexpected number of entries for stdout appender: %d should be %d", stdoutAppender.Size(), 2)
		}
		if apiAppender.Size() != 0 {
			t.Fatalf("Unexpected number of entries for api appender: %d should be %d", apiAppender.Size(), 0)
		}
		if authAppender.Size() != 0 {
			t.Fatalf("Unexpected number of entries for auth appender: %d should be %d", authAppender.Size(), 0)
		}
		if dbAppender.Size() != 0 {
			t.Fatalf("Unexpected number of entries for db appender: %d should be %d", dbAppender.Size(), 0)
		}

		expected1 := fmt.Sprint(
			`{"logger":"pkg.core","level":"debug","message":"Account closed",`,
			`"id":"01D7JJEQASQXNZP5C6DFT1J7RN"}`,
			"\n",
		)
		expected2 := fmt.Sprint(
			`{"logger":"pkg.core","level":"debug","message":"Account closed",`,
			`"id":"01D7JJEQASQXNZP5C6DFT1J7RN"}`,
			"\n",
		)

		if stdoutAppender.Log(0) != expected1 {
			t.Fatalf("Unexpected log message #1: '%s' should be '%s'", stdoutAppender.Log(0), expected1)
		}
		if stdoutAppender.Log(1) != expected2 {
			t.Fatalf("Unexpected log message #1: '%s' should be '%s'", stdoutAppender.Log(1), expected2)
		}

		stdoutAppender.Clear()

	}

	//
	// Testing pkg.proxy logger
	//
	{

		handler.New("pkg.proxy").Debug("Request forwarded", soba.String("trace_id", "b2711a28753adb247822ad9f27a3cc"))
		handler.New("pkg.proxy").Debug("Request forwarded", soba.String("trace_id", "953bfedc8bfcc3ac6559c2920c0e37"))

		if stdoutAppender.Size() != 2 {
			t.Fatalf("Unexpected number of entries for stdout appender: %d should be %d", stdoutAppender.Size(), 2)
		}
		if apiAppender.Size() != 0 {
			t.Fatalf("Unexpected number of entries for api appender: %d should be %d", apiAppender.Size(), 0)
		}
		if authAppender.Size() != 0 {
			t.Fatalf("Unexpected number of entries for auth appender: %d should be %d", authAppender.Size(), 0)
		}
		if dbAppender.Size() != 0 {
			t.Fatalf("Unexpected number of entries for db appender: %d should be %d", dbAppender.Size(), 0)
		}

		expected1 := fmt.Sprint(
			`{"logger":"pkg.proxy","level":"debug","message":"Request forwarded",`,
			`"trace_id":"b2711a28753adb247822ad9f27a3cc"}`,
			"\n",
		)
		expected2 := fmt.Sprint(
			`{"logger":"pkg.proxy","level":"debug","message":"Request forwarded",`,
			`"trace_id":"953bfedc8bfcc3ac6559c2920c0e37"}`,
			"\n",
		)

		if stdoutAppender.Log(0) != expected1 {
			t.Fatalf("Unexpected log message #1: '%s' should be '%s'", stdoutAppender.Log(0), expected1)
		}
		if stdoutAppender.Log(1) != expected2 {
			t.Fatalf("Unexpected log message #1: '%s' should be '%s'", stdoutAppender.Log(1), expected2)
		}

		stdoutAppender.Clear()

	}

	//
	// Testing random logger
	//
	{

		handler.New("services.webhooks").Info("User email confirmed",
			soba.String("id", "01CV5FN4JF1STZMYDJWMGQR68W"),
			soba.String("webhook", "email.confirm"), soba.String("status", "success"),
		)
		handler.New("services.webhooks").Debug("Webhook received",
			soba.String("id", "01CV5FN4JF1STZMYDJWMGQR68W"),
			soba.String("webhook", "email.confirm"), soba.String("status", "received"),
		)

		if stdoutAppender.Size() != 1 {
			t.Fatalf("Unexpected number of entries for stdout appender: %d should be %d", dbAppender.Size(), 1)
		}
		if apiAppender.Size() != 0 {
			t.Fatalf("Unexpected number of entries for api appender: %d should be %d", apiAppender.Size(), 0)
		}
		if authAppender.Size() != 0 {
			t.Fatalf("Unexpected number of entries for auth appender: %d should be %d", authAppender.Size(), 0)
		}
		if dbAppender.Size() != 0 {
			t.Fatalf("Unexpected number of entries for db appender: %d should be %d", dbAppender.Size(), 0)
		}

		expected1 := fmt.Sprint(
			`{"logger":"services.webhooks","level":"info","message":"User email confirmed",`,
			`"id":"01CV5FN4JF1STZMYDJWMGQR68W","webhook":"email.confirm","status":"success"}`,
			"\n",
		)
		if stdoutAppender.Log(0) != expected1 {
			t.Fatalf("Unexpected log message #1: '%s' should be '%s'", stdoutAppender.Log(0), expected1)
		}

		stdoutAppender.Clear()

	}

}
