// Package coroutine 协程处理程序
// 处理等待协程等
package coroutine

import (
	"reflect"
	"sync"

	"github.com/freshcn/log"
)

var (
	wait sync.WaitGroup
)

// Add 添加一个新的等待项
func Add(nums ...int) {
	var num = 1
	if len(nums) > 0 {
		num = nums[0]
	}
	wait.Add(num)
}

// Done 完成一个等待
func Done() {
	wait.Done()
}

// Wait 等候完成
func Wait() {
	wait.Wait()
}

// Run 运行一个协程
func Run(handler interface{}, parames ...interface{}) bool {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
		}
	}()

	var (
		handlerFUNC    = reflect.ValueOf(handler)
		handlerParames = make([]reflect.Value, len(parames))
	)

	if handlerFUNC.Kind() != reflect.Func {
		return false
	}

	if len(parames) > 0 {
		for i, v := range parames {
			handlerParames[i] = reflect.ValueOf(v)
		}
	}

	// 开始运行一个新的协程
	Add()
	go func() {
		defer Done()
		handlerFUNC.Call(handlerParames)
	}()

	return true
}
