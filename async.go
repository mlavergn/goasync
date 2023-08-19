package goasync

import "context"

// Async task interface
type Task interface {
	Await() (interface{}, error)
}

// Go channel context wrapper
type async struct {
	await func(ctx context.Context) (interface{}, error)
}

// Generic async function wrapper
func (any async) Await() (interface{}, error) {
	return any.await(context.Background())
}

// Async task factory
func Async(any func() (interface{}, error)) Task {
	var result interface{}
	ch := make(chan struct{})
	go func() {
		defer close(ch)
		result, _ = any()
	}()
	return async{
		await: func(ctx context.Context) (interface{}, error) {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-ch:
				return result, nil
			}
		},
	}
}
