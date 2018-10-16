package json

import (
	"sync"
)

type Encoder struct {
	buffer []byte
}

func (encoder *Encoder) Bytes() []byte {
	return []byte{}
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
