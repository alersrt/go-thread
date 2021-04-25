package pool_executor

type PoolExecutor struct {
	capacity   int      // The max amount of concurrent tasks.
	threads    int      // Amount of the currently executed tasks.
	counter    int      // Amount tasks left for execution.
	socket     chan int // Socket for signalling about finishing of the task execution.
	tasks      []func() // Tasks for execution.
	isFinished bool     // Is all tasks finished.
}

// CreatePoolExecutor created new instance of the pool executor
// with set of tasks for execution. Capacity parameter defines
// the amount of tasks what will be executed at the one moment.
// Zero value of capacity means the unlimited amount.
func CreatePoolExecutor(tasks []func(), capacity int) *PoolExecutor {
	return &PoolExecutor{
		tasks:    tasks,
		capacity: capacity,
		counter:  len(tasks),
		socket:   make(chan int),
	}
}

func (executor *PoolExecutor) startExecution() {
	for i, t := range executor.tasks {
		if executor.capacity != 0 {
			for executor.threads >= executor.capacity {
			}
		}
		executor.threads++
		i := i
		t := t
		go func() {
			t()
			executor.socket <- i
		}()
	}
}

func (executor *PoolExecutor) startListening() {
	for executor.counter > 0 {
		select {
		case _ = <-executor.socket:
			executor.counter--
			executor.threads--
		}
	}
	executor.isFinished = true
}

// Execute Executes set of passed tasks.
func (executor *PoolExecutor) Execute() bool {
	// The listening should be started before
	// to avoid lack of performance.
	go executor.startListening()
	go executor.startExecution()

	for !executor.isFinished {
	}

	return executor.isFinished
}
