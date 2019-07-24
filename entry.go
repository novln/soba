package soba

import (
	"sync"
	"time"
)

const (
	// LoggerKey is the json key used for the entry name.
	LoggerKey = "logger"
	// TimeKey is the json key used for the entry timestamp.
	TimeKey = "time"
	// LevelKey is the json key used for the entry level.
	LevelKey = "level"
	// MessageKey is the json key used for the entry message.
	MessageKey = "message"
)

// Entry represents a log event.
type Entry struct {
	name    string
	unix    int64
	level   Level
	message string
	fields  []Field
	indexes map[string]int
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
		entry.fields = entry.fields[:0]
		for name := range entry.indexes {
			delete(entry.indexes, name)
		}
		entryPool.Put(entry)
	}
}

// NewEntry creates a new entry with given configuration.
// In case of field with duplicate name, the last one will be kept.
func NewEntry(name string, level Level, message string, fields ...[]Field) *Entry {
	entry := entryPool.Get().(*Entry)
	entry.name = name
	entry.level = level
	entry.message = message
	entry.unix = time.Now().Unix()

	for x := range fields {
		for y := range fields[x] {
			// Avoid duplication of field with the same name.
			// Last field overwrite the previous one.
			name := fields[x][y].name
			i, ok := entry.indexes[name]
			if !ok {
				entry.indexes[name] = len(entry.fields)
				entry.fields = append(entry.fields, fields[x][y])
			} else {
				entry.fields[i] = fields[x][y]
			}
		}
	}

	return entry
}

// WriteEntry writes entry informations on the given encoder.
func WriteEntry(entry *Entry, encoder Encoder) []byte {
	return encoder.Encode(func(encoder Encoder) {
		encoder.AddString(LoggerKey, entry.name)
		encoder.AddTime(TimeKey, time.Unix(entry.unix, 0).UTC())
		encoder.AddStringer(LevelKey, entry.level)
		encoder.AddString(MessageKey, entry.message)
		for _, field := range entry.fields {
			field.Write(encoder)
		}
	})
}

// An entry pool to reduce memory allocation pressure.
var entryPool = &sync.Pool{
	New: func() interface{} {
		return &Entry{
			indexes: make(map[string]int, 64),
			fields:  make([]Field, 0, 64),
		}
	},
}
