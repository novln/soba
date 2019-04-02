package json

import (
	"encoding/base64"
	"fmt"
	"math"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/novln/soba/encoder"
)

// Source forked from https://github.com/uber-go/zap and from https://github.com/rs/zerolog

// For JSON-escaping. See Encoder.safeAddString(string) below.
const hex = "0123456789abcdef"

// AddArray adds the field key with given ArrayMarshaler to the encoder buffer.
func (encoder *Encoder) AddArray(key string, value encoder.ArrayMarshaler) {
	encoder.AppendKey(key)
	encoder.AppendArray(value)
}

// AddObject adds the field key with given ObjectMarshaler to the encoder buffer.
func (encoder *Encoder) AddObject(key string, value encoder.ObjectMarshaler) {
	encoder.AppendKey(key)
	encoder.AppendObject(value)
}

// AddObjects adds the field key with given list of ObjectMarshaler to the encoder buffer.
func (encoder *Encoder) AddObjects(key string, values []encoder.ObjectMarshaler) {
	encoder.AppendKey(key)
	encoder.AppendArrayStart()
	for i := range values {
		encoder.AppendObject(values[i])
	}
	encoder.AppendArrayEnd()
}

// AddInt adds the field key with given integer to the encoder buffer.
func (encoder *Encoder) AddInt(key string, value int) {
	encoder.AppendKey(key)
	encoder.AppendInt(value)
}

// AddInts adds the field key with given list of integer to the encoder buffer.
func (encoder *Encoder) AddInts(key string, values []int) {
	encoder.AppendKey(key)
	encoder.AppendArrayStart()
	for i := range values {
		encoder.AppendInt(values[i])
	}
	encoder.AppendArrayEnd()
}

// AddInt8 adds the field key with given integer to the encoder buffer.
func (encoder *Encoder) AddInt8(key string, value int8) {
	encoder.AppendKey(key)
	encoder.AppendInt8(value)
}

// AddInt8s adds the field key with given list of integer to the encoder buffer.
func (encoder *Encoder) AddInt8s(key string, values []int8) {
	encoder.AppendKey(key)
	encoder.AppendArrayStart()
	for i := range values {
		encoder.AppendInt8(values[i])
	}
	encoder.AppendArrayEnd()
}

// AddInt16 adds the field key with given integer to the encoder buffer.
func (encoder *Encoder) AddInt16(key string, value int16) {
	encoder.AppendKey(key)
	encoder.AppendInt16(value)
}

// AddInt16s adds the field key with given list of integer to the encoder buffer.
func (encoder *Encoder) AddInt16s(key string, values []int16) {
	encoder.AppendKey(key)
	encoder.AppendArrayStart()
	for i := range values {
		encoder.AppendInt16(values[i])
	}
	encoder.AppendArrayEnd()
}

// AddInt32 adds the field key with given integer to the encoder buffer.
func (encoder *Encoder) AddInt32(key string, value int32) {
	encoder.AppendKey(key)
	encoder.AppendInt32(value)
}

// AddInt32s adds the field key with given list of integer to the encoder buffer.
func (encoder *Encoder) AddInt32s(key string, values []int32) {
	encoder.AppendKey(key)
	encoder.AppendArrayStart()
	for i := range values {
		encoder.AppendInt32(values[i])
	}
	encoder.AppendArrayEnd()
}

// AddInt64 adds the field key with given integer to the encoder buffer.
func (encoder *Encoder) AddInt64(key string, value int64) {
	encoder.AppendKey(key)
	encoder.AppendInt64(value)
}

// AddInt64s adds the field key with given list of integer to the encoder buffer.
func (encoder *Encoder) AddInt64s(key string, values []int64) {
	encoder.AppendKey(key)
	encoder.AppendArrayStart()
	for i := range values {
		encoder.AppendInt64(values[i])
	}
	encoder.AppendArrayEnd()
}

// AddUint adds the field key with given unsigned integer to the encoder buffer.
func (encoder *Encoder) AddUint(key string, value uint) {
	encoder.AppendKey(key)
	encoder.AppendUint(value)
}

// AddUints adds the field key with given list of unsigned integer to the encoder buffer.
func (encoder *Encoder) AddUints(key string, values []uint) {
	encoder.AppendKey(key)
	encoder.AppendArrayStart()
	for i := range values {
		encoder.AppendUint(values[i])
	}
	encoder.AppendArrayEnd()
}

// AddUint8 adds the field key with given unsigned integer to the encoder buffer.
func (encoder *Encoder) AddUint8(key string, value uint8) {
	encoder.AppendKey(key)
	encoder.AppendUint8(value)
}

// AddUint8s adds the field key with given list of unsigned integer to the encoder buffer.
func (encoder *Encoder) AddUint8s(key string, values []uint8) {
	encoder.AppendKey(key)
	encoder.AppendArrayStart()
	for i := range values {
		encoder.AppendUint8(values[i])
	}
	encoder.AppendArrayEnd()
}

// AddUint16 adds the field key with given unsigned integer to the encoder buffer.
func (encoder *Encoder) AddUint16(key string, value uint16) {
	encoder.AppendKey(key)
	encoder.AppendUint16(value)
}

// AddUint16s adds the field key with given list of unsigned integer to the encoder buffer.
func (encoder *Encoder) AddUint16s(key string, values []uint16) {
	encoder.AppendKey(key)
	encoder.AppendArrayStart()
	for i := range values {
		encoder.AppendUint16(values[i])
	}
	encoder.AppendArrayEnd()
}

// AddUint32 adds the field key with given unsigned integer to the encoder buffer.
func (encoder *Encoder) AddUint32(key string, value uint32) {
	encoder.AppendKey(key)
	encoder.AppendUint32(value)
}

// AddUint32s adds the field key with given list of unsigned integer to the encoder buffer.
func (encoder *Encoder) AddUint32s(key string, values []uint32) {
	encoder.AppendKey(key)
	encoder.AppendArrayStart()
	for i := range values {
		encoder.AppendUint32(values[i])
	}
	encoder.AppendArrayEnd()
}

// AddUint64 adds the field key with given unsigned integer to the encoder buffer.
func (encoder *Encoder) AddUint64(key string, value uint64) {
	encoder.AppendKey(key)
	encoder.AppendUint64(value)
}

// AddUint64s adds the field key with given list of unsigned integer to the encoder buffer.
func (encoder *Encoder) AddUint64s(key string, values []uint64) {
	encoder.AppendKey(key)
	encoder.AppendArrayStart()
	for i := range values {
		encoder.AppendUint64(values[i])
	}
	encoder.AppendArrayEnd()
}

// AddFloat32 adds the field key with given number to the encoder buffer.
func (encoder *Encoder) AddFloat32(key string, value float32) {
	encoder.AppendKey(key)
	encoder.AppendFloat32(value)
}

// AddFloat32s adds the field key with given list of number to the encoder buffer.
func (encoder *Encoder) AddFloat32s(key string, values []float32) {
	encoder.AppendKey(key)
	encoder.AppendArrayStart()
	for i := range values {
		encoder.AppendFloat32(values[i])
	}
	encoder.AppendArrayEnd()
}

// AddFloat64 adds the field key with given number to the encoder buffer.
func (encoder *Encoder) AddFloat64(key string, value float64) {
	encoder.AppendKey(key)
	encoder.AppendFloat64(value)
}

// AddFloat64s adds the field key with given list of number to the encoder buffer.
func (encoder *Encoder) AddFloat64s(key string, values []float64) {
	encoder.AppendKey(key)
	encoder.AppendArrayStart()
	for i := range values {
		encoder.AppendFloat64(values[i])
	}
	encoder.AppendArrayEnd()
}

// AddString adds the field key with given string to the encoder buffer.
func (encoder *Encoder) AddString(key string, value string) {
	encoder.AppendKey(key)
	encoder.AppendString(value)
}

// AddStrings adds the field key with given list of string to the encoder buffer.
func (encoder *Encoder) AddStrings(key string, values []string) {
	encoder.AppendKey(key)
	encoder.AppendArrayStart()
	for i := range values {
		encoder.AppendString(values[i])
	}
	encoder.AppendArrayEnd()
}

// AddStringer adds the field key with given Stringer to the encoder buffer.
func (encoder *Encoder) AddStringer(key string, value fmt.Stringer) {
	encoder.AppendKey(key)
	encoder.AppendString(value.String())
}

// AddStringers adds the field key with given list of Stringer to the encoder buffer.
func (encoder *Encoder) AddStringers(key string, values []fmt.Stringer) {
	encoder.AppendKey(key)
	encoder.AppendArrayStart()
	for i := range values {
		encoder.AppendString(values[i].String())
	}
	encoder.AppendArrayEnd()
}

// AddTime adds the field key with given Time to the encoder buffer.
func (encoder *Encoder) AddTime(key string, value time.Time) {
	encoder.AppendKey(key)
	encoder.AppendTime(value)
}

// AddTimes adds the field key with given list of Time to the encoder buffer.
func (encoder *Encoder) AddTimes(key string, values []time.Time) {
	encoder.AppendKey(key)
	encoder.AppendArrayStart()
	for i := range values {
		encoder.AppendTime(values[i])
	}
	encoder.AppendArrayEnd()
}

// AddDuration adds the field key with given Duration to the encoder buffer.
func (encoder *Encoder) AddDuration(key string, value time.Duration) {
	encoder.AppendKey(key)
	encoder.AppendDuration(value)
}

// AddDurations adds the field key with given list of Duration to the encoder buffer.
func (encoder *Encoder) AddDurations(key string, values []time.Duration) {
	encoder.AppendKey(key)
	encoder.AppendArrayStart()
	for i := range values {
		encoder.AppendDuration(values[i])
	}
	encoder.AppendArrayEnd()
}

// AddBool adds the field key with given boolean to the encoder buffer.
func (encoder *Encoder) AddBool(key string, value bool) {
	encoder.AppendKey(key)
	encoder.AppendBool(value)
}

// AddBools adds the field key with given list of boolean to the encoder buffer.
func (encoder *Encoder) AddBools(key string, values []bool) {
	encoder.AppendKey(key)
	encoder.AppendArrayStart()
	for i := range values {
		encoder.AppendBool(values[i])
	}
	encoder.AppendArrayEnd()
}

// AddBinary adds the field key with given buffer or bytes to the encoder buffer.
func (encoder *Encoder) AddBinary(key string, value []byte) {
	encoder.AppendKey(key)
	encoder.AppendBinary(value)
}

// AddNull adds the field key as a null value to the encoder buffer.
func (encoder *Encoder) AddNull(key string) {
	encoder.AppendKey(key)
	encoder.AppendNull()
}

// AppendArray converts the input array marshaler and appends the encoded value to the encoder buffer.
func (encoder *Encoder) AppendArray(value encoder.ArrayMarshaler) {
	encoder.AppendElementSeparator()
	encoder.AppendArrayStart()
	value.Encode(encoder)
	encoder.AppendArrayEnd()
}

// AppendObject converts the input object marshaler and appends the encoded value to the encoder buffer.
func (encoder *Encoder) AppendObject(value encoder.ObjectMarshaler) {
	encoder.AppendElementSeparator()
	encoder.AppendBeginMarker()
	value.Encode(encoder)
	encoder.AppendEndMarker()
}

// AppendInt converts the input integer and appends the encoded value to the encoder buffer.
func (encoder *Encoder) AppendInt(value int) {
	encoder.AppendInt64(int64(value))
}

// AppendInt8 converts the input integer and appends the encoded value to the encoder buffer.
func (encoder *Encoder) AppendInt8(value int8) {
	encoder.AppendInt64(int64(value))
}

// AppendInt16 converts the input integer and appends the encoded value to the encoder buffer.
func (encoder *Encoder) AppendInt16(value int16) {
	encoder.AppendInt64(int64(value))
}

// AppendInt32 converts the input integer and appends the encoded value to the encoder buffer.
func (encoder *Encoder) AppendInt32(value int32) {
	encoder.AppendInt64(int64(value))
}

// AppendInt64 converts the input integer and appends the encoded value to the encoder buffer.
func (encoder *Encoder) AppendInt64(value int64) {
	encoder.AppendElementSeparator()
	encoder.buffer = strconv.AppendInt(encoder.buffer, value, 10)
}

// AppendUint converts the input integer and appends the encoded value to the encoder buffer.
func (encoder *Encoder) AppendUint(value uint) {
	encoder.AppendUint64(uint64(value))
}

// AppendUint8 converts the input integer and appends the encoded value to the encoder buffer.
func (encoder *Encoder) AppendUint8(value uint8) {
	encoder.AppendUint64(uint64(value))
}

// AppendUint16 converts the input integer and appends the encoded value to the encoder buffer.
func (encoder *Encoder) AppendUint16(value uint16) {
	encoder.AppendUint64(uint64(value))
}

// AppendUint32 converts the input integer and appends the encoded value to the encoder buffer.
func (encoder *Encoder) AppendUint32(value uint32) {
	encoder.AppendUint64(uint64(value))
}

// AppendUint64 converts the input integer and appends the encoded value to the encoder buffer.
func (encoder *Encoder) AppendUint64(value uint64) {
	encoder.AppendElementSeparator()
	encoder.buffer = strconv.AppendUint(encoder.buffer, value, 10)
}

// AppendFloat32 converts the input number and appends the encoded value to the encoder buffer.
func (encoder *Encoder) AppendFloat32(value float32) {
	encoder.appendFloat(float64(value), 32)
}

// AppendFloat64 converts the input number and appends the encoded value to the encoder buffer.
func (encoder *Encoder) AppendFloat64(value float64) {
	encoder.appendFloat(value, 64)
}

// AppendString converts and escapes the input string and appends the encoded value to the encoder buffer.
func (encoder *Encoder) AppendString(value string) {
	encoder.AppendElementSeparator()
	encoder.buffer = append(encoder.buffer, '"')
	encoder.safeAddString(value)
	encoder.buffer = append(encoder.buffer, '"')
}

// AppendBool converts the input bool to a string and appends the encoded value to the encoder buffer.
func (encoder *Encoder) AppendBool(value bool) {
	encoder.AppendElementSeparator()
	encoder.buffer = strconv.AppendBool(encoder.buffer, value)
}

// AppendTime converts the input time to a string and appends the encoded value to the encoder buffer.
func (encoder *Encoder) AppendTime(value time.Time) {
	encoder.AppendElementSeparator()
	encoder.buffer = append(encoder.buffer, '"')
	encoder.safeAddByteString(value.AppendFormat(nil, time.RFC3339Nano))
	encoder.buffer = append(encoder.buffer, '"')
}

// AppendDuration converts the input duration to a string and appends the encoded value to the encoder buffer.
func (encoder *Encoder) AppendDuration(value time.Duration) {
	encoder.AppendElementSeparator()
	encoder.buffer = append(encoder.buffer, '"')
	encoder.safeAddString(value.String())
	encoder.buffer = append(encoder.buffer, '"')
}

// AppendBinary converts the input buffer or bytes to a string and appends the encoded value to the encoder buffer.
func (encoder *Encoder) AppendBinary(value []byte) {
	b64 := base64.StdEncoding
	encoder.AppendElementSeparator()
	encoder.buffer = append(encoder.buffer, '"')
	buffer := make([]byte, b64.EncodedLen(len(value)))
	b64.Encode(buffer, value)
	encoder.safeAddByteString(buffer)
	encoder.buffer = append(encoder.buffer, '"')
}

// AppendNull appends a null value to the encoder buffer.
func (encoder *Encoder) AppendNull() {
	encoder.AppendElementSeparator()
	encoder.buffer = append(encoder.buffer, 'n', 'u', 'l', 'l')
}

// appendFloat converts a number and appends it to the encoder buffer.
// Since JSON does not permit NaN or Infinity, we make a tradeoff and store those types as string.
func (encoder *Encoder) appendFloat(value float64, size int) {
	encoder.AppendElementSeparator()
	switch {
	case math.IsNaN(value):
		encoder.buffer = append(encoder.buffer, `"NaN"`...)
	case math.IsInf(value, 1):
		encoder.buffer = append(encoder.buffer, `"+Inf"`...)
	case math.IsInf(value, -1):
		encoder.buffer = append(encoder.buffer, `"-Inf"`...)
	default:
		encoder.buffer = strconv.AppendFloat(encoder.buffer, value, 'f', -1, size)
	}
}

// safeAddString JSON-escapes a string and appends it to the encoder buffer.
// Unlike the standard library's encoder, it doesn't attempt to protect the
// user from browser vulnerabilities or JSONP-related problems.
func (encoder *Encoder) safeAddString(value string) {
	i := 0
	for i < len(value) {
		if encoder.tryAddRuneSelf(value[i]) {
			i++
			continue
		}
		char, size := utf8.DecodeRuneInString(value[i:])
		if encoder.tryAddRuneError(char, size) {
			i++
			continue
		}
		encoder.buffer = append(encoder.buffer, value[i:i+size]...)
		i += size
	}
}

// safeAddByteString is no-alloc equivalent of safeAddString(string(s)) for a slice of byte.
func (encoder *Encoder) safeAddByteString(value []byte) {
	i := 0
	for i < len(value) {
		if encoder.tryAddRuneSelf(value[i]) {
			i++
			continue
		}
		char, size := utf8.DecodeRune(value[i:])
		if encoder.tryAddRuneError(char, size) {
			i++
			continue
		}
		encoder.buffer = append(encoder.buffer, value[i:i+size]...)
		i += size
	}
}

// tryAddRuneSelf appends given value if it is valid UTF-8 character represented in a single byte.
func (encoder *Encoder) tryAddRuneSelf(char byte) bool {
	if char >= utf8.RuneSelf {
		return false
	}
	if 0x20 <= char && char != '\\' && char != '"' {
		encoder.buffer = append(encoder.buffer, char)
		return true
	}
	switch char {
	case '\\', '"':
		encoder.buffer = append(encoder.buffer, '\\')
		encoder.buffer = append(encoder.buffer, char)
	case '\n':
		encoder.buffer = append(encoder.buffer, '\\')
		encoder.buffer = append(encoder.buffer, 'n')
	case '\r':
		encoder.buffer = append(encoder.buffer, '\\')
		encoder.buffer = append(encoder.buffer, 'r')
	case '\t':
		encoder.buffer = append(encoder.buffer, '\\')
		encoder.buffer = append(encoder.buffer, 't')
	default:
		// Encode bytes < 0x20, except for the escape sequences above.
		encoder.buffer = append(encoder.buffer, '\\', 'u', '0', '0')
		encoder.buffer = append(encoder.buffer, hex[char>>4], hex[char&0xF])
	}
	return true
}

func (encoder *Encoder) tryAddRuneError(char rune, size int) bool {
	if char == utf8.RuneError && size == 1 {
		encoder.buffer = append(encoder.buffer, '\\', 'u', 'f', 'f', 'f', 'd')
		return true
	}
	return false
}
