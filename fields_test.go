package soba_test

import (
	"github.com/novln/soba"
	"github.com/novln/soba/encoder/json"
)

// DebugField returns a human readable field by using a JSON encoder.
func DebugField(field soba.Field) string {
	encoder := json.NewEncoder()
	defer encoder.Close()
	field.Write(encoder)
	buffer := encoder.Bytes()
	return string(buffer)
}

// TODO: Add unit test for every field functions.
