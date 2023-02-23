package async

type worker struct {
	subs map[Type][]Handler
}
