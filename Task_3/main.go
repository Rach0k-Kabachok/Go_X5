package main

import "fmt"

func main() {
	notifiers := []Notifier{
		EmailNotifier{
			BaseNotifier: BaseNotifier{channel: "email"},
		},
		PushNotifier{
			BaseNotifier: BaseNotifier{channel: "push"},
		},
	}

	if err := Dispatch(notifiers, "Order #42 is ready"); err != nil {
		fmt.Println("Dispatch:", err)
		return
	}

	for _, notifier := range notifiers {
		if email, ok := notifier.(EmailNotifier); ok {
			fmt.Println("Email channel:", email.channel)
		}
	}
}
