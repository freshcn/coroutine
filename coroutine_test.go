package coroutine

import (
	"testing"
	"time"
)

func TestDefault(t *testing.T) {
	for i := 0; i < 10; i++ {
		Default.Run(func(i int) {
			t.Log(i)
			time.Sleep(5 * time.Second)
		}, i)
	}
	Default.Wait()
}

func TestNewGroup(t *testing.T) {
	group := Group{}

	for i := 0; i < 10; i++ {
		group.Run(func(i int) {
			t.Log(i)
			time.Sleep(5 * time.Second)
		}, i)
	}
	group.Wait()
}
