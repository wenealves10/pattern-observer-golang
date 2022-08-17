package observer

import "github.com/wenealves10/pattern-observer-golang/pkg/pattern"

type observer struct {
	id int
}

func NewObserver(id int) *observer {
	return &observer{id: id}
}

func (o *observer) OnNotify(e pattern.Event) {
	println("Observer", o.id, "received", e.Data)
}
