package async

type NoSubscribersError struct {
}

func (e *NoSubscribersError) Error() string {
	return "no subscribers found"
}

func (e *worker) GetSubscribers(name Type) ([]Handler, error) {
	val, ok := e.subs[name]
	if ok {
		return val, nil
	}
	return []Handler{}, &NoSubscribersError{}
}
