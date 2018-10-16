package soba_test

import (
	"strings"
	"testing"

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
