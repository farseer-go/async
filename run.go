/*
作者：黄山鸣
联系：1581400575@qq.com
创建时间：2022/09/18
修订时间：2022/09/18
*/
package async

import (
	"errors"
	"fmt"
	"sync"
)

// Async 异步结构体
type Async struct {
	wg  *sync.WaitGroup // sync.WaitGroup
	err error           // 返回错误
}

// Parallel 并行执行fns
func Parallel(fns ...func()) *Async {
	async := &Async{
		wg: &sync.WaitGroup{},
	}
	if len(fns) > 0 {
		return async.Add(fns...)
	}
	return async
}

// Add 添加异步执行的方法
func (ac *Async) Add(fns ...func()) *Async {
	for _, fn := range fns {
		ac.wg.Add(1)
		go ac.executeFunc(fn)
	}
	return ac
}

func (ac *Async) executeFunc(fn func()) {
	defer func() {
		// 异常处理
		if err := recover(); err != nil {
			switch err.(type) {
			case error:
				ac.err = err.(error)
			default:
				ac.err = errors.New(fmt.Sprint(err))
			}
		}
		ac.wg.Done()
	}()
	fn()
}

// ContinueWith 当并行任务执行完后，以非阻塞方式执行callbacks
func (ac *Async) ContinueWith(callbacks ...func()) {
	// 使用异步等待，并执行callbacks
	go func() {
		ac.wg.Wait()
		for _, callback := range callbacks {
			callback()
		}
	}()
}

// Wait 阻塞等待执行完成
func (ac *Async) Wait() error {
	ac.wg.Wait()
	return ac.err
}
