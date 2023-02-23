package async

import "sync"

func (e *worker) Dispatch(wg *sync.WaitGroup, event Event) (int, error) {
	subscribers, err := e.GetSubscribers(event.Type)
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
