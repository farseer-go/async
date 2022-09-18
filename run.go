/*
作者：黄山鸣
联系：1581400575@qq.com
创建时间：2022/09/18
修订时间：2022/09/18
*/
package async

import "sync"

// Async 异步结构体，wg：sync.WaitGroup，fnCount为所有需要执行的异步方法，fnDoneCount为当前已执行完成的异步方法
type Async struct {
	wg          *sync.WaitGroup
	fnCount     int
	fnDoneCount int
}

// Callback 回调方法，Add的所有方法执行结束后执行回调方法
func (ac *Async) Callback(fns ...func()) *Async {
	go func(*Async, *[]func()) {
		for ac.fnDoneCount != ac.fnCount {
		}
		for _, fn := range fns {
			fn()
		}
		ac.wg.Done()
	}(ac, &fns)
	return ac
}

// Add 添加异步执行的方法
func (ac *Async) Add(fns ...func()) *Async {
	for _, fn := range fns {
		ac.wg.Add(1)
		ac.fnCount++
		go func(ac *Async, nfn func()) {
			nfn()
			ac.fnDoneCount++
			if ac.fnDoneCount < ac.fnCount {
				defer ac.wg.Done()
			}
		}(ac, fn)
	}
	return ac
}

// New 初始化异步执行结构体asyncStruct
func New() *Async {
	asyncStruct := &Async{}
	wg := sync.WaitGroup{}
	asyncStruct.wg = &wg
	return asyncStruct
}

// Run 执行异步方法
func (ac *Async) Run() *Async {
	ac.wg.Wait()
	return ac
}
