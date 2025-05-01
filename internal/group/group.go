package group

import (
	"context"
	"errors"
	"sync"
)

type Group struct {
	wg     sync.WaitGroup
	errCh  chan error
	cancel context.CancelFunc
}

func NewWithContext(ctx context.Context) (*Group, context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	return &Group{
		errCh:  make(chan error),
		cancel: cancel,
	}, ctx
}

func (g *Group) Go(f func() error) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		if err := f(); err != nil {
			g.cancel()
			g.errCh <- err
		}
	}()
}

func (g *Group) Wait() error {
	go func() {
		defer close(g.errCh)
		g.wg.Wait()
	}()

	var err error
	for nextErr := range g.errCh {
		err = errors.Join(err, nextErr)
	}

	return err
}
