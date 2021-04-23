package main

type PoolExecutor struct {
	socket chan bool
}

func CreatePoolExecutor() *PoolExecutor {
	return &PoolExecutor{socket: make(chan bool)}
}

// Execute Executes set of passed tasks and handles results by passed handler.
func (executor *PoolExecutor) Execute(tasks []func()) {
	counter := len(tasks)

	for _, t := range tasks {
		go func() {
			t()
			executor.socket <- true
		}()
	}

	for counter > 0 {
		select {
		case _ = <-executor.socket:
			counter--
		}
	}
}
