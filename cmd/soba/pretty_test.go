package main

import (
	"fmt"
	"testing"
)

var EOL = "\n"

func TestPretty_Null(t *testing.T) {
	input := `null`
	expected := fmt.Sprint(`null`, EOL)

	testPretty(t, input, expected)
}

func TestPretty_Boolean(t *testing.T) {
	{
		input := `true`
		expected := fmt.Sprint(`true`, EOL)

		testPretty(t, input, expected)
	}
	{
		input := `false`
		expected := fmt.Sprint(`false`, EOL)

		testPretty(t, input, expected)
	}
}

func TestPretty_Number(t *testing.T) {
	{
		input := `10`
		expected := fmt.Sprint(`10`, EOL)

		testPretty(t, input, expected)
	}
	{
		input := `12`
		expected := fmt.Sprint(`12`, EOL)

		testPretty(t, input, expected)
	}
	{
		input := `-43`
		expected := fmt.Sprint(`-43`, EOL)

		testPretty(t, input, expected)
	}
	{
		input := `-200`
		expected := fmt.Sprint(`-200`, EOL)

		testPretty(t, input, expected)
	}
	{
		input := `1.34`
		expected := fmt.Sprint(`1.34`, EOL)

		testPretty(t, input, expected)
	}
	{
		input := `-6.141759`
		expected := fmt.Sprint(`-6.141759`, EOL)

		testPretty(t, input, expected)
	}
}

func TestPretty_String(t *testing.T) {
	{
		input := `"foobar"`
		expected := fmt.Sprint(`"foobar"`, EOL)

		testPretty(t, input, expected)
	}
	{
		input := `"bar"`
		expected := fmt.Sprint(`"bar"`, EOL)

		testPretty(t, input, expected)
	}
	{
		input := `"xyz"`
		expected := fmt.Sprint(`"xyz"`, EOL)

		testPretty(t, input, expected)
	}
}

func TestPretty_Object(t *testing.T) {
	{
		input := `{"key":"foobar"}`
		expected := fmt.Sprint(
			`{`, EOL,
			`  "key": "foobar"`, EOL,
			`}`, EOL,
		)

		testPretty(t, input, expected)
	}
	{
		input := `{"user":42,"first_name":"Allan","last_name":"GUICHARD","age":99}`
		expected := fmt.Sprint(
			`{`, EOL,
			`  "user": 42,`, EOL,
			`  "first_name": "Allan",`, EOL,
			`  "last_name": "GUICHARD",`, EOL,
			`  "age": 99`, EOL,
			`}`, EOL,
		)

		testPretty(t, input, expected)
	}
	{
		input := fmt.Sprint(
			`{"user":{"id":42,"identity":{"first_name":"Allan",`,
			`"last_name":"GUICHARD"},"deleted_at":null,"is_admin":true}}`,
		)
		expected := fmt.Sprint(
			`{`, EOL,
			`  "user": {`, EOL,
			`    "id": 42,`, EOL,
			`    "identity": {`, EOL,
			`      "first_name": "Allan",`, EOL,
			`      "last_name": "GUICHARD"`, EOL,
			`    },`, EOL,
			`    "deleted_at": null,`, EOL,
			`    "is_admin": true`, EOL,
			`  }`, EOL,
			`}`, EOL,
		)

		testPretty(t, input, expected)
	}
	{
		input := `{"keys":["k\\\"0","k\\\"1","k\\\"2"]}`
		expected := fmt.Sprint(
			`{`, EOL,
			`  "keys": [`, EOL,
			`    "k\\\"0",`, EOL,
			`    "k\\\"1",`, EOL,
			`    "k\\\"2"`, EOL,
			`  ]`, EOL,
			`}`, EOL,
		)

		testPretty(t, input, expected)
	}
}

func TestPretty_Array(t *testing.T) {
	{
		input := `[1,2,3,4,5,6,7]`
		expected := fmt.Sprint(
			`[`, EOL,
			`  1,`, EOL,
			`  2,`, EOL,
			`  3,`, EOL,
			`  4,`, EOL,
			`  5,`, EOL,
			`  6,`, EOL,
			`  7`, EOL,
			`]`, EOL,
		)

		testPretty(t, input, expected)
	}
	{
		input := `[{"key":42},{"key":22},{"key":68},{"key":6}]`
		expected := fmt.Sprint(
			`[`, EOL,
			`  {`, EOL,
			`    "key": 42`, EOL,
			`  },`, EOL,
			`  {`, EOL,
			`    "key": 22`, EOL,
			`  },`, EOL,
			`  {`, EOL,
			`    "key": 68`, EOL,
			`  },`, EOL,
			`  {`, EOL,
			`    "key": 6`, EOL,
			`  }`, EOL,
			`]`, EOL,
		)

		testPretty(t, input, expected)
	}
	{
		input := `[{"pack":[0,1,2]},{"pack":[3,4,5]},{"pack":[6,7,8]},{"pack":[9,10,11]}]`
		expected := fmt.Sprint(
			`[`, EOL,
			`  {`, EOL,
			`    "pack": [`, EOL,
			`      0,`, EOL,
			`      1,`, EOL,
			`      2`, EOL,
			`    ]`, EOL,
			`  },`, EOL,
			`  {`, EOL,
			`    "pack": [`, EOL,
			`      3,`, EOL,
			`      4,`, EOL,
			`      5`, EOL,
			`    ]`, EOL,
			`  },`, EOL,
			`  {`, EOL,
			`    "pack": [`, EOL,
			`      6,`, EOL,
			`      7,`, EOL,
			`      8`, EOL,
			`    ]`, EOL,
			`  },`, EOL,
			`  {`, EOL,
			`    "pack": [`, EOL,
			`      9,`, EOL,
			`      10,`, EOL,
			`      11`, EOL,
			`    ]`, EOL,
			`  }`, EOL,
			`]`, EOL,
		)

		testPretty(t, input, expected)
	}
}

func testPretty(t *testing.T, input string, expected string) {
	output := Pretty([]byte(input), &ColorOptions{
		EnableColor: false,
	})

	if string(output) != expected {
		t.Fatalf("Unexpected json output: '%s' should be '%s'", string(output), expected)
	}
}
