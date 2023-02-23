package sync

import "testing"

func TestWorker_Subscribe(t *testing.T) {
	t.Run("Single subscription", func(t *testing.T) {
		t.Run("No target", func(t *testing.T) {
			t.Parallel()
			s := New()

			s.Subscribe(&testHandler{
				called: false,
			}, []Type{})
		})

		t.Run("Single target", func(t *testing.T) {
			t.Parallel()
			s := New()

			s.Subscribe(&testHandler{
				called: false,
			}, []Type{"test"})
		})

		t.Run("Multiple targets", func(t *testing.T) {
			t.Parallel()
			s := New()

			t1 := Type("test1")
			t2 := Type("test2")

			s.Subscribe(&testHandler{
				called: false,
			}, []Type{t1, t2})
		})
	})

	t.Run("Multiple subscriptions", func(t *testing.T) {
		t.Run("No target", func(t *testing.T) {
			t.Parallel()
			s := New()

			s.Subscribe(&testHandler{
				called: false,
			}, []Type{})

			s.Subscribe(&testHandler{
				called: false,
			}, []Type{})
		})

		t.Run("Single target", func(t *testing.T) {
			t.Parallel()
			s := New()

			s.Subscribe(&testHandler{
				called: false,
			}, []Type{"test"})

			s.Subscribe(&testHandler{
				called: false,
			}, []Type{"test"})
		})

		t.Run("Multiple targets", func(t *testing.T) {
			t.Parallel()
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
	})
}
