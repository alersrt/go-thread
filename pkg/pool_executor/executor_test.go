package pool_executor

import (
	"fmt"
	guuid "github.com/google/uuid"
	"testing"
	"time"
)

func TestExecute(t *testing.T) {
	// Arranges

	socket := make(chan string)
	tMap := make(map[string]bool, 10)
	tSlice := make([]func(), 10)

	testFunc := func(uuid string) string {
		time.Sleep(2 * time.Second)
		return uuid
	}

	for i := 0; i < 10; i++ {
		uuid, _ := guuid.NewRandom()
		tMap[uuid.String()] = false
		tSlice[i] = func() {
			u := testFunc(uuid.String())
			socket <- u
		}
	}

	// Actions

	poolExecutor := CreatePoolExecutor(tSlice, 4)
	isFinished := false
	go func() {
		for !isFinished {
			select {
			case u := <-socket:
				tMap[u] = true
			}
		}
	}()
	isFinished = poolExecutor.Execute()

	// Asserts

	isTrue := true
	for k, v := range tMap {
		fmt.Printf("%s %t\n", k, v)
		isTrue = isTrue && v
	}

	if !isTrue {
		t.Fail()
	}
}
