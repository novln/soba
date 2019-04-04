package soba

import (
	"fmt"
	"strings"
	"time"
)

// A Field is an operation that add a key-value pair to the logger's context.
// Most fields are lazily marshaled, so it's inexpensive to add fields to disabled debug-level log statements.
type Field struct {
	name    string
	handler func(Encoder)
}

// Name returns field key.
func (field Field) Name() string {
	return field.name
}

// Write marshaled current field to given encoder so its key-value pair will be available in logger's context.
func (field Field) Write(encoder Encoder) {
	field.handler(encoder)
}

// NewField creates a new field.
func NewField(name string, handler func(Encoder)) Field {
	return Field{
		name:    strings.ToLower(name),
		handler: handler,
	}
}

// ----------------------------------------------------------------------------
// Object
// ----------------------------------------------------------------------------

// Object creates a typesafe Field with given key and ObjectMarshaler.
func Object(key string, value ObjectMarshaler) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddObject(key, value)
	})
}

// Int creates a typesafe Field with given key and int.
func Int(key string, value int) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddInt(key, value)
	})
}

// Int8 creates a typesafe Field with given key and int8.
func Int8(key string, value int8) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddInt8(key, value)
	})
}

// Int16 creates a typesafe Field with given key and int16.
func Int16(key string, value int16) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddInt16(key, value)
	})
}

// Int32 creates a typesafe Field with given key and int32.
func Int32(key string, value int32) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddInt32(key, value)
	})
}

// Int64 creates a typesafe Field with given key and int64.
func Int64(key string, value int64) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddInt64(key, value)
	})
}

// Uint creates a typesafe Field with given key and uint.
func Uint(key string, value uint) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddUint(key, value)
	})
}

// Uint8 creates a typesafe Field with given key and uint8.
func Uint8(key string, value uint8) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddUint8(key, value)
	})
}

// Uint16 creates a typesafe Field with given key and uint16.
func Uint16(key string, value uint16) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddUint16(key, value)
	})
}

// Uint32 creates a typesafe Field with given key and uint32.
func Uint32(key string, value uint32) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddUint32(key, value)
	})
}

// Uint64 creates a typesafe Field with given key and uint64.
func Uint64(key string, value uint64) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddUint64(key, value)
	})
}

// Float32 creates a typesafe Field with given key and float32.
func Float32(key string, value float32) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddFloat32(key, value)
	})
}

// Float64 creates a typesafe Field with given key and float64.
func Float64(key string, value float64) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddFloat64(key, value)
	})
}

// String creates a typesafe Field with given key and string.
func String(key, value string) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddString(key, value)
	})
}

// Stringer creates a typesafe Field with given key and Stringer.
func Stringer(key string, value fmt.Stringer) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddStringer(key, value)
	})
}

// Time creates a typesafe Field with given key and Time.
func Time(key string, value time.Time) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddTime(key, value)
	})
}

// Duration creates a typesafe Field with given key and Duration.
func Duration(key string, value time.Duration) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddDuration(key, value)
	})
}

// Bool creates a typesafe Field with given key and Bool.
func Bool(key string, value bool) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddBool(key, value)
	})
}

// Binary creates a typesafe Field with given key and slice of byte.
func Binary(key string, value []byte) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddBinary(key, value)
	})
}

// Skip is a no-op Field
func Skip(key string) Field {
	return NewField(key, func(encoder Encoder) {})
}

// Error is an alias of NamedError("error", err).
func Error(err error) Field {
	return NamedError("error", err)
}

// NamedError creates a typesafe Field with given key and error.
func NamedError(key string, err error) Field {

	if err == nil {
		return Skip(key)
	}

	return String(key, err.Error())
}

// ----------------------------------------------------------------------------
// Array
// ----------------------------------------------------------------------------

// Objects creates a typesafe Field with given key and collection of ObjectMarshaler.
func Objects(key string, values []ObjectMarshaler) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddObjects(key, values)
	})
}

// Array creates a typesafe Field with given key and ArrayMarshaler.
func Array(key string, value ArrayMarshaler) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddArray(key, value)
	})
}

// Ints creates a typesafe Field with given key and slice of int.
func Ints(key string, values []int) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddInts(key, values)
	})
}

// Int8s creates a typesafe Field with given key and slice of int8.
func Int8s(key string, values []int8) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddInt8s(key, values)
	})
}

// Int16s creates a typesafe Field with given key and slice of int16.
func Int16s(key string, values []int16) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddInt16s(key, values)
	})
}

// Int32s creates a typesafe Field with given key and slice of int32.
func Int32s(key string, values []int32) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddInt32s(key, values)
	})
}

// Int64s creates a typesafe Field with given key and slice of int64.
func Int64s(key string, values []int64) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddInt64s(key, values)
	})
}

// Uints creates a typesafe Field with given key and slice of uint.
func Uints(key string, values []uint) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddUints(key, values)
	})
}

// Uint8s creates a typesafe Field with given key and slice of uint8.
func Uint8s(key string, values []uint8) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddUint8s(key, values)
	})
}

// Uint16s creates a typesafe Field with given key and slice of uint16.
func Uint16s(key string, values []uint16) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddUint16s(key, values)
	})
}

// Uint32s creates a typesafe Field with given key and slice of uint32.
func Uint32s(key string, values []uint32) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddUint32s(key, values)
	})
}

// Uint64s creates a typesafe Field with given key and slice of uint64.
func Uint64s(key string, values []uint64) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddUint64s(key, values)
	})
}

// Float32s creates a typesafe Field with given key and slice of float32.
func Float32s(key string, values []float32) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddFloat32s(key, values)
	})
}

// Float64s creates a typesafe Field with given key and slice of float64.
func Float64s(key string, values []float64) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddFloat64s(key, values)
	})
}

// Strings creates a typesafe Field with given key and slice of string.
func Strings(key string, values []string) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddStrings(key, values)
	})
}

// Stringers creates a typesafe Field with given key and slice of Stringer.
func Stringers(key string, values []fmt.Stringer) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddStringers(key, values)
	})
}

// Times creates a typesafe Field with given key and slice of Time.
func Times(key string, values []time.Time) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddTimes(key, values)
	})
}

// Durations creates a typesafe Field with given key and slice of Duration.
func Durations(key string, values []time.Duration) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddDurations(key, values)
	})
}

// Bools creates a typesafe Field with given key and slice of boolean.
func Bools(key string, values []bool) Field {
	return NewField(key, func(encoder Encoder) {
		encoder.AddBools(key, values)
	})
}

// Errors creates a typesafe Field with given key and slice of error.
func Errors(key string, errors []error) Field {
	return NewField(key, func(encoder Encoder) {
		list := []string{}
		for i := range errors {
			list = append(list, errors[i].Error())
		}
		encoder.AddStrings(key, list)
	})
}
