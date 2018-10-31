package soba_test

import (
	"strings"
	"testing"
	"time"

	random "github.com/Pallinder/go-randomdata"

	"github.com/novln/soba"
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

// Test creation of new entry.
func TestEntry_New(t *testing.T) {
	name := "spectrum"
	level := soba.WarnLevel
	message := "Pellentesque non massa libero. Praesent a semper orci."
	before := time.Now()

	entry := soba.NewEntry(name, level, message, []soba.Field{
		soba.Bool("flower", true),
		soba.Time("stream", time.Now()),
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
	if entry.Fields()[0].Name() != "flower" {
		t.Fatalf(`Unexpected result for "fields" at position 0: it should have "%s" as name, not "%s"`,
			"flower", entry.Fields()[0].Name())
	}
	if entry.Fields()[1].Name() != "stream" {
		t.Fatalf(`Unexpected result for "fields" at position 1: it should have "%s" as name, not "%s"`,
			"stream", entry.Fields()[1].Name())
	}
	// TODO (novln): Check content of fields ?
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
	if entry.Fields()[0].Name() != "silver" {
		t.Fatalf(`Unexpected result for "fields" at position 0: it should have "%s" as name, not "%s"`,
			"silver", entry.Fields()[0].Name())
	}
	if entry.Fields()[1].Name() != "chrome" {
		t.Fatalf(`Unexpected result for "fields" at position 1: it should have "%s" as name, not "%s"`,
			"chrome", entry.Fields()[1].Name())
	}
	if entry.Fields()[2].Name() != "iron" {
		t.Fatalf(`Unexpected result for "fields" at position 1: it should have "%s" as name, not "%s"`,
			"iron", entry.Fields()[2].Name())
	}
	if entry.Fields()[3].Name() != "quartz" {
		t.Fatalf(`Unexpected result for "fields" at position 1: it should have "%s" as name, not "%s"`,
			"quartz", entry.Fields()[3].Name())
	}
	// TODO (novln): Check content of fields ?
}
