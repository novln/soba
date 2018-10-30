package json

import (
	"strconv"
	"time"
	"unicode/utf8"
)

// Source forked from https://github.com/uber-go/zap and from https://github.com/rs/zerolog

// For JSON-escaping. See Encoder.safeAddString(string) below.
const hex = "0123456789abcdef"

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

// safeAddString JSON-escapes a string and appends it to the internal buffer.
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

// safeAddByteString is no-alloc equivalent of safeAddString(string(s)) for s []byte.
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

// tryAddRuneSelf appends b if it is valid UTF-8 character represented in a single byte.
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
