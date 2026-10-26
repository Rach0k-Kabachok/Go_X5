package main

import (
	"errors"
	"fmt"
)

var ErrEmptyMessage = errors.New("message can't be empty")

type Notifier interface {
	Send(message string) error
}

type BaseNotifier struct {
	channel string
}

func (notifier BaseNotifier) Log(message string) {
	fmt.Printf("[%s] %s\n", notifier.channel, message)
}

type EmailNotifier struct {
	BaseNotifier
}

func (notifier EmailNotifier) Send(message string) error {
	if message == "" {
		return ErrEmptyMessage
	}

	notifier.Log(message)
	return nil
}

type PushNotifier struct {
	BaseNotifier
}

func (notifier PushNotifier) Send(message string) error {
	if message == "" {
		return ErrEmptyMessage
	}

	notifier.Log(message)
	return nil
}
