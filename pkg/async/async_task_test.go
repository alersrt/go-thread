package async

import (
	"testing"
	"time"
)

func TestAsynchronousTask_Result(t *testing.T) {
	//	Arranges
	testFunc := func() any {
		time.Sleep(2 * time.Second)
		return true
	}

	//	Actions
	testTask := NewTask[any](testFunc)
	res := <-testTask.Result()

	//	Asserts
	if res == nil {
		t.Fail()
	}
	if res == 0 {
		t.Fail()
	}
	if res == false {
		t.Fail()
	}
}
