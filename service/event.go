package service

type Event interface {
	Sync() SyncEvent
	Async() AsyncEvent
}

func New() Event {
	return &event{
		sync:  NewSync(),
		async: NewAsync(),
	}
}

type event struct {
	sync  SyncEvent
	async AsyncEvent
}

func (e *event) Sync() SyncEvent {
	return e.sync
}

func (e *event) Async() AsyncEvent {
	return e.async
}
