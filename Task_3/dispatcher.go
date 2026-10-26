package main

func Dispatch(notifiers []Notifier, message string) error {
	for _, notifier := range notifiers {
		if err := notifier.Send(message); err != nil {
			return err
		}
	}

	return nil
}
