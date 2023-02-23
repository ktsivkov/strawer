package sync

type worker struct {
	subs map[Type][]Handler
}
