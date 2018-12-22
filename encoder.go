package soba

import (
	"github.com/novln/soba/encoder"
)

// Aliasing from github.com/novln/soba/encoder package to avoid circular imports.

// Encoder is a strongly-typed, encoding-agnostic interface for adding array, map or struct-like object to the
// logging context.
// Also, be advised that Encoder aren't safe for concurrent use.
type Encoder = encoder.Encoder

// ArrayEncoder is a strongly-typed, encoding-agnostic interface for adding array to the logging context.
// Also, be advised that Encoder aren't safe for concurrent use.
type ArrayEncoder = encoder.ArrayEncoder

// ObjectEncoder is a strongly-typed, encoding-agnostic interface for adding map or struct-like object to the
// logging context.
// Also, be advised that Encoder aren't safe for concurrent use.
type ObjectEncoder = encoder.ObjectEncoder

// ObjectMarshaler define how an object can register itself in the logging context.
type ObjectMarshaler = encoder.ObjectMarshaler

// ArrayMarshaler define how an array can register itself in the logging context.
type ArrayMarshaler = encoder.ArrayMarshaler
