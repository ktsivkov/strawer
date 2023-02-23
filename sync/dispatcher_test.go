package sync

import (
	"errors"
	"testing"
)

func TestWorker_Dispatch(t *testing.T) {
	t.Run("Single subscriber", func(t *testing.T) {
		t.Parallel()
		s := New()

		h := &testHandler{
			called: false,
		}
		t1 := Type("test")
		s.Subscribe(h, []Type{t1})

		hc, err := s.Dispatch(Event{
			Type: t1,
			Data: nil,
		})
		if err != nil {
			t.Error("Dispatch resulted in an unexpected error.")
			t.Error(err)
		}
		if hc != 1 {
			t.Errorf("Dispatch was expected to call exactly 1 handler, got %d.", hc)
		}

		if h.called != true {
			t.Error("Dispatch method did not call the subscribed handler.")
		}
	})

	t.Run("Multiple subscribers", func(t *testing.T) {
		t.Parallel()
		s := New()

		h1 := &testHandler{
			called: false,
		}
		h2 := &testHandler{
			called: false,
		}
		t1 := Type("test")
		s.Subscribe(h1, []Type{t1})
		s.Subscribe(h2, []Type{t1})

		hc, err := s.Dispatch(Event{
			Type: t1,
			Data: nil,
		})
		if err != nil {
			t.Error("Dispatch resulted in an unexpected error.")
			t.Error(err)
		}
		if hc != 2 {
			t.Errorf("Dispatch was expected to call exactly 2 handlers, got %d.", hc)
		}

		if h1.called != true {
			t.Error("Dispatch method did not call the first subscribed handler.")
		}
		if h2.called != true {
			t.Error("Dispatch method did not call the second subscribed handler.")
		}
	})

	t.Run("No subscribers", func(t *testing.T) {
		t.Parallel()
		s := New()

		t1 := Type("test")

		hc, err := s.Dispatch(Event{
			Type: t1,
			Data: nil,
		})

		if err == nil {
			t.Error("Dispatch was expected to return an error, got nil.")
			t.Error(err)
		}
		if errors.Is(err, &NoSubscribersError{}) == false {
			t.Error("Dispatch was expected to return an error of type NoSubscribersError.")
		}
		if err.Error() != "no subscribers found" {
			t.Errorf("Dispatch was expected to return an error with message \"no subscribers found\", got %s.", err.Error())
		}

		if hc != 0 {
			t.Errorf("Dispatch was expected to call exactly 0 handlers, got %d.", hc)
		}
	})
}

func BenchmarkWorker_Dispatch(b *testing.B) {
	b.Run("Dispatch with a handler", func(b *testing.B) {
		s := New()
		t1 := Type("test")
		h1 := &testHandler{
			called: false,
		}

		s.Subscribe(h1, []Type{t1})

		for n := 0; n < b.N; n++ {
			_, _ = s.Dispatch(Event{
				Type: t1,
				Data: nil,
			})
		}
	})

	b.Run("Dispatch with two handlers", func(b *testing.B) {
		s := New()
		t1 := Type("test")

		s.Subscribe(&testHandler{
			called: false,
		}, []Type{t1})
		s.Subscribe(&testHandler{
			called: false,
		}, []Type{t1})

		for n := 0; n < b.N; n++ {
			_, _ = s.Dispatch(Event{
				Type: t1,
				Data: nil,
			})
		}
	})

	b.Run("Dispatch with no handlers", func(b *testing.B) {
		s := New()
		t1 := Type("test")

		for n := 0; n < b.N; n++ {
			_, _ = s.Dispatch(Event{
				Type: t1,
				Data: nil,
			})
		}
	})
}
