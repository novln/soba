package json_test

import (
	"fmt"
	"testing"

	"github.com/novln/soba"
	"github.com/novln/soba/json"
)

// Source forked from https://github.com/uber-go/zap and from https://github.com/rs/zerolog

// Global encoder for benchmark, used to avoid compiler optimization.
var ge soba.Encoder

// Benchmark allocation of new JSON Encoder.
func BenchmarkJSON_NewEncoder(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {

		// Create a local logger and update it's value from handler to prevent the
		// compiler to eliminate function execution.
		var e *json.Encoder

		for pb.Next() {
			e = json.NewEncoder()
			e.Close()
		}

		// Store logger instance in a global variable so the compiler cannot eliminate the benchmark.
		// It create a race conditions but it's okay since it's only a benchmark and not a unit test.
		//ge = e

	})
}

func TestJSON_Encoder_AddBool(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":false`

		encoder.AddBool("foo", false)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":false,"bar":true`
		encoder.AddBool("foo", false)
		encoder.AddBool("bar", true)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":false,"bar":true,"enabled":true`
		encoder.AddBool("foo", false)
		encoder.AddBool("bar", true)
		encoder.AddBool("enabled", true)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddString(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":"bar"`

		encoder.AddString("foo", "bar")
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"user":"01CV2N1Z7BE4AV5NBRMZBRT9CZ","group":"01CV2N2BJP3F6V5DA1QA9N5Q2D"`
		encoder.AddString("user", "01CV2N1Z7BE4AV5NBRMZBRT9CZ")
		encoder.AddString("group", "01CV2N2BJP3F6V5DA1QA9N5Q2D")
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"user":"01CV2N1Z7BE4AV5NBRMZBRT9CZ","group":"01CV2N2BJP3F6V5DA1QA9N5Q2D","description":"foobar"`
		encoder.AddString("user", "01CV2N1Z7BE4AV5NBRMZBRT9CZ")
		encoder.AddString("group", "01CV2N2BJP3F6V5DA1QA9N5Q2D")
		encoder.AddString("description", "foobar")
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := fmt.Sprint(
			`"name":"Les \"artisans\"","description":"\\\\ Communauté d'artisans //\n\n\t`,
			`Découvrez les ateliers d'initiation","website":"https://artisans.io"`,
		)
		encoder.AddString("name", `Les "artisans"`)
		encoder.AddString("description", `\\ Communauté d'artisans //`+"\n\n\t"+"Découvrez les ateliers d'initiation")
		encoder.AddString("website", "https://artisans.io")
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}
