package soba_test

import (
	"fmt"
	"testing"

	"github.com/novln/soba"
)

// Test logger name format.
func TestLogger_IsNameValid(t *testing.T) {

	scenarios := []struct {
		name  string
		valid bool
	}{
		{
			// Scenario #1
			name:  "hello",
			valid: true,
		},
		{
			// Scenario #2
			name:  "hello.world",
			valid: true,
		},
		{
			// Scenario #3
			name:  "Hello.World",
			valid: false,
		},
		{
			// Scenario #4
			name:  "Hello/World",
			valid: false,
		},
		{
			// Scenario #5
			name:  "hello/world",
			valid: false,
		},
		{
			// Scenario #6
			name:  "hello.world0",
			valid: true,
		},
		{
			// Scenario #7
			name:  "hello_world0",
			valid: true,
		},
		{
			// Scenario #8
			name:  "0hello_world0",
			valid: false,
		},
		{
			// Scenario #9
			name:  "_hello_world_",
			valid: false,
		},
		{
			// Scenario #10
			name:  "hello_world_",
			valid: false,
		},
		{
			// Scenario #11
			name:  "hello_world",
			valid: true,
		},
		{
			// Scenario #12
			name:  "xx",
			valid: false,
		},
		{
			// Scenario #13
			name:  "",
			valid: false,
		},
		{
			// Scenario #14
			name:  "foo",
			valid: true,
		},
		{
			// Scenario #15
			name:  "hello-world",
			valid: true,
		},
		{
			// Scenario #16
			name:  "hello-world0",
			valid: true,
		},
	}

	for i, scenario := range scenarios {
		message := fmt.Sprintf("scenario #%d", (i + 1))
		result := soba.IsLoggerNameValid(scenario.name)
		if result != scenario.valid {
			t.Fatalf("Unexpected result for %s: %v should be %v", message, result, scenario.valid)
		}
	}

}

// TODO More test on logger
