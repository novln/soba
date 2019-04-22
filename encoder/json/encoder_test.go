package json_test

import (
	"fmt"
	"testing"

	"github.com/novln/soba"
	libencoder "github.com/novln/soba/encoder"
	"github.com/novln/soba/encoder/json"
)

// Global encoder for benchmark, used to avoid compiler optimization.
var ge soba.Encoder

// TestArray is a simple struct to test ArrayMarshaler interface.
type TestArray struct {
	Enabled bool
	Status  string
	ID      int64
}

func (array TestArray) Encode(encoder libencoder.ArrayEncoder) {
	encoder.AppendBool(array.Enabled)
	encoder.AppendString(array.Status)
	encoder.AppendInt64(array.ID)
}

// TestObject is a simple struct to test ObjectMarshaler interface.
type TestObject struct {
	Enabled bool
	Status  string
	ID      int64
}

func (array TestObject) Encode(encoder libencoder.ObjectEncoder) {
	encoder.AddBool("enabled", array.Enabled)
	encoder.AddString("status", array.Status)
	encoder.AddInt64("id", array.ID)
}

// Benchmark allocation of new JSON Encoder.
func BenchmarkJSON_NewEncoder(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {

		// Create a local logger and update it's value from to prevent the
		// compiler to eliminate function execution.
		var e *json.Encoder

		for pb.Next() {
			e = json.NewEncoder()
			e.Close()
		}

		// Store logger instance in a global variable so the compiler cannot eliminate the benchmark.
		// It create a race conditions but it's okay since it's only a benchmark and not a test.
		ge = e

	})
}

func TestJSON_Encoder_Encode(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := fmt.Sprint(`{"foobar":42}`, "\n")

		buffer := encoder.Encode(func(e libencoder.Encoder) {
			e.AddInt("foobar", 42)
		})

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := fmt.Sprint(`{"shared":true}`, "\n")

		buffer := encoder.Encode(func(e libencoder.Encoder) {
			e.AddBool("shared", true)
		})

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}
