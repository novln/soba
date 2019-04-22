package soba_test

import (
	"errors"
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

// TestArray is an ArrayEncoder for test.
type TestArray struct {
	Objects []TestObject
}

func (o TestArray) Encode(encoder soba.ArrayEncoder) {
	for _, object := range o.Objects {
		encoder.AppendObject(object)
	}
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

// Test field with a collection of objects.
func TestField_Objects(t *testing.T) {
	objects := []soba.ObjectMarshaler{
		&TestObject{
			Key:   "10.0.6.113",
			Value: 12,
		},
		&TestObject{
			Key:   "10.0.6.46",
			Value: 36,
		},
	}

	field := soba.Objects("key", objects)
	expected := `"key":[{"key":"10.0.6.113","value":12},{"key":"10.0.6.46","value":36}]`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with an array.
func TestField_Array(t *testing.T) {
	array := &TestArray{
		Objects: []TestObject{
			{
				Key:   "10.0.4.177",
				Value: 23,
			},
			{
				Key:   "10.0.4.21",
				Value: 255,
			},
		},
	}

	field := soba.Array("key", array)
	expected := `"key":[{"key":"10.0.4.177","value":23},{"key":"10.0.4.21","value":255}]`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with a list of int.
func TestField_Ints(t *testing.T) {
	list := []int{1, 2, 3, 4}

	field := soba.Ints("key", list)
	expected := `"key":[1,2,3,4]`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with a list of int8.
func TestField_Int8s(t *testing.T) {
	list := []int8{1, 2, 3, 4}

	field := soba.Int8s("key", list)
	expected := `"key":[1,2,3,4]`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with a list of int16.
func TestField_Int16s(t *testing.T) {
	list := []int16{1, 2, 3, 4}

	field := soba.Int16s("key", list)
	expected := `"key":[1,2,3,4]`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with a list of int32.
func TestField_Int32s(t *testing.T) {
	list := []int32{1, 2, 3, 4}

	field := soba.Int32s("key", list)
	expected := `"key":[1,2,3,4]`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with a list of int64.
func TestField_Int64s(t *testing.T) {
	list := []int64{1, 2, 3, 4}

	field := soba.Int64s("key", list)
	expected := `"key":[1,2,3,4]`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with a list of uint.
func TestField_Uints(t *testing.T) {
	list := []uint{1, 2, 3, 4}

	field := soba.Uints("key", list)
	expected := `"key":[1,2,3,4]`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with a list of uint8.
func TestField_Uint8s(t *testing.T) {
	list := []uint8{1, 2, 3, 4}

	field := soba.Uint8s("key", list)
	expected := `"key":[1,2,3,4]`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with a list of uint16.
func TestField_Uint16s(t *testing.T) {
	list := []uint16{1, 2, 3, 4}

	field := soba.Uint16s("key", list)
	expected := `"key":[1,2,3,4]`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with a list of uint32.
func TestField_Uint32s(t *testing.T) {
	list := []uint32{1, 2, 3, 4}

	field := soba.Uint32s("key", list)
	expected := `"key":[1,2,3,4]`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with a list of uint64.
func TestField_Uint64s(t *testing.T) {
	list := []uint64{1, 2, 3, 4}

	field := soba.Uint64s("key", list)
	expected := `"key":[1,2,3,4]`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with a list of float32.
func TestField_Float32s(t *testing.T) {
	list := []float32{1.1, 1.2, 1.3, 1.4}

	field := soba.Float32s("key", list)
	expected := `"key":[1.1,1.2,1.3,1.4]`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with a list of float64.
func TestField_Float64s(t *testing.T) {
	list := []float64{1.1, 1.2, 1.3, 1.4}

	field := soba.Float64s("key", list)
	expected := `"key":[1.1,1.2,1.3,1.4]`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with a list of string.
func TestField_Strings(t *testing.T) {
	list := []string{"abc", "ijk", "xyz"}

	field := soba.Strings("key", list)
	expected := `"key":["abc","ijk","xyz"]`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with a list of fmt.Stringer.
func TestField_Stringers(t *testing.T) {
	list := []fmt.Stringer{
		230 * time.Millisecond,
	}

	field := soba.Stringers("key", list)
	expected := `"key":["230ms"]`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with a list of time.Time.
func TestField_Times(t *testing.T) {
	timezone, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		t.Fatalf("Unexpected error to load timezone: %s", err)
	}
	list := []time.Time{
		time.Unix(1540912844, 0).In(timezone),
		time.Unix(1541113832, 0).In(timezone),
	}

	field := soba.Times("key", list)
	expected := `"key":["2018-10-30T16:20:44+01:00","2018-11-02T00:10:32+01:00"]`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with a list of time.Duration.
func TestField_Durations(t *testing.T) {
	list := []time.Duration{
		(3 * time.Millisecond) + (762 * time.Microsecond),
		(4 * time.Millisecond) + (275 * time.Microsecond),
	}

	field := soba.Durations("key", list)
	expected := `"key":["3.762ms","4.275ms"]`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with a list of bool.
func TestField_Bools(t *testing.T) {
	list := []bool{
		true, false, true,
	}

	field := soba.Bools("key", list)
	expected := `"key":[true,false,true]`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}

// Test field with a list of errors.
func TestField_Errors(t *testing.T) {
	list := []error{
		fmt.Errorf("unexpected error: %d", 80),
		errors.New("cannot open: /var/lib/mongodb/db"),
	}

	field := soba.Errors("key", list)
	expected := `"key":["unexpected error: 80","cannot open: /var/lib/mongodb/db"]`

	value := DebugField(field)

	if expected != value {
		t.Fatalf("Unexpected value: '%s' should be '%s'", value, expected)
	}
}
