package main

import "fmt"

func SafeSend(notifier Notifier, message string) (err error) {
	defer func() {
		if value := recover(); value != nil {
			err = fmt.Errorf("send panic: %v", value)
		}
	}()

	return notifier.Send(message)
}

func Dispatch(notifiers []Notifier, message string) error {
	for index, notifier := range notifiers {
		if err := SafeSend(notifier, message); err != nil {
			return fmt.Errorf("notifier %d: %w", index+1, err)
		}
	}

	return nil
}
