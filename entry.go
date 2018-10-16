package soba

import (
	"sync"
	"time"
)

// Entry represents a log event.
type Entry struct {
	name    string
	unix    int64
	level   Level
	message string
	fields  []Field
}

// Name returns entry name.
func (entry Entry) Name() string {
	return entry.name
}

// Level returns entry level.
func (entry Entry) Level() Level {
	return entry.level
}

// Message returns entry message.
func (entry Entry) Message() string {
	return entry.message
}

// Unix returns entry timestamp.
func (entry Entry) Unix() int64 {
	return entry.unix
}

// Fields returns entry fields.
func (entry Entry) Fields() []Field {
	return entry.fields
}

// Flush recycles entry.
func (entry *Entry) Flush() {
	if entry != nil {
		entryPool.Put(entry)
	}
}

// NewEntry creates a new entry with given configuration.
func NewEntry(name string, level Level, message string, fields ...[]Field) *Entry {
	e := entryPool.Get().(*Entry)
	e.name = name
	e.level = level
	e.message = message
	e.unix = time.Now().Unix()
	e.fields = e.fields[:0]
	for i := range fields {
		e.fields = append(e.fields, fields[i]...)
	}
	return e
}

// An entry pool to reduce memory allocation pressure.
var entryPool = &sync.Pool{
	New: func() interface{} {
		return &Entry{
			fields: make([]Field, 0, 64),
		}
	},
}
