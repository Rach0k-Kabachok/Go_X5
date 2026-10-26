package main

import "fmt"

type Notifier interface {
	Send(message string) error
}

type BaseNotifier struct {
	channel string
}

func (notifier BaseNotifier) Log(message string) {
	if message == "panic" {
		panic("logger failed")
	}
	fmt.Printf("[%s] %s\n", notifier.channel, message)
}

type EmailNotifier struct {
	BaseNotifier
}

func (notifier EmailNotifier) Send(message string) error {
	if message == "" {
		return &SendError{channel: "email", err: ErrEmptyMessage}
	}

	notifier.Log(message)
	return nil
}

type PushNotifier struct {
	BaseNotifier
}

func (notifier PushNotifier) Send(message string) error {
	if message == "" {
		return &SendError{channel: "push", err: ErrEmptyMessage}
	}

	notifier.Log(message)
	return nil
}

type BrokenNotifier struct{}

func (BrokenNotifier) Send(message string) error {
	panic("channel driver failed")
}
