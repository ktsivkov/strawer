package sync

import "testing"

func TestWorker_Subscribe(t *testing.T) {
	t.Run("Single subscription - No target", func(t *testing.T) {
		s := New()

		s.Subscribe(&testHandler{
			called: false,
		}, []Type{})
	})

	t.Run("Single subscription - Single target", func(t *testing.T) {
		s := New()

		s.Subscribe(&testHandler{
			called: false,
		}, []Type{"test"})
	})

	t.Run("Single subscription - Multiple targets", func(t *testing.T) {
		s := New()

		t1 := Type("test1")
		t2 := Type("test2")

		s.Subscribe(&testHandler{
			called: false,
		}, []Type{t1, t2})
	})

	t.Run("Multiple subscriptions - No target", func(t *testing.T) {
		s := New()

		s.Subscribe(&testHandler{
			called: false,
		}, []Type{})

		s.Subscribe(&testHandler{
			called: false,
		}, []Type{})
	})

	t.Run("Multiple subscriptions - Single target", func(t *testing.T) {
		s := New()

		s.Subscribe(&testHandler{
			called: false,
		}, []Type{"test"})

		s.Subscribe(&testHandler{
			called: false,
		}, []Type{"test"})
	})

	t.Run("Multiple subscriptions - Multiple targets", func(t *testing.T) {
		s := New()

		t1 := Type("test1")
		t2 := Type("test2")

		s.Subscribe(&testHandler{
			called: false,
		}, []Type{t1, t2})

		s.Subscribe(&testHandler{
			called: false,
		}, []Type{t1, t2})
	})
}
