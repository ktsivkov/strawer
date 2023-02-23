package async

import (
	"sync"
)

type Worker interface {
	Subscriber
	Dispatcher
	Getter
}

type Subscriber interface {
	Subscribe(handler Handler, targets []Type)
}

type Dispatcher interface {
	Dispatch(wg *sync.WaitGroup, event Event) (int, error)
}

type Getter interface {
	GetSubscribers(name Type) ([]Handler, error)
}

type Handler interface {
	Handle(e Event)
}

func New() Worker {
	return &worker{
		subs: map[Type][]Handler{},
	}
}
