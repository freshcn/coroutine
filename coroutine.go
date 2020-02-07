// Package coroutine 协程处理程序
// 处理等待协程等
package coroutine

import (
	"reflect"
	"sync"

	"github.com/freshcn/log"
)

// Group 主处理组
type Group struct {
	wait sync.WaitGroup
}

// Default 默认的处理组
var Default = Group{}

// add 添加一个新的等待项
func (g *Group) add(nums ...int) {
	var num = 1
	if len(nums) > 0 {
		num = nums[0]
	}
	g.wait.Add(num)
}

// done 完成一个等待
func (g *Group) done() {
	g.wait.Done()
}

// Wait 等候完成
func (g *Group) Wait() {
	g.wait.Wait()
}

// Run 运行一个协程
// handler 是要运行的处理函数，可以接受多个参数
// parames 需要发给handler接受的参数，顺序和函数需要接受的顺序一样
func (g *Group) Run(handler interface{}, parames ...interface{}) bool {
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
	g.add()
	go func() {
		defer g.done()
		handlerFUNC.Call(handlerParames)
	}()

	return true
}
