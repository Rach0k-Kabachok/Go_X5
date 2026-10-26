package main

import (
	"errors"
	"fmt"
)

func main() {
	email := EmailNotifier{BaseNotifier: BaseNotifier{channel: "email"}}
	push := PushNotifier{BaseNotifier: BaseNotifier{channel: "push"}}
	notifiers := []Notifier{email, push}

	if err := Dispatch(notifiers, "Order #42 is ready"); err != nil {
		fmt.Println("Dispatch:", err)
	}

	err := Dispatch(notifiers, "")
	fmt.Println("Dispatch:", err)
	fmt.Println("Empty message:", errors.Is(err, ErrEmptyMessage))
	fmt.Println("Email error:", errors.Is(err, &SendError{channel: "email"}))

	var sendErr *SendError
	if errors.As(err, &sendErr) {
		fmt.Println("Failed channel:", sendErr.channel)
	}

	fmt.Println("Send panic:", SafeSend(BrokenNotifier{}, "ready"))
	fmt.Println("Log panic:", SafeSend(email, "panic"))
}
