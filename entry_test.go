package soba_test

import (
	"strings"
	"testing"
	"time"

	random "github.com/Pallinder/go-randomdata"

	"github.com/novln/soba"
	"github.com/novln/soba/encoder/json"
)

// Benchmark allocation of new Entry.
func BenchmarkEntry_NewEntry(b *testing.B) {
	type entry struct {
		name    string
		message string
		fields  []soba.Field
	}

	list := []entry{}

	number := random.Number(100, 500)
	for i := 0; i < number; i++ {
		item := entry{}
		levels := []string{}
		for y := 0; y < random.Number(1, 10); y++ {
			if (y % 2) == 0 {
				levels = append(levels, strings.ToLower(random.Noun()))
			} else {
				levels = append(levels, strings.ToLower(random.Adjective()))
			}
		}
		item.name = strings.Join(levels, ".")
		item.message = random.Paragraph()
		for y := 0; y < random.Number(0, 20); y++ {
			if random.Number(0, 2) == 0 {
				item.fields = append(item.fields, soba.String(random.Noun(), random.Letters(20)))
			} else {
				item.fields = append(item.fields, soba.String(random.Adjective(), random.Letters(20)))
			}
		}
		list = append(list, item)
	}

	max := len(list)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		var e entry
		for pb.Next() {
			i = (i + 1) % max
			e = list[i]
			entry := soba.NewEntry(e.name, soba.DebugLevel, e.message, e.fields)
			entry.Flush()
		}
	})
}

// Benchmark allocation of writing Entry.
func BenchmarkEntry_WriteEntry(b *testing.B) {
	list := []*soba.Entry{}

	number := random.Number(100, 500)
	for i := 0; i < number; i++ {
		levels := []string{}
		for y := 0; y < random.Number(1, 10); y++ {
			if (y % 2) == 0 {
				levels = append(levels, strings.ToLower(random.Noun()))
			} else {
				levels = append(levels, strings.ToLower(random.Adjective()))
			}
		}
		name := strings.Join(levels, ".")
		message := random.Paragraph()
		fields := []soba.Field{}
		for y := 0; y < random.Number(0, 20); y++ {
			if random.Number(0, 2) == 0 {
				fields = append(fields, soba.String(random.Noun(), random.Letters(20)))
			} else {
				fields = append(fields, soba.String(random.Adjective(), random.Letters(20)))
			}
		}
		entry := soba.NewEntry(name, soba.InfoLevel, message, fields)
		list = append(list, entry)
	}

	max := len(list)
	encoder := NewTestEncoder()
	defer encoder.Close()

	b.ResetTimer()
	b.Run("no-encoder", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			i := 0
			var e *soba.Entry
			for pb.Next() {
				i = (i + 1) % max
				e = list[i]
				soba.WriteEntry(e, encoder)
			}
		})
	})
	b.Run("json-encoder", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			i := 0
			var e *soba.Entry
			for pb.Next() {
				encoder := json.NewEncoder()
				i = (i + 1) % max
				e = list[i]
				soba.WriteEntry(e, encoder)
				encoder.Close()
			}
		})
	})

}

// Test creation of new entry.
func TestEntry_New(t *testing.T) {
	name := "spectrum"
	level := soba.WarnLevel
	message := "Pellentesque non massa libero. Praesent a semper orci."
	before := time.Now()

	entry := soba.NewEntry(name, level, message, []soba.Field{
		soba.Bool("flower", true),
		soba.Int64("stream", 1545492862),
	})
	defer entry.Flush()

	after := time.Now()

	if entry.Name() != name {
		t.Fatalf(`Unexpected result for "name": "%v" should be "%v"`, entry.Name(), name)
	}
	if entry.Level() != level {
		t.Fatalf(`Unexpected result for "level": "%v" should be "%v"`, entry.Level(), level)
	}
	if entry.Message() != message {
		t.Fatalf(`Unexpected result for "message": "%v" should be "%v"`, entry.Message(), message)
	}
	if entry.Unix() < before.Unix() {
		t.Fatalf(`Unexpected result for "unix": "%v" should be greater than or equals to "%v"`,
			entry.Unix(), before.Unix())
	}
	if entry.Unix() > after.Unix() {
		t.Fatalf(`Unexpected result for "unix": "%v" should be less than or equals to "%v"`,
			entry.Unix(), after.Unix())
	}
	if len(entry.Fields()) != 2 {
		t.Fatalf(`Unexpected result for "fields": it should have "%d" fields, not "%d"`, 2, len(entry.Fields()))
	}
	f0 := DebugField(entry.Fields()[0])
	if f0 != `"flower":true` {
		t.Fatalf(`Unexpected result for "fields" at position 0: it should be "%s", not "%s"`,
			`"flower":true`, f0)
	}
	f1 := DebugField(entry.Fields()[1])
	if f1 != `"stream":1545492862` {
		t.Fatalf(`Unexpected result for "fields" at position 1: it should be "%s", not "%s"`,
			`"stream":1545492862`, f1)
	}
}

// Test entry in case of a duplicate name in fields.
func TestEntry_DuplicateFieldName(t *testing.T) {
	name := "prism"
	message := "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
	entry := soba.NewEntry(name, soba.InfoLevel, message,
		[]soba.Field{
			soba.Int("silver", 1),
			soba.Int("chrome", 2),
		},
		[]soba.Field{
			soba.Int("iron", 3),
			soba.Int("silver", 4),
			soba.Int("iron", 5),
			soba.Int("quartz", 6),
		},
	)
	defer entry.Flush()

	if len(entry.Fields()) != 4 {
		t.Fatalf(`Unexpected result for "fields": it should have "%d" fields, not "%d"`, 4, len(entry.Fields()))
	}
	f0 := DebugField(entry.Fields()[0])
	if f0 != `"silver":4` {
		t.Fatalf(`Unexpected result for "fields" at position 0: it should be "%s", not "%s"`,
			`"silver":4`, f0)
	}
	f1 := DebugField(entry.Fields()[1])
	if f1 != `"chrome":2` {
		t.Fatalf(`Unexpected result for "fields" at position 1: it should be "%s", not "%s"`,
			`"chrome":2`, f1)
	}
	f2 := DebugField(entry.Fields()[2])
	if f2 != `"iron":5` {
		t.Fatalf(`Unexpected result for "fields" at position 2: it should be "%s", not "%s"`,
			`"iron":5`, f2)
	}
	f3 := DebugField(entry.Fields()[3])
	if f3 != `"quartz":6` {
		t.Fatalf(`Unexpected result for "fields" at position 3: it should be "%s", not "%s"`,
			`"quartz":6`, f3)
	}
}

// Test entry in case of a duplicate name in fields.
func TestEntry_WriteEntry(t *testing.T) {
	encoder := json.NewEncoder()
	defer encoder.Close()

	entry := soba.NewEntry("test", soba.InfoLevel, "A log message", []soba.Field{
		soba.Bool("test", true),
	})
	defer entry.Flush()

	buffer := soba.WriteEntry(entry, encoder)
	line := strings.TrimSpace(string(buffer))

	// The variable line should have something like:
	// {"logger":"test","time":"2019-04-20T09:53:13Z","level":"info","message":"A log message","test":true}
	// In order to avoid guessing the time of generation, we'll only verify the
	// prefix and the suffix of this entry line...
	prefix := `{"logger":"test","time":`
	suffix := `","level":"info","message":"A log message","test":true}`

	if !strings.HasPrefix(line, prefix) {
		t.Fatalf("Unexpected entry line: %s", line)
	}

	if !strings.HasSuffix(line, suffix) {
		t.Fatalf("Unexpected entry line: %s", line)
	}
}
