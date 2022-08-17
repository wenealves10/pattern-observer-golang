package notifier

import (
	"github.com/wenealves10/pattern-observer-golang/pkg/pattern"
)

type notifier struct {
	observers map[pattern.Observer]struct{}
}

func NewNotifier() *notifier {
	return &notifier{
		observers: make(map[pattern.Observer]struct{}),
	}
}

func (n *notifier) Register(o pattern.Observer) {
	n.observers[o] = struct{}{}
}

func (n *notifier) Unregister(o pattern.Observer) {
	delete(n.observers, o)
}

func (n *notifier) Notify(e pattern.Event) {
	for o := range n.observers {
		o.OnNotify(e)
	}
}
