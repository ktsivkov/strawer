package service

import (
	"errors"
	"github.com/ktsivkov/strawer"
	"sync"
)

type AsyncEvent interface {
	AsyncSubscriber
	AsyncDispatcher
	AsyncGetter
}

func NewAsync() AsyncEvent {
	return &asyncEvent{
		asyncSubs: map[strawer.Type][]strawer.AsyncHandler{},
	}
}

type asyncEvent struct {
	asyncSubs map[strawer.Type][]strawer.AsyncHandler
}

type AsyncSubscriber interface {
	SubscribeAsync(handler strawer.AsyncHandler, targets []strawer.Type)
}

type AsyncDispatcher interface {
	DispatchAsync(wg *sync.WaitGroup, event strawer.AsyncEvent) (int, error)
}

type AsyncGetter interface {
	GetAsyncSubscribers(name strawer.Type) ([]strawer.AsyncHandler, error)
}

func (e *asyncEvent) SubscribeAsync(handler strawer.AsyncHandler, targets []strawer.Type) {
	for _, eType := range targets {
		subscribers, _ := e.GetAsyncSubscribers(eType)
		e.asyncSubs[eType] = append(subscribers, handler)
	}
}

func (e *asyncEvent) DispatchAsync(wg *sync.WaitGroup, event strawer.AsyncEvent) (int, error) {
	subscribers, err := e.GetAsyncSubscribers(event.Type)
	if err != nil {
		return 0, err
	}
	for _, handler := range subscribers {
		if wg != nil {
			wg.Add(1)
		}
		h := handler
		go func() {
			h.Handle(event)
			if wg != nil {
				defer wg.Done()
			}
		}()
	}
	return len(subscribers), nil
}

func (e *asyncEvent) GetAsyncSubscribers(name strawer.Type) ([]strawer.AsyncHandler, error) {
	val, ok := e.asyncSubs[name]
	if ok {
		return val, nil
	}
	return []strawer.AsyncHandler{}, errors.New("no subscribers found")
}
