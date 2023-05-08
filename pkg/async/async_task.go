package async

// AsynchronousTask describes task with asynchronous
// execution where.
type AsynchronousTask[T any] struct {
	f func() T // f Common function what should be executed.
}

// NewTask Creates new AsynchronousTask instance.
func NewTask[T any](f func() T) *AsynchronousTask[T] {
	return &AsynchronousTask[T]{f: f}
}

// Result Returns the channel for awaiting of the executed
// function result.
func (t *AsynchronousTask[T]) Result() <-chan T {
	c := make(chan T)
	go func() {
		defer close(c)
		res := t.f()
		c <- res
	}()
	return c
}
