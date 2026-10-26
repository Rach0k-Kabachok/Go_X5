package main

import (
	"errors"
	"fmt"
)

var ErrEmptyMessage = errors.New("message can't be empty")

type SendError struct {
	channel string
	err     error
}

func (err *SendError) Error() string {
	return fmt.Sprintf("%s: %v", err.channel, err.err)
}

func (err *SendError) Unwrap() error {
	return err.err
}

func (err *SendError) Is(target error) bool {
	other, ok := target.(*SendError)
	return ok && err != nil && other != nil && err.channel == other.channel
}
