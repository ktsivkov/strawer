package async

import (
	"errors"
	"testing"
)

func TestWorker_GetSubscribers(t *testing.T) {
	t.Run("Single target - Single subscription", func(t *testing.T) {
		s := New()

		s.Subscribe(&testHandler{
			called: false,
		}, []Type{"test"})

		subs, err := s.GetSubscribers("test")
		if err != nil {
			t.Error("GetSubscribers resulted in an unexpected error.")
			t.Error(err)
		}

		if l := len(subs); l != 1 {
			t.Errorf("GetSubscribers expected exactly 1 result, got %d.", l)
		}
	})

	t.Run("Multiple targets - Single subscription", func(t *testing.T) {
		s := New()

		t1 := Type("test1")
		t2 := Type("test2")

		s.Subscribe(&testHandler{
			called: false,
		}, []Type{t1, t2})

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
	})

	t.Run("Single Target - Multiple subscriptions", func(t *testing.T) {
		s := New()

		s.Subscribe(&testHandler{
			called: false,
		}, []Type{"test"})
		s.Subscribe(&testHandler{
			called: false,
		}, []Type{"test"})

		subs, err := s.GetSubscribers("test")
		if err != nil {
			t.Error("GetSubscribers resulted in an unexpected error.")
			t.Error(err)
		}

		if l := len(subs); l != 2 {
			t.Errorf("GetSubscribers expected exactly 1 result, got %d.", l)
		}
	})

	t.Run("Multiple targets - Multiple subscriptions", func(t *testing.T) {
		s := New()

		t1 := Type("test1")
		t2 := Type("test2")

		s.Subscribe(&testHandler{
			called: false,
		}, []Type{t1, t2})
		s.Subscribe(&testHandler{
			called: false,
		}, []Type{t1, t2})

		ts1, err1 := s.GetSubscribers(t2)
		if err1 != nil {
			t.Error("GetSubscribers resulted in an unexpected error.")
			t.Error(err1)
		}
		if l1 := len(ts1); l1 != 2 {
			t.Errorf("GetSubscribers expected exactly 1 result, got %d.", l1)
		}

		ts2, err2 := s.GetSubscribers(t2)
		if err2 != nil {
			t.Error("GetSubscribers resulted in an unexpected error.")
			t.Error(err2)
		}
		if l2 := len(ts2); l2 != 2 {
			t.Errorf("GetSubscribers expected exactly 1 result, got %d.", l2)
		}
	})

	t.Run("No subscribers", func(t *testing.T) {
		s := New()

		t1 := Type("test")

		ts, err := s.GetSubscribers(t1)
		if err == nil {
			t.Error("GetSubscribers was expected to return an error, got nil.")
			t.Error(err)
		}

		if errors.Is(err, &NoSubscribersError{}) == false {
			t.Error("GetSubscribers was expected to return an error of type NoSubscribersError.")
		}

		if err.Error() != "no subscribers found" {
			t.Errorf("GetSubscribers was expected to return an error with message \"no subscribers found\", got %s.", err.Error())
		}

		tsl := len(ts)
		if tsl != 0 {
			t.Errorf("GetSubscribers was expected to return no handlers, got %d.", tsl)
		}
	})
}
