package main

import (
	"errors"
	"testing"
)

func TestSendError(t *testing.T) {
	err := Dispatch([]Notifier{EmailNotifier{}}, "")
	if !errors.Is(err, ErrEmptyMessage) {
		t.Fatalf("Dispatch error = %v; want %v", err, ErrEmptyMessage)
	}
	if !errors.Is(err, &SendError{channel: "email"}) {
		t.Error("Is did not recognize the email channel")
	}
	if errors.Is(err, &SendError{channel: "push"}) {
		t.Error("Is recognized the wrong channel")
	}

	var sendErr *SendError
	if !errors.As(err, &sendErr) || sendErr.channel != "email" {
		t.Fatalf("errors.As returned %v; want email SendError", sendErr)
	}
	if errors.Unwrap(sendErr) != ErrEmptyMessage {
		t.Error("Unwrap lost the original cause")
	}
}
