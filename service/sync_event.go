package service

import (
	"errors"
	"github.com/ktsivkov/strawer"
)

type SyncEvent interface {
	SyncSubscriber
	SyncDispatcher
	SyncGetter
}

func NewSync() SyncEvent {
	return &syncEvent{
		subs: map[strawer.Type][]strawer.SyncHandler{},
	}
}

type syncEvent struct {
	subs map[strawer.Type][]strawer.SyncHandler
}

type SyncSubscriber interface {
	Subscribe(handler strawer.SyncHandler, targets []strawer.Type)
}

type SyncDispatcher interface {
	Dispatch(event strawer.Event) error
}

type SyncGetter interface {
	GetSubscribers(name strawer.Type) ([]strawer.SyncHandler, error)
}

func (e *syncEvent) Subscribe(handler strawer.SyncHandler, targets []strawer.Type) {
	for _, eType := range targets {
		subscribers, _ := e.GetSubscribers(eType)
		e.subs[eType] = append(subscribers, handler)
	}
}

func (e *syncEvent) Dispatch(event strawer.Event) error {
	subscribers, err := e.GetSubscribers(event.Type)
	if err != nil {
		return err
	}
	for _, h := range subscribers {
		h.Handle(event)
	}
	return nil
}

func (e *syncEvent) GetSubscribers(name strawer.Type) ([]strawer.SyncHandler, error) {
	val, ok := e.subs[name]
	if ok {
		return val, nil
	}
	return []strawer.SyncHandler{}, errors.New("no subscribers found")
}
