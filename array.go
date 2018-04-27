package soba

import (
	"time"
)

// Array creates a typesafe Field with given key and ArrayMarshaler.
func Array(key string, value ArrayMarshaler) Field {
	return field(func(encoder Encoder) {
		encoder.AddArray(key, value)
	})
}

// Ints creates a typesafe Field with given key and slice of int.
func Ints(key string, values []int) Field {
	return field(func(encoder Encoder) {
		encoder.AddInts(key, values)
	})
}

// Int8s creates a typesafe Field with given key and slice of int8.
func Int8s(key string, values []int8) Field {
	return field(func(encoder Encoder) {
		encoder.AddInt8s(key, values)
	})
}

// Int16s creates a typesafe Field with given key and slice of int16.
func Int16s(key string, values []int16) Field {
	return field(func(encoder Encoder) {
		encoder.AddInt16s(key, values)
	})
}

// Int32s creates a typesafe Field with given key and slice of int32.
func Int32s(key string, values []int32) Field {
	return field(func(encoder Encoder) {
		encoder.AddInt32s(key, values)
	})
}

// Int64s creates a typesafe Field with given key and slice of int64.
func Int64s(key string, values []int64) Field {
	return field(func(encoder Encoder) {
		encoder.AddInt64s(key, values)
	})
}

// Uints creates a typesafe Field with given key and slice of uint.
func Uints(key string, values []uint) Field {
	return field(func(encoder Encoder) {
		encoder.AddUints(key, values)
	})
}

// Uint8s creates a typesafe Field with given key and slice of uint8.
func Uint8s(key string, values []uint8) Field {
	return field(func(encoder Encoder) {
		encoder.AddUint8s(key, values)
	})
}

// Uint16s creates a typesafe Field with given key and slice of uint16.
func Uint16s(key string, values []uint16) Field {
	return field(func(encoder Encoder) {
		encoder.AddUint16s(key, values)
	})
}

// Uint32s creates a typesafe Field with given key and slice of uint32.
func Uint32s(key string, values []uint32) Field {
	return field(func(encoder Encoder) {
		encoder.AddUint32s(key, values)
	})
}

// Uint64s creates a typesafe Field with given key and slice of uint64.
func Uint64s(key string, values []uint64) Field {
	return field(func(encoder Encoder) {
		encoder.AddUint64s(key, values)
	})
}

// Float32s creates a typesafe Field with given key and slice of float32.
func Float32s(key string, values []float32) Field {
	return field(func(encoder Encoder) {
		encoder.AddFloat32s(key, values)
	})
}

// Float64s creates a typesafe Field with given key and slice of float64.
func Float64s(key string, values []float64) Field {
	return field(func(encoder Encoder) {
		encoder.AddFloat64s(key, values)
	})
}

// Strings creates a typesafe Field with given key and slice of string.
func Strings(key string, values []string) Field {
	return field(func(encoder Encoder) {
		encoder.AddStrings(key, values)
	})
}

// Times creates a typesafe Field with given key and slice of Time.
func Times(key string, values []time.Time) Field {
	return field(func(encoder Encoder) {
		encoder.AddTimes(key, values)
	})
}

// Durations creates a typesafe Field with given key and slice of Duration.
func Durations(key string, values []time.Duration) Field {
	return field(func(encoder Encoder) {
		encoder.AddDurations(key, values)
	})
}

// Bools creates a typesafe Field with given key and slice of boolean.
func Bools(key string, values []bool) Field {
	return field(func(encoder Encoder) {
		encoder.AddBools(key, values)
	})
}

// Errors creates a typesafe Field with given key and slice of error.
func Errors(key string, errors []error) Field {
	return field(func(encoder Encoder) {
		list := []string{}
		for i := range errors {
			list = append(list, errors[i].Error())
		}
		encoder.AddStrings(key, list)
	})
}
