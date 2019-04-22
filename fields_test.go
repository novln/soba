package soba_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/novln/soba"
	"github.com/novln/soba/encoder/json"
)

// DebugField returns a human readable field by using a JSON encoder.
func DebugField(field soba.Field) string {
	encoder := json.NewEncoder()
	defer encoder.Close()
	field.Write(encoder)
	src := encoder.Bytes()
	dst := make([]byte, len(src))
	copy(dst, src)
	return string(dst)
}

// TestObject is an ObjectMarshaler for test.
type TestObject struct {
	Key   string
	Value int64
}

func (o TestObject) Encode(encoder soba.ObjectEncoder) {
	encoder.AddString("key", o.Key)
	encoder.AddInt64("value", o.Value)
}

// Test creation of new field.
func TestField_New(t *testing.T) {
	encoder := json.NewEncoder()
	defer encoder.Close()

	name := "alpha"
	message := `"alpha":false`
	handler := func(encoder soba.Encoder) {
		encoder.AddBool("alpha", false)
	}

	field := soba.NewField(name, handler)

	if name != field.Name() {
		t.Fatalf("Unexpected field name: '%s' should be '%s'", field.Name(), name)
	}

	field.Write(encoder)
	buffer := encoder.Bytes()

	if message != string(buffer) {
		t.Fatalf("Unexpected field message: '%s' should be '%s'", message, string(buffer))
	}
}

// Test field with object.
func TestField_Object(t *testing.T) {
	object := &TestObject{
		Key:   "10.0.7.23",
		Value: 20,
	}
	field := soba.Object("key", object)
	expected := `"key":{"key":"10.0.7.23","value":20}`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with int.
func TestField_Int(t *testing.T) {
	field := soba.Int("key", 1)
	expected := `"key":1`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with int8.
func TestField_Int8(t *testing.T) {
	field := soba.Int8("key", 1)
	expected := `"key":1`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with int16.
func TestField_Int16(t *testing.T) {
	field := soba.Int16("key", 1)
	expected := `"key":1`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with int32.
func TestField_Int32(t *testing.T) {
	field := soba.Int32("key", 1)
	expected := `"key":1`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with int64.
func TestField_Int64(t *testing.T) {
	field := soba.Int64("key", 1)
	expected := `"key":1`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with uint.
func TestField_Uint(t *testing.T) {
	field := soba.Uint("key", 1)
	expected := `"key":1`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with uint8.
func TestField_Uint8(t *testing.T) {
	field := soba.Uint8("key", 1)
	expected := `"key":1`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with uint16.
func TestField_Uint16(t *testing.T) {
	field := soba.Uint16("key", 1)
	expected := `"key":1`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with uint32.
func TestField_Uint32(t *testing.T) {
	field := soba.Uint32("key", 1)
	expected := `"key":1`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with uint64.
func TestField_Uint64(t *testing.T) {
	field := soba.Uint64("key", 1)
	expected := `"key":1`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with float32.
func TestField_Float32(t *testing.T) {
	field := soba.Float32("key", 1.1)
	expected := `"key":1.1`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with float64.
func TestField_Float64(t *testing.T) {
	field := soba.Float64("key", 1.1)
	expected := `"key":1.1`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with string.
func TestField_String(t *testing.T) {
	field := soba.String("key", "1")
	expected := `"key":"1"`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with fmt.Stringer.
func TestField_Stringer(t *testing.T) {
	latency := 230 * time.Millisecond
	field := soba.Stringer("key", latency)
	expected := `"key":"230ms"`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with time.Time.
func TestField_Time(t *testing.T) {
	timezone, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		t.Fatalf("Unexpected error to load timezone: %s", err)
	}
	when := time.Unix(1540912844, 0).In(timezone)
	field := soba.Time("key", when)
	expected := `"key":"2018-10-30T16:20:44+01:00"`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with time.Duration.
func TestField_Duration(t *testing.T) {
	latency := (2 * time.Millisecond) + (523 * time.Microsecond)
	field := soba.Duration("key", latency)
	expected := `"key":"2.523ms"`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with bool.
func TestField_Bool(t *testing.T) {
	field := soba.Bool("key", true)
	expected := `"key":true`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with binary.
func TestField_Binary(t *testing.T) {
	field := soba.Binary("key", []byte("Hello world"))
	expected := `"key":"SGVsbG8gd29ybGQ="`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with error.
func TestField_Error(t *testing.T) {
	err := fmt.Errorf("foobar")
	field := soba.Error(err)
	expected := `"error":"foobar"`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with named error.
func TestField_NamedError(t *testing.T) {
	err := fmt.Errorf("foobar")
	field := soba.NamedError("validation", err)
	expected := `"validation":"foobar"`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with empty error.
func TestField_EmptyError(t *testing.T) {
	field := soba.Error(nil)
	expected := ``

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with null value.
func TestField_Null(t *testing.T) {
	field := soba.Null("key")
	expected := `"key":null`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// TODO: Add unit test for every field functions.
// - Arrays
