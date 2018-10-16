package soba_test

import (
	"testing"

	"github.com/novln/soba"
)

// Test registration of appenders.
func TestPlugins_RegisterAppenders(t *testing.T) {
	apiAppender := NewTestAppender("api-log")
	dbAppender := NewTestAppender("db-log")
	authAppender := NewTestAppender("auth-log")
	stdoutAppender := NewTestAppender("stdout")

	err := soba.RegisterAppenders(apiAppender, dbAppender, authAppender, stdoutAppender)
	if err != nil {
		t.Fatal(err)
	}

	invalidAppender := NewTestAppender("Notifier")
	err = soba.RegisterAppenders(invalidAppender)
	if err == nil {
		t.Fatal("An error was expected")
	}
}
