package async

// AsynchronousTask describes task with asynchronous
// execution where.
type AsynchronousTask struct {
	// f Common function what should be executed.
	f func() interface{}
}

// NewTask Creates new AsynchronousTask instance.
func NewTask(f func() interface{}) *AsynchronousTask {
	return &AsynchronousTask{f: f}
}

// Result Returns the channel for awaiting of the executed
// function result.
func (t *AsynchronousTask) Result() <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		res := t.f()
		c <- res
	}()
	return c
}
