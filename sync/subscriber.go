package sync

func (e *worker) Subscribe(handler Handler, targets []Type) {
	for _, eType := range targets {
		subscribers, _ := e.GetSubscribers(eType)
		e.subs[eType] = append(subscribers, handler)
	}
}
