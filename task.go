package main

// Task Representation of the task for execution.
type Task interface {

	// Result Returns result of the execution.
	Result() interface{}

	// Error Returns error what is happened during execution.
	Error() error
}
