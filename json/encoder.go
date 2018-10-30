package json

import (
	"sync"
)

// Source forked from https://github.com/uber-go/zap and from https://github.com/rs/zerolog

// Encoder is a JSON encoder that isn't safe for concurrent access.
type Encoder struct {
	buffer []byte
}

// Bytes return the encoder content buffer.
func (encoder *Encoder) Bytes() []byte {
	return encoder.buffer
}

// Close recycles underlying resources of encoder.
func (encoder *Encoder) Close() {
	// Proper usage of a sync.Pool requires each entry to have approximately
	// the same memory cost. To obtain this property when the stored type
	// contains a variably-sized buffer, we add a hard limit on the maximum buffer
	// to place back in the pool.
	//
	// See https://golang.org/issue/23199
	if encoder != nil && cap(encoder.buffer) < (1<<16) {
		encoderPool.Put(encoder)
	}
}

// AddString adds the field key with given string value to the encoder buffer.
func (encoder *Encoder) AddString(key string, value string) {
	encoder.AppendKey(key)
	encoder.AppendString(value)
}

// AddBool adds the field key with given boolean value to the encoder buffer.
func (encoder *Encoder) AddBool(key string, value bool) {
	encoder.AppendKey(key)
	encoder.AppendBool(value)
}

// NewEncoder creates a new JSON Encoder.
func NewEncoder() *Encoder {
	entry := encoderPool.Get().(*Encoder)
	entry.buffer = entry.buffer[:0]
	return entry
}

// An encoder pool to reduce memory allocation pressure.
var encoderPool = &sync.Pool{
	New: func() interface{} {
		return &Encoder{
			buffer: make([]byte, 0, 1024),
		}
	},
}
