package json

// Source forked from https://github.com/uber-go/zap and from https://github.com/rs/zerolog

// AppendBeginMarker inserts a map start into the internal buffer.
func (encoder *Encoder) AppendBeginMarker() {
	encoder.buffer = append(encoder.buffer, '{')
}

// AppendEndMarker inserts a map end into the internal buffer.
func (encoder *Encoder) AppendEndMarker() {
	encoder.buffer = append(encoder.buffer, '}')
}

// AppendLineBreak appends a line break.
func (encoder *Encoder) AppendLineBreak() {
	encoder.buffer = append(encoder.buffer, '\n')
}

// AppendArrayStart adds markers to indicate the start of an array.
func (encoder *Encoder) AppendArrayStart() {
	encoder.buffer = append(encoder.buffer, '[')
}

// AppendArrayEnd adds markers to indicate the end of an array.
func (encoder *Encoder) AppendArrayEnd() {
	encoder.buffer = append(encoder.buffer, ']')
}

// AppendKey appends a new key into the internal buffer.
func (encoder *Encoder) AppendKey(key string) {
	if len(encoder.buffer) > 1 && encoder.buffer[len(encoder.buffer)-1] != '{' {
		encoder.buffer = append(encoder.buffer, ',')
	}
	encoder.AppendString(key)
	encoder.buffer = append(encoder.buffer, ':')
}

// AppendElementSeparator appends a new separator into the internal buffer when required.
func (encoder *Encoder) AppendElementSeparator() {
	last := len(encoder.buffer) - 1
	if last < 0 {
		return
	}

	switch encoder.buffer[last] {
	case '{', '[', ':', ',', ' ':
		return
	default:
		encoder.buffer = append(encoder.buffer, ',')
	}
}
