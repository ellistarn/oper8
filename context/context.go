package context

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	// AwaitInterrupt is a singleton context that is canceled when os.Interrupt or SIGTERM is received.
	AwaitInterrupt = new()
)

func new() context.Context {
	done := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, []os.Signal{os.Interrupt, syscall.SIGTERM}...)
	go func() {
		<-c
		close(done) // first signal, cancel context
		<-c
		os.Exit(1) // second signal, exit directly.
	}()
	return &singleton{done: done}
}

type singleton struct {
	done <-chan struct{}
}

func (s *singleton) Deadline() (time.Time, bool) {
	return time.Time{}, false
}

func (s *singleton) Done() <-chan struct{} {
	return s.done
}

func (s *singleton) Err() error {
	select {
	case _, ok := <-s.Done():
		if !ok {
			return context.Canceled
		}
	default:
	}
	return nil
}

func (s *singleton) Value(key interface{}) interface{} {
	return nil
}
