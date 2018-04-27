package soba

// A Field is an operation that add a key-value pair to the logger's context.
// Most fields are lazily marshaled, so it's inexpensive to add fields to disabled debug-level log statements.
type Field interface {
	// Write must register Field in given Encoder so its key-value pair will be available in logger's context.
	Write(Encoder)
}

// field is a lambda wrapper to implement Field interface.
type field func(Encoder)

func (e field) Write(encoder Encoder) {
	e(encoder)
}
