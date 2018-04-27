package soba

import (
	"time"
)

// EncoderMode defines the behavior on an Encoder.
type EncoderMode uint8

const (
	// UnknownMode indicates that the Encoder mode is undefined for the moment.
	UnknownMode = EncoderMode(iota)
	// ArrayMode indicates that the Encoder will encode array entities.
	ArrayMode
	// ObjectMode indicates that the Encoder will encode object entities.
	ObjectMode
)

func (m EncoderMode) String() string {
	switch m {
	case ArrayMode:
		return "ArrayMode"
	case ObjectMode:
		return "ObjectMode"
	default:
		return "UnknownMode"
	}
}

// Encoder is a strongly-typed, encoding-agnostic interface for adding array, map or struct-like object to the
// logging context. Also, be advised that Encoder aren't safe for concurrent use.
type Encoder interface {
	// Mode indicates the encoder's mode: object, array or neither.
	Mode() EncoderMode
	// Encoder is an ObjectEncoder in ObjectMode.
	ObjectEncoder
	// Encoder is an ArrayEncoder in ArrayMode.
	ArrayEncoder
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
	AppendByte(value byte)
}

// ObjectEncoder is a strongly-typed, encoding-agnostic interface for adding map or struct-like object to the
// logging context. Also, be advised that Encoder aren't safe for concurrent use.
type ObjectEncoder interface {
	AddArray(key string, value ArrayMarshaler)
	AddObject(key string, value ObjectMarshaler)
	AddObjects(key string, value []ObjectMarshaler)
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
	AddTime(key string, value time.Time)
	AddTimes(key string, values []time.Time)
	AddDuration(key string, value time.Duration)
	AddDurations(key string, values []time.Duration)
	AddBool(key string, value bool)
	AddBools(key string, values []bool)
	AddBinary(key string, value []byte)
	AddRaw(key string, value interface{})
}

// ObjectMarshaler define how an object can register itself in the logging context.
type ObjectMarshaler interface {
	EncodeObject(encoder ObjectEncoder)
}

// ArrayMarshaler define how an array can register itself in the logging context.
type ArrayMarshaler interface {
	EncodeArray(encoder ArrayEncoder)
}
