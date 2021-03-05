package rutina

import (
	"context"
	"sync"
)

// Rutina main object that rules coroutinas
type Rutina struct {
	Ctx              context.Context
	Cancel           context.CancelFunc
	wg               sync.WaitGroup
	Errors           chan error
	StopOnFirstError bool
}

// New returns new instance of Rutina
func New(ctx context.Context) *Rutina {
	ctx, cancel := context.WithCancel(ctx)
	return &Rutina{
		Ctx:    ctx,
		Cancel: cancel,
		wg:     sync.WaitGroup{},
		Errors: make(chan error),
	}
}

// Go runs new coroutine that managed with Rutina
func (r *Rutina) Go(f func(context.Context) error) context.CancelFunc {
	r.wg.Add(1)
	ctx, cancel := context.WithCancel(r.Ctx)
	go func() {
		defer r.wg.Done()
		if err := f(ctx); err != nil {
			r.Errors <- err
		}
	}()
	return cancel
}

// Wait all coroutines to complete
func (r *Rutina) Wait() (err error) {
	if r.StopOnFirstError {
		go func(err *error) {
			rerr := <-r.Errors
			err = &rerr
			r.Cancel()
		}(&err)
	}
	r.wg.Wait()
	return
}
