package service

import (
	"github.com/ktsivkov/strawer"
	"sync"
	"testing"
)

type testAsyncHandler struct {
	called bool
}

func (h *testAsyncHandler) Handle(_ strawer.AsyncEvent) {
	h.called = true
}

func TestAsyncEvent_SubscribeAsyncSingleTarget(t *testing.T) {
	s := NewAsync()

	s.SubscribeAsync(&testAsyncHandler{
		called: false,
	}, []strawer.Type{"test"})
}

func TestAsyncEvent_SubscribeAsyncMultipleTargets(t *testing.T) {
	s := NewAsync()

	t1 := strawer.Type("test1")
	t2 := strawer.Type("test2")

	s.SubscribeAsync(&testAsyncHandler{
		called: false,
	}, []strawer.Type{t1, t2})
}

func TestAsyncEvent_GetAsyncSubscribersSingleTarget(t *testing.T) {
	s := NewAsync()

	s.SubscribeAsync(&testAsyncHandler{
		called: false,
	}, []strawer.Type{"test"})
	subs, err := s.GetAsyncSubscribers("test")

	if err != nil {
		t.Error("GetAsyncSubscribers resulted in an unexpected error.")
		t.Error(err)
	}

	if l := len(subs); l != 1 {
		t.Errorf("GetAsyncSubscribers expected exactly 1 result, got %d.", l)
	}
}

func TestAsyncEvent_GetAsyncSubscribersMultipleTargets(t *testing.T) {
	s := NewAsync()

	t1 := strawer.Type("test1")
	t2 := strawer.Type("test2")

	s.SubscribeAsync(&testAsyncHandler{
		called: false,
	}, []strawer.Type{t1, t2})
	ts1, err1 := s.GetAsyncSubscribers(t2)
	if err1 != nil {
		t.Error("GetAsyncSubscribers resulted in an unexpected error.")
		t.Error(err1)
	}
	if l1 := len(ts1); l1 != 1 {
		t.Errorf("GetAsyncSubscribers expected exactly 1 result, got %d.", l1)
	}

	ts2, err2 := s.GetAsyncSubscribers(t2)
	if err2 != nil {
		t.Error("GetAsyncSubscribers resulted in an unexpected error.")
		t.Error(err2)
	}
	if l2 := len(ts2); l2 != 1 {
		t.Errorf("GetAsyncSubscribers expected exactly 1 result, got %d.", l2)
	}
}

func TestAsyncEvent_DispatchAsyncWithSingleSubscriberAndNoWaitGroup(t *testing.T) {
	s := NewAsync()

	h := &testAsyncHandler{
		called: false,
	}
	t1 := strawer.Type("test")
	s.SubscribeAsync(h, []strawer.Type{t1})

	hc, err := s.DispatchAsync(nil, strawer.AsyncEvent{
		Type: t1,
		Data: nil,
	})
	if err != nil {
		t.Error("DispatchAsync resulted in an unexpected error.")
		t.Error(err)
	}
	if hc != 1 {
		t.Errorf("DispatchAsync was expected to call exactly 1 handler, got %d.", hc)
	}
}

func TestAsyncEvent_DispatchAsyncWithMultipleSubscribersAndNoWaitGroup(t *testing.T) {
	s := NewAsync()

	h1 := &testAsyncHandler{
		called: false,
	}
	h2 := &testAsyncHandler{
		called: false,
	}
	t1 := strawer.Type("test")
	s.SubscribeAsync(h1, []strawer.Type{t1})
	s.SubscribeAsync(h2, []strawer.Type{t1})

	hc, err := s.DispatchAsync(nil, strawer.AsyncEvent{
		Type: t1,
		Data: nil,
	})
	if err != nil {
		t.Error("DispatchAsync resulted in an unexpected error.")
		t.Error(err)
	}
	if hc != 2 {
		t.Errorf("DispatchAsync was expected to call exactly 2 handler, got %d.", hc)
	}
}

func TestAsyncEvent_DispatchAsyncWithSingleSubscriberAndWaitGroup(t *testing.T) {
	s := NewAsync()
	wg := &sync.WaitGroup{}

	h := &testAsyncHandler{
		called: false,
	}
	t1 := strawer.Type("test")
	s.SubscribeAsync(h, []strawer.Type{t1})

	hc, err := s.DispatchAsync(wg, strawer.AsyncEvent{
		Type: t1,
		Data: nil,
	})
	if err != nil {
		t.Error("DispatchAsync resulted in an unexpected error.")
		t.Error(err)
	}
	if hc != 1 {
		t.Errorf("DispatchAsync was expected to call exactly 1 handler, got %d.", hc)
	}

	wg.Wait()
	if h.called != true {
		t.Error("DispatchAsync method did not call the subscribed handler.")
	}
}

func TestAsyncEvent_DispatchAsyncWithMultipleSubscribersAndWaitGroup(t *testing.T) {
	s := NewAsync()
	wg := &sync.WaitGroup{}

	h1 := &testAsyncHandler{
		called: false,
	}
	h2 := &testAsyncHandler{
		called: false,
	}
	t1 := strawer.Type("test")
	s.SubscribeAsync(h1, []strawer.Type{t1})
	s.SubscribeAsync(h2, []strawer.Type{t1})

	hc, err := s.DispatchAsync(wg, strawer.AsyncEvent{
		Type: t1,
		Data: nil,
	})
	if err != nil {
		t.Error("DispatchAsync resulted in an unexpected error.")
		t.Error(err)
	}
	if hc != 2 {
		t.Errorf("DispatchAsync was expected to call exactly 2 handler, got %d.", hc)
	}

	wg.Wait()
	if h1.called != true {
		t.Error("DispatchAsync method did not call the first subscribed handler.")
	}
	if h2.called != true {
		t.Error("DispatchAsync method did not call the second subscribed handler.")
	}
}
