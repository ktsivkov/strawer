package sync

func (e *worker) Dispatch(event Event) (int, error) {
	subscribers, err := e.GetSubscribers(event.Type)
	if err != nil {
		return 0, err
	}
	for _, h := range subscribers {
		h.Handle(event)
	}
	return len(subscribers), nil
}
