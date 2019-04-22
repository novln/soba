package json

import (
	"sync"

	"github.com/novln/soba/encoder"
)

// Source forked from https://github.com/uber-go/zap and from https://github.com/rs/zerolog

// Encoder is a JSON encoder that isn't safe for concurrent access.
// To encode a new instance/object, you should use Encode() method that will handles a lot of boilerplate for you.
// Finally, when you have retrieve the buffer content, execute Close() method to recycles underlying resources.
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

// Encode start the initialization of a new instance/object.
// The given callback is used to provides object properties.
// At the end of it, it will returns the encoder content buffer, finished by a line break.
func (encoder *Encoder) Encode(handler func(encoder encoder.Encoder)) []byte {
	encoder.AppendBeginMarker()
	handler(encoder)
	encoder.AppendEndMarker()
	encoder.AppendLineBreak()
	return encoder.Bytes()
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

// Ensure Encoder implements encoder.Encoder interface at compile time.
var _ encoder.Encoder = &Encoder{}
