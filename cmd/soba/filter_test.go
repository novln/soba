package main

import (
	"fmt"
	"testing"
)

func TestFilter_Validator(t *testing.T) {
	testCases := []struct {
		input    string
		filtered bool
	}{
		{
			input:    `Lorem ipsum dolor sit amet, consectetur adipiscing elit.`,
			filtered: true,
		},
		{
			input:    `Foobar`,
			filtered: true,
		},
		{
			input:    `{]`,
			filtered: true,
		},
		{
			input:    `:(){ :|:& };:`,
			filtered: true,
		},
		{
			input:    `logger: prog.go:15: Hello world!`,
			filtered: true,
		},
		{
			input:    `{"key":"val}`,
			filtered: true,
		},
		{
			input:    `{"key":"val",}`,
			filtered: true,
		},
		{
			input:    `{[]}`,
			filtered: true,
		},
		{
			input:    `null`,
			filtered: false,
		},
		{
			input:    `true`,
			filtered: false,
		},
		{
			input:    `false`,
			filtered: false,
		},
		{
			input:    `"foobar"`,
			filtered: false,
		},
		{
			input:    `10`,
			filtered: false,
		},
		{
			input:    `12`,
			filtered: false,
		},
		{
			input:    `-43`,
			filtered: false,
		},
		{
			input:    `-200`,
			filtered: false,
		},
		{
			input:    `1.34`,
			filtered: false,
		},
		{
			input:    `-6.141759`,
			filtered: false,
		},
		{
			input:    `{"key":"val"}`,
			filtered: false,
		},
		{
			input:    `{"key":"foobar"}`,
			filtered: false,
		},
		{
			input:    `{"user":42,"first_name":"Allan","last_name":"GUICHARD","age":99}`,
			filtered: false,
		},
		{
			input: fmt.Sprint(
				`{"user":{"id":42,"identity":{"first_name":"Allan",`,
				`"last_name":"GUICHARD"},"deleted_at":null,"is_admin":true}}`,
			),
			filtered: false,
		},
		{
			input:    `{"keys":["k\\\"0","k\\\"1","k\\\"2"]}`,
			filtered: false,
		},
		{
			input:    `[1,2,3,4,5,6,7]`,
			filtered: false,
		},
		{
			input:    `[{"key":42},{"key":22},{"key":68},{"key":6}]`,
			filtered: false,
		},
	}

	for _, testCase := range testCases {
		output := Filter([]byte(testCase.input), nil)

		if testCase.filtered {
			if len(output) != 0 {
				t.Fatalf("Unexpected json output: '%s'", string(output))
			}
		} else {
			if len(output) == 0 {
				t.Fatalf("Unexpected json filter: '%s'", testCase.input)
			}
			if string(output) != testCase.input {
				t.Fatalf("Unexpected json output: '%s' should be '%s'", string(output), testCase.input)
			}
		}
	}
}

func TestFilter_Loggers(t *testing.T) {
	testCases := []struct {
		input    string
		rules    FilterAction
		filtered bool
	}{
		{
			input: fmt.Sprint(
				`{"logger":"repositories.users","level":"info",`,
				`"message":"User created","id":"01CV5FN4JF1STZMYDJWMGQR68W"}`,
			),
			rules: FilterAction{
				Exclude: []string{},
				Include: []string{
					"repositories.users",
				},
			},
			filtered: false,
		},
		{
			input: fmt.Sprint(
				`{"logger":"repositories.users","level":"info",`,
				`"message":"User created","id":"01CV5FN4JF1STZMYDJWMGQR68W"}`,
			),
			rules: FilterAction{
				Exclude: []string{
					"repositories.users",
				},
				Include: []string{},
			},
			filtered: true,
		},
		{
			input: fmt.Sprint(
				`{"logger":"repositories.users","level":"info",`,
				`"message":"User created","id":"01CV5FN4JF1STZMYDJWMGQR68W"}`,
			),
			rules: FilterAction{
				Exclude: []string{
					"repositories.organizations",
				},
				Include: []string{},
			},
			filtered: false,
		},
		{
			input: fmt.Sprint(
				`{"logger":"repositories.users","level":"info",`,
				`"message":"User created","id":"01CV5FN4JF1STZMYDJWMGQR68W"}`,
			),
			rules: FilterAction{
				Exclude: []string{},
				Include: []string{
					"repositories.organizations",
				},
			},
			filtered: true,
		},
		{
			input: fmt.Sprint(
				`{"logger":"repositories.users","level":"info",`,
				`"message":"User created","id":"01CV5FN4JF1STZMYDJWMGQR68W"}`,
			),
			rules: FilterAction{
				Exclude: []string{
					"repositories",
				},
				Include: []string{
					"repositories.users",
				},
			},
			filtered: false,
		},
		{
			input: fmt.Sprint(
				`{"logger":"repositories.users","level":"info",`,
				`"message":"User created","id":"01CV5FN4JF1STZMYDJWMGQR68W"}`,
			),
			rules: FilterAction{
				Exclude: []string{
					"repositories",
				},
				Include: []string{
					"repositories.organizations",
				},
			},
			filtered: true,
		},
	}

	for _, testCase := range testCases {
		testFilter(t, testCase.input, testCase.filtered, &FilterOptions{
			Loggers: testCase.rules,
		})
	}
}

func TestFilter_Messages(t *testing.T) {
	testCases := []struct {
		input    string
		rules    FilterAction
		filtered bool
	}{
		{
			input: fmt.Sprint(
				`{"logger":"repositories.users","level":"info",`,
				`"message":"User created","id":"01CV5FN4JF1STZMYDJWMGQR68W"}`,
			),
			rules: FilterAction{
				Exclude: []string{},
				Include: []string{
					"created",
				},
			},
			filtered: false,
		},
		{
			input: fmt.Sprint(
				`{"logger":"repositories.users","level":"info",`,
				`"message":"User created","id":"01CV5FN4JF1STZMYDJWMGQR68W"}`,
			),
			rules: FilterAction{
				Exclude: []string{
					"created",
				},
				Include: []string{},
			},
			filtered: true,
		},
		{
			input: fmt.Sprint(
				`{"logger":"repositories.users","level":"info",`,
				`"message":"User created","id":"01CV5FN4JF1STZMYDJWMGQR68W"}`,
			),
			rules: FilterAction{
				Exclude: []string{
					"updated",
				},
				Include: []string{},
			},
			filtered: false,
		},
		{
			input: fmt.Sprint(
				`{"logger":"repositories.users","level":"info",`,
				`"message":"User created","id":"01CV5FN4JF1STZMYDJWMGQR68W"}`,
			),
			rules: FilterAction{
				Exclude: []string{},
				Include: []string{
					"updated",
				},
			},
			filtered: true,
		},
		{
			input: fmt.Sprint(
				`{"logger":"repositories.users","level":"info",`,
				`"message":"User created","id":"01CV5FN4JF1STZMYDJWMGQR68W"}`,
			),
			rules: FilterAction{
				Exclude: []string{
					"User",
				},
				Include: []string{
					"User created",
				},
			},
			filtered: false,
		},
		{
			input: fmt.Sprint(
				`{"logger":"repositories.users","level":"info",`,
				`"message":"User created","id":"01CV5FN4JF1STZMYDJWMGQR68W"}`,
			),
			rules: FilterAction{
				Exclude: []string{
					"User",
				},
				Include: []string{
					"User updated",
				},
			},
			filtered: true,
		},
	}

	for _, testCase := range testCases {
		testFilter(t, testCase.input, testCase.filtered, &FilterOptions{
			Messages: testCase.rules,
		})
	}
}

func TestFilter_Levels(t *testing.T) {
	testCases := []struct {
		input    string
		rules    FilterAction
		filtered bool
	}{
		{
			input: fmt.Sprint(
				`{"logger":"repositories.users","level":"info",`,
				`"message":"User created","id":"01CV5FN4JF1STZMYDJWMGQR68W"}`,
			),
			rules: FilterAction{
				Exclude: []string{},
				Include: []string{
					"info",
				},
			},
			filtered: false,
		},
		{
			input: fmt.Sprint(
				`{"logger":"repositories.users","level":"info",`,
				`"message":"User created","id":"01CV5FN4JF1STZMYDJWMGQR68W"}`,
			),
			rules: FilterAction{
				Exclude: []string{
					"info",
				},
				Include: []string{},
			},
			filtered: true,
		},
		{
			input: fmt.Sprint(
				`{"logger":"repositories.users","level":"info",`,
				`"message":"User created","id":"01CV5FN4JF1STZMYDJWMGQR68W"}`,
			),
			rules: FilterAction{
				Exclude: []string{
					"debug",
					"warning",
					"error",
				},
				Include: []string{},
			},
			filtered: false,
		},
		{
			input: fmt.Sprint(
				`{"logger":"repositories.users","level":"info",`,
				`"message":"User created","id":"01CV5FN4JF1STZMYDJWMGQR68W"}`,
			),
			rules: FilterAction{
				Exclude: []string{},
				Include: []string{
					"debug",
					"warning",
					"error",
				},
			},
			filtered: true,
		},
		{
			input: fmt.Sprint(
				`{"logger":"repositories.users","level":"info",`,
				`"message":"User created","id":"01CV5FN4JF1STZMYDJWMGQR68W"}`,
			),
			rules: FilterAction{
				Exclude: []string{
					"info",
				},
				Include: []string{
					"info",
				},
			},
			filtered: false,
		},
	}

	for _, testCase := range testCases {
		testFilter(t, testCase.input, testCase.filtered, &FilterOptions{
			Levels: testCase.rules,
		})
	}
}

func TestFilter_Combined(t *testing.T) {
	testCases := []struct {
		input    string
		opts     *FilterOptions
		filtered bool
	}{
		{
			input: fmt.Sprint(
				`{"logger":"repositories.users","level":"info",`,
				`"message":"User created","id":"01CV5FN4JF1STZMYDJWMGQR68W"}`,
			),
			opts: &FilterOptions{
				Levels: FilterAction{
					Exclude: []string{},
					Include: []string{
						"info",
					},
				},
				Loggers: FilterAction{
					Exclude: []string{},
					Include: []string{
						"repositories.users",
					},
				},
			},
			filtered: false,
		},
		{
			input: fmt.Sprint(
				`{"logger":"repositories.users","level":"info",`,
				`"message":"User created","id":"01CV5FN4JF1STZMYDJWMGQR68W"}`,
			),
			opts: &FilterOptions{
				Levels: FilterAction{
					Exclude: []string{
						"info",
					},
					Include: []string{},
				},
				Loggers: FilterAction{
					Exclude: []string{},
					Include: []string{
						"repositories.users",
					},
				},
			},
			filtered: true,
		},
	}

	for _, testCase := range testCases {
		testFilter(t, testCase.input, testCase.filtered, testCase.opts)
	}
}

func testFilter(t *testing.T, input string, filtered bool, filter *FilterOptions) {
	output := Filter([]byte(input), filter)

	if filtered {
		if len(output) != 0 {
			t.Fatalf("Unexpected json output: '%s' using %+v", string(output), *filter)
		}
	} else {
		if len(output) == 0 {
			t.Fatalf("Unexpected json filter: '%s' using %+v", input, *filter)
		}
		if string(output) != input {
			t.Fatalf("Unexpected json output: '%s' should be '%s' using %+v",
				string(output), input, *filter)
		}
	}
}
