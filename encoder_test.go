package soba_test

import (
	"fmt"
	"time"

	"github.com/novln/soba"
	"github.com/novln/soba/encoder"
)

// TestEncoder is an encoder for unit test.
type TestEncoder struct {
}

func (encoder *TestEncoder) Open(handler func(encoder.Encoder)) {
	handler(encoder)
}

func (encoder *TestEncoder) Bytes() []byte {
	return nil
}

func (encoder *TestEncoder) Close() {}

func (encoder *TestEncoder) AddArray(key string, value encoder.ArrayMarshaler) {}

func (encoder *TestEncoder) AddObject(key string, value encoder.ObjectMarshaler) {}

func (encoder *TestEncoder) AddObjects(key string, values []encoder.ObjectMarshaler) {}

func (encoder *TestEncoder) AddInt(key string, value int) {}

func (encoder *TestEncoder) AddInts(key string, values []int) {}

func (encoder *TestEncoder) AddInt8(key string, value int8) {}

func (encoder *TestEncoder) AddInt8s(key string, values []int8) {}

func (encoder *TestEncoder) AddInt16(key string, value int16) {}

func (encoder *TestEncoder) AddInt16s(key string, values []int16) {}

func (encoder *TestEncoder) AddInt32(key string, value int32) {}

func (encoder *TestEncoder) AddInt32s(key string, values []int32) {}

func (encoder *TestEncoder) AddInt64(key string, value int64) {}

func (encoder *TestEncoder) AddInt64s(key string, values []int64) {}

func (encoder *TestEncoder) AddUint(key string, value uint) {}

func (encoder *TestEncoder) AddUints(key string, values []uint) {}

func (encoder *TestEncoder) AddUint8(key string, value uint8) {}

func (encoder *TestEncoder) AddUint8s(key string, values []uint8) {}

func (encoder *TestEncoder) AddUint16(key string, value uint16) {}

func (encoder *TestEncoder) AddUint16s(key string, values []uint16) {}

func (encoder *TestEncoder) AddUint32(key string, value uint32) {}

func (encoder *TestEncoder) AddUint32s(key string, values []uint32) {}

func (encoder *TestEncoder) AddUint64(key string, value uint64) {}

func (encoder *TestEncoder) AddUint64s(key string, values []uint64) {}

func (encoder *TestEncoder) AddFloat32(key string, value float32) {}

func (encoder *TestEncoder) AddFloat32s(key string, values []float32) {}

func (encoder *TestEncoder) AddFloat64(key string, value float64) {}

func (encoder *TestEncoder) AddFloat64s(key string, values []float64) {}

func (encoder *TestEncoder) AddString(key string, value string) {}

func (encoder *TestEncoder) AddStrings(key string, values []string) {}

func (encoder *TestEncoder) AddStringer(key string, value fmt.Stringer) {}

func (encoder *TestEncoder) AddStringers(key string, values []fmt.Stringer) {}

func (encoder *TestEncoder) AddTime(key string, value time.Time) {}

func (encoder *TestEncoder) AddTimes(key string, values []time.Time) {}

func (encoder *TestEncoder) AddDuration(key string, value time.Duration) {}

func (encoder *TestEncoder) AddDurations(key string, values []time.Duration) {}

func (encoder *TestEncoder) AddBool(key string, value bool) {}

func (encoder *TestEncoder) AddBools(key string, values []bool) {}

func (encoder *TestEncoder) AddBinary(key string, value []byte) {}

func (encoder *TestEncoder) AppendArray(value encoder.ArrayMarshaler) {}

func (encoder *TestEncoder) AppendObject(value encoder.ObjectMarshaler) {}

func (encoder *TestEncoder) AppendInt(value int) {}

func (encoder *TestEncoder) AppendInt8(value int8) {}

func (encoder *TestEncoder) AppendInt16(value int16) {}

func (encoder *TestEncoder) AppendInt32(value int32) {}

func (encoder *TestEncoder) AppendInt64(value int64) {}

func (encoder *TestEncoder) AppendUint(value uint) {}

func (encoder *TestEncoder) AppendUint8(value uint8) {}

func (encoder *TestEncoder) AppendUint16(value uint16) {}

func (encoder *TestEncoder) AppendUint32(value uint32) {}

func (encoder *TestEncoder) AppendUint64(value uint64) {}

func (encoder *TestEncoder) AppendFloat32(value float32) {}

func (encoder *TestEncoder) AppendFloat64(value float64) {}

func (encoder *TestEncoder) AppendString(value string) {}

func (encoder *TestEncoder) AppendBool(value bool) {}

func (encoder *TestEncoder) AppendTime(value time.Time) {}

func (encoder *TestEncoder) AppendDuration(value time.Duration) {}

func (encoder *TestEncoder) AppendBinary(value []byte) {}

// NewTestEncoder creates a new TestEncoder.
func NewTestEncoder() *TestEncoder {
	return &TestEncoder{}
}

// Ensure TestEncoder implements Encoder interface at compile time.
var _ soba.Encoder = &TestEncoder{}
