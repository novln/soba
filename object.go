package soba

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
)

// Object creates a typesafe Field with given key and ObjectMarshaler.
func Object(key string, value ObjectMarshaler) Field {
	return field(func(encoder Encoder) {
		encoder.AddObject(key, value)
	})
}

// Int creates a typesafe Field with given key and int.
func Int(key string, value int) Field {
	return field(func(encoder Encoder) {
		encoder.AddInt(key, value)
	})
}

// Int8 creates a typesafe Field with given key and int8.
func Int8(key string, value int8) Field {
	return field(func(encoder Encoder) {
		encoder.AddInt8(key, value)
	})
}

// Int16 creates a typesafe Field with given key and int16.
func Int16(key string, value int16) Field {
	return field(func(encoder Encoder) {
		encoder.AddInt16(key, value)
	})
}

// Int32 creates a typesafe Field with given key and int32.
func Int32(key string, value int32) Field {
	return field(func(encoder Encoder) {
		encoder.AddInt32(key, value)
	})
}

// Int64 creates a typesafe Field with given key and int64.
func Int64(key string, value int64) Field {
	return field(func(encoder Encoder) {
		encoder.AddInt64(key, value)
	})
}

// Uint creates a typesafe Field with given key and uint.
func Uint(key string, value uint) Field {
	return field(func(encoder Encoder) {
		encoder.AddUint(key, value)
	})
}

// Uint8 creates a typesafe Field with given key and uint8.
func Uint8(key string, value uint8) Field {
	return field(func(encoder Encoder) {
		encoder.AddUint8(key, value)
	})
}

// Uint16 creates a typesafe Field with given key and uint16.
func Uint16(key string, value uint16) Field {
	return field(func(encoder Encoder) {
		encoder.AddUint16(key, value)
	})
}

// Uint32 creates a typesafe Field with given key and uint32.
func Uint32(key string, value uint32) Field {
	return field(func(encoder Encoder) {
		encoder.AddUint32(key, value)
	})
}

// Uint64 creates a typesafe Field with given key and uint64.
func Uint64(key string, value uint64) Field {
	return field(func(encoder Encoder) {
		encoder.AddUint64(key, value)
	})
}

// Float32 creates a typesafe Field with given key and float32.
func Float32(key string, value float32) Field {
	return field(func(encoder Encoder) {
		encoder.AddFloat32(key, value)
	})
}

// Float64 creates a typesafe Field with given key and float64.
func Float64(key string, value float64) Field {
	return field(func(encoder Encoder) {
		encoder.AddFloat64(key, value)
	})
}

// String creates a typesafe Field with given key and string.
func String(key, value string) Field {
	return field(func(encoder Encoder) {
		encoder.AddString(key, value)
	})
}

// Stringer creates a typesafe Field with given key and Stringer.
func Stringer(key string, value fmt.Stringer) Field {
	return String(key, value.String())
}

// Time creates a typesafe Field with given key and Time.
func Time(key string, value time.Time) Field {
	return field(func(encoder Encoder) {
		encoder.AddTime(key, value)
	})
}

// Duration creates a typesafe Field with given key and Duration.
func Duration(key string, value time.Duration) Field {
	return field(func(encoder Encoder) {
		encoder.AddDuration(key, value)
	})
}

// Bool creates a typesafe Field with given key and Bool.
func Bool(key string, value bool) Field {
	return field(func(encoder Encoder) {
		encoder.AddBool(key, value)
	})
}

// Binary creates a typesafe Field with given key and slice of byte.
func Binary(key string, value []byte) Field {
	return field(func(encoder Encoder) {
		encoder.AddBinary(key, value)
	})
}

// Skip is a no-op Field
func Skip() Field {
	return field(func(encoder Encoder) {})
}

// Error is an alias of NamedError("error", err).
func Error(err error) Field {
	return NamedError("error", err)
}

// NamedError creates a typesafe Field with given key and error.
func NamedError(key string, err error) Field {
	if err == nil {
		return Skip()
	}

	value := err.Error()
	sterr, ok := err.(interface {
		StackTrace() errors.StackTrace
	})
	if ok {
		value += fmt.Sprintf(" %+v", sterr.StackTrace())
	}

	return String(key, value)
}
