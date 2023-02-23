package async

type Event struct {
	Type    Type
	Data    any
	Channel chan any
}
