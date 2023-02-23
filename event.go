package strawer

import (
	"github.com/ktsivkov/strawer/async"
	"github.com/ktsivkov/strawer/sync"
)

type Worker interface {
	Sync() sync.Worker
	Async() async.Worker
}

func New() Worker {
	return &worker{
		sync:  sync.New(),
		async: async.New(),
	}
}

type worker struct {
	sync  sync.Worker
	async async.Worker
}

func (e *worker) Sync() sync.Worker {
	return e.sync
}

func (e *worker) Async() async.Worker {
	return e.async
}
