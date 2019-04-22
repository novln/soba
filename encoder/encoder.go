package encoder

import (
	"fmt"
	"time"
)

// Encoder is a strongly-typed, encoding-agnostic interface for adding array, map or struct-like object to the
// logging context.
// Also, be advised that Encoder aren't safe for concurrent use.
// Finally, when an encoder is finished, always execute Close() method to recycles underlying resources of encoder.
type Encoder interface {
	// An Encoder is a ObjectEncoder.
	ObjectEncoder
	// An Encoder is a ArrayEncoder.
	ArrayEncoder
	// Bytes returns the encoder content buffer.
	Bytes() []byte
	// Close recycles underlying resources of encoder.
	Close()
	// Encode start the initialization of a new instance/object.
	// The given callback is used to provides object properties.
	// At the end of it, it will returns the encoder content buffer, finished by a line break.
	Encode(handler func(encoder Encoder)) []byte
}

// ArrayEncoder is a strongly-typed, encoding-agnostic interface for adding array to the logging context.
// Also, be advised that Encoder aren't safe for concurrent use.
type ArrayEncoder interface {
	AppendArray(value ArrayMarshaler)
	AppendObject(value ObjectMarshaler)
	AppendInt(value int)
	AppendInt8(value int8)
	AppendInt16(value int16)
	AppendInt32(value int32)
	AppendInt64(value int64)
	AppendUint(value uint)
	AppendUint8(value uint8)
	AppendUint16(value uint16)
	AppendUint32(value uint32)
	AppendUint64(value uint64)
	AppendFloat32(value float32)
	AppendFloat64(value float64)
	AppendString(value string)
	AppendTime(value time.Time)
	AppendDuration(value time.Duration)
	AppendBool(value bool)
	AppendBinary(value []byte)
	AppendNull()
}

// ObjectEncoder is a strongly-typed, encoding-agnostic interface for adding map or struct-like object to the
// logging context.
// Also, be advised that Encoder aren't safe for concurrent use.
type ObjectEncoder interface {
	AddArray(key string, value ArrayMarshaler)
	AddObject(key string, value ObjectMarshaler)
	AddObjects(key string, values []ObjectMarshaler)
	AddInt(key string, value int)
	AddInts(key string, values []int)
	AddInt8(key string, value int8)
	AddInt8s(key string, values []int8)
	AddInt16(key string, value int16)
	AddInt16s(key string, values []int16)
	AddInt32(key string, value int32)
	AddInt32s(key string, values []int32)
	AddInt64(key string, value int64)
	AddInt64s(key string, values []int64)
	AddUint(key string, value uint)
	AddUints(key string, values []uint)
	AddUint8(key string, value uint8)
	AddUint8s(key string, values []uint8)
	AddUint16(key string, value uint16)
	AddUint16s(key string, values []uint16)
	AddUint32(key string, value uint32)
	AddUint32s(key string, values []uint32)
	AddUint64(key string, value uint64)
	AddUint64s(key string, values []uint64)
	AddFloat32(key string, value float32)
	AddFloat32s(key string, values []float32)
	AddFloat64(key string, value float64)
	AddFloat64s(key string, values []float64)
	AddString(key string, value string)
	AddStrings(key string, values []string)
	AddStringer(key string, value fmt.Stringer)
	AddStringers(key string, values []fmt.Stringer)
	AddTime(key string, value time.Time)
	AddTimes(key string, values []time.Time)
	AddDuration(key string, value time.Duration)
	AddDurations(key string, values []time.Duration)
	AddBool(key string, value bool)
	AddBools(key string, values []bool)
	AddBinary(key string, value []byte)
	AddNull(key string)
}

// ObjectMarshaler define how an object can register itself in the logging context.
type ObjectMarshaler interface {
	Encode(encoder ObjectEncoder)
}

// ArrayMarshaler define how an array can register itself in the logging context.
type ArrayMarshaler interface {
	Encode(encoder ArrayEncoder)
}
