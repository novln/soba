package soba_test

import (
	"testing"

	"github.com/novln/soba"
)

// Test parsing of entry level.
func TestLevel_Parse(t *testing.T) {
	scenarios := []struct {
		input    string
		valid    bool
		expected soba.Level
	}{
		{
			input:    "never",
			valid:    true,
			expected: soba.NoLevel,
		},
		{
			input:    "no",
			valid:    true,
			expected: soba.NoLevel,
		},
		{
			input:    "none",
			valid:    true,
			expected: soba.NoLevel,
		},
		{
			input:    "disable",
			valid:    true,
			expected: soba.NoLevel,
		},
		{
			input:    "disabled",
			valid:    true,
			expected: soba.NoLevel,
		},
		{
			input:    "verbose",
			valid:    true,
			expected: soba.DebugLevel,
		},
		{
			input:    "debug",
			valid:    true,
			expected: soba.DebugLevel,
		},
		{
			input:    "info",
			valid:    true,
			expected: soba.InfoLevel,
		},
		{
			input:    "warning",
			valid:    true,
			expected: soba.WarnLevel,
		},
		{
			input:    "warn",
			valid:    true,
			expected: soba.WarnLevel,
		},
		{
			input:    "error",
			valid:    true,
			expected: soba.ErrorLevel,
		},
		{
			input:    "unknown",
			valid:    false,
			expected: soba.UnknownLevel,
		},
		{
			input:    "foobar",
			valid:    false,
			expected: soba.UnknownLevel,
		},
		{
			input:    "Zzzz",
			valid:    false,
			expected: soba.UnknownLevel,
		},
	}

	for _, scenario := range scenarios {
		level, ok := soba.ParseLevel(scenario.input)

		if scenario.valid {
			if !ok {
				t.Fatalf(`Unexpected result for "%s": must be a valid level`, scenario.input)
			}
		} else {
			if ok {
				t.Fatalf(`Unexpected result for "%s": should be an invalid level`, scenario.input)
			}
		}

		if level != scenario.expected {
			t.Fatalf("Unexpected result for %s: %v should be %v", scenario.input, level, scenario.expected)
		}
	}
}

// Test printing of entry level.
func TestLevel_String(t *testing.T) {
	scenarios := []struct {
		input    soba.Level
		name     string
		expected string
	}{
		{
			input:    soba.NoLevel,
			name:     "soba.NoLevel",
			expected: "never",
		},
		{
			input:    soba.DebugLevel,
			name:     "soba.DebugLevel",
			expected: "debug",
		},
		{
			input:    soba.InfoLevel,
			name:     "soba.InfoLevel",
			expected: "info",
		},
		{
			input:    soba.WarnLevel,
			name:     "soba.WarnLevel",
			expected: "warning",
		},
		{
			input:    soba.ErrorLevel,
			name:     "soba.ErrorLevel",
			expected: "error",
		},
		{
			input:    soba.UnknownLevel,
			name:     "soba.UnknownLevel",
			expected: "unknown",
		},
	}

	for _, scenario := range scenarios {
		output := scenario.input.String()
		if output != scenario.expected {
			t.Fatalf("Unexpected result for %s: %v should be %v", scenario.name, output, scenario.expected)
		}
	}
}
