package main

import (
	"github.com/wenealves10/pattern-observer-golang/pkg/pattern"
	"github.com/wenealves10/pattern-observer-golang/pkg/pattern/notifier"
	"github.com/wenealves10/pattern-observer-golang/pkg/pattern/observer"
)

func main() {
	notifier := notifier.NewNotifier()

	notifier.Register(observer.NewObserver(1))
	notifier.Register(observer.NewObserver(2))

	notifier.Notify(pattern.Event{Data: 1})
	notifier.Notify(pattern.Event{Data: 2})
	notifier.Notify(pattern.Event{Data: 3})
	notifier.Notify(pattern.Event{Data: 4})
}
