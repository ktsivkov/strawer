package strawer

type Type string

type Event struct {
	Type Type
	Data any
}

type SyncHandler interface {
	Handle(e Event)
}

type AsyncHandler interface {
	Handle(e AsyncEvent)
}

type AsyncEvent struct {
	Type    Type
	Data    any
	Channel chan any
}
