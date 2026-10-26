package main

import (
	"errors"
	"testing"
)

type testNotifier struct {
	calls int
	err   error
}

func (notifier *testNotifier) Send(message string) error {
	notifier.calls++
	return notifier.err
}

func TestDispatch(t *testing.T) {
	tests := []struct {
		name      string
		firstErr  error
		wantCalls int
	}{
		{name: "all channels", wantCalls: 1},
		{name: "stop after error", firstErr: ErrEmptyMessage},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			first := &testNotifier{err: tt.firstErr}
			second := &testNotifier{}
			err := Dispatch([]Notifier{first, second}, "ready")
			if !errors.Is(err, tt.firstErr) {
				t.Errorf("Dispatch error = %v; want %v", err, tt.firstErr)
			}
			if first.calls != 1 || second.calls != tt.wantCalls {
				t.Errorf("calls = (%d, %d); want (1, %d)", first.calls, second.calls, tt.wantCalls)
			}
		})
	}
}

func TestSafeSend(t *testing.T) {
	tests := []struct {
		name     string
		notifier Notifier
		message  string
	}{
		{name: "panic in Send", notifier: BrokenNotifier{}, message: "ready"},
		{name: "panic in Log", notifier: EmailNotifier{}, message: "panic"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SafeSend(tt.notifier, tt.message); err == nil {
				t.Error("SafeSend returned nil after panic")
			}
		})
	}
}
