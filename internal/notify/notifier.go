package notify

import (
	gosxnotifier "github.com/deckarep/gosx-notifier"
)

type Notifier struct{}

func New() Notifier {
	return Notifier{}
}

type Note struct {
	Title string
	Body  string
	Link  string
}

func (n Notifier) Notify(note Note) error {
	notif := gosxnotifier.NewNotification(note.Body)
	notif.Title = note.Title
	notif.Link = note.Link

	return notif.Push()
}
