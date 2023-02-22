package service

import (
	"github.com/ktsivkov/strawer"
	"testing"
)

type testSyncHandler struct {
	called bool
}

func (h *testSyncHandler) Handle(_ strawer.Event) {
	h.called = true
}

func TestSyncEvent_SubscribeSingleTarget(t *testing.T) {
	s := NewSync()

	s.Subscribe(&testSyncHandler{
		called: false,
	}, []strawer.Type{"test"})
}

func TestSyncEvent_SubscribeMultipleTargets(t *testing.T) {
	s := NewSync()

	t1 := strawer.Type("test1")
	t2 := strawer.Type("test2")

	s.Subscribe(&testSyncHandler{
		called: false,
	}, []strawer.Type{t1, t2})
}

func TestSyncEvent_GetSubscribersSingleTarget(t *testing.T) {
	s := NewSync()

	s.Subscribe(&testSyncHandler{
		called: false,
	}, []strawer.Type{"test"})
	subs, err := s.GetSubscribers("test")

	if err != nil {
		t.Error("GetSubscribers resulted in an unexpected error.")
		t.Error(err)
	}

	if l := len(subs); l != 1 {
		t.Errorf("GetSubscribers expected exactly 1 result, got %d.", l)
	}
}

func TestSyncEvent_GetSubscribersMultipleTargets(t *testing.T) {
	s := NewSync()

	t1 := strawer.Type("test1")
	t2 := strawer.Type("test2")

	s.Subscribe(&testSyncHandler{
		called: false,
	}, []strawer.Type{t1, t2})
	ts1, err1 := s.GetSubscribers(t2)
	if err1 != nil {
		t.Error("GetSubscribers resulted in an unexpected error.")
		t.Error(err1)
	}
	if l1 := len(ts1); l1 != 1 {
		t.Errorf("GetSubscribers expected exactly 1 result, got %d.", l1)
	}

	ts2, err2 := s.GetSubscribers(t2)
	if err2 != nil {
		t.Error("GetSubscribers resulted in an unexpected error.")
		t.Error(err2)
	}
	if l2 := len(ts2); l2 != 1 {
		t.Errorf("GetSubscribers expected exactly 1 result, got %d.", l2)
	}
}

func TestSyncEvent_DispatchWithSingleSubscriber(t *testing.T) {
	s := NewSync()

	h := &testSyncHandler{
		called: false,
	}
	t1 := strawer.Type("test")
	s.Subscribe(h, []strawer.Type{t1})

	err := s.Dispatch(strawer.Event{
		Type: t1,
		Data: nil,
	})
	if err != nil {
		t.Error("Dispatch resulted in an unexpected error.")
		t.Error(err)
	}
	if h.called != true {
		t.Error("Dispatch method did not call the subscribed handler.")
	}
}

func TestSyncEvent_DispatchWithMultipleSubscribers(t *testing.T) {
	s := NewSync()

	h1 := &testSyncHandler{
		called: false,
	}
	h2 := &testSyncHandler{
		called: false,
	}
	t1 := strawer.Type("test")
	s.Subscribe(h1, []strawer.Type{t1})
	s.Subscribe(h2, []strawer.Type{t1})

	err := s.Dispatch(strawer.Event{
		Type: t1,
		Data: nil,
	})
	if err != nil {
		t.Error("Dispatch resulted in an unexpected error.")
		t.Error(err)
	}
	if h1.called != true {
		t.Error("Dispatch method did not call the first subscribed handler.")
	}
	if h2.called != true {
		t.Error("Dispatch method did not call the second subscribed handler.")
	}
}
