package main

import (
	"errors"
	"testing"
)

func TestSend(t *testing.T) {
	tests := []struct {
		name     string
		notifier Notifier
		message  string
		wantErr  error
	}{
		{name: "email", notifier: EmailNotifier{}, message: "ready"},
		{name: "push", notifier: PushNotifier{}, message: "ready"},
		{name: "empty email message", notifier: EmailNotifier{}, wantErr: ErrEmptyMessage},
		{name: "empty push message", notifier: PushNotifier{}, wantErr: ErrEmptyMessage},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.notifier.Send(tt.message); !errors.Is(err, tt.wantErr) {
				t.Errorf("Send error = %v; want %v", err, tt.wantErr)
			}
		})
	}
}

type testNotifier struct {
	messages []string
	err      error
}

func (notifier *testNotifier) Send(message string) error {
	notifier.messages = append(notifier.messages, message)
	return notifier.err
}

func TestDispatch(t *testing.T) {
	first, second := &testNotifier{}, &testNotifier{}
	notifiers := []Notifier{first, second}
	if err := Dispatch(notifiers, "ready"); err != nil {
		t.Fatalf("Dispatch: %v", err)
	}
	for _, notifier := range []*testNotifier{first, second} {
		if len(notifier.messages) != 1 || notifier.messages[0] != "ready" {
			t.Errorf("messages = %v; want [ready]", notifier.messages)
		}
	}

	first.err = errors.New("send failed")
	second.messages = nil
	if err := Dispatch(notifiers, "ready"); !errors.Is(err, first.err) {
		t.Errorf("Dispatch error = %v; want %v", err, first.err)
	}
	if len(second.messages) != 0 {
		t.Error("Dispatch continued after error")
	}
}
