package soba_test

import (
	"testing"

	"github.com/novln/soba"
)

// Test registration of appenders.
func TestPlugins_RegisterAppenders(t *testing.T) {
	apiAppender := NewTestAppender("api-log")
	defer CloseAppender(t, apiAppender)

	dbAppender := NewTestAppender("db-log")
	defer CloseAppender(t, dbAppender)

	authAppender := NewTestAppender("auth-log")
	defer CloseAppender(t, authAppender)

	stdoutAppender := NewTestAppender("stdout")
	defer CloseAppender(t, stdoutAppender)

	err := soba.RegisterAppenders(apiAppender, dbAppender, authAppender, stdoutAppender)
	if err != nil {
		t.Fatalf("Unexpected error: %+v", err)
	}

	invalidAppender := NewTestAppender("Notifier")
	defer CloseAppender(t, invalidAppender)

	err = soba.RegisterAppenders(invalidAppender)
	if err == nil {
		t.Fatal("An error was expected")
	}
}
