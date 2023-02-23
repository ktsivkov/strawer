package async

type testHandler struct {
	called bool
}

func (h *testHandler) Handle(_ Event) {
	h.called = true
}
