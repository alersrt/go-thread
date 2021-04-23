package pool_executor

type PoolExecutor struct {

	// The max amount of concurrent tasks.
	capacity int

	// Amount of the currently executed tasks.
	threads int

	// Amount tasks left for execution.
	counter int

	// Socket for signalling about finishing of the task execution.
	socket chan int

	// Tasks for execution.
	tasks []func()

	// Is all tasks finished.
	isFinished bool
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
func (executor *PoolExecutor) Execute() {
	// The listening should be started before
	// to avoid lack of performance.
	go executor.startListening()
	go executor.startExecution()

	for !executor.isFinished {
	}
}
