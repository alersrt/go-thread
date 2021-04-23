package main

// CompletableTask Representation of the task for execution.
type CompletableTask interface {

	// Result Returns result of the execution.
	Result() interface{}

	// Error Returns error what is happened during execution.
	Error() error
}
