package async

import "sync"

type Async struct {
	wg          *sync.WaitGroup
	fnCount     int
	fnDoneCount int
}

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

func (ac *Async) Add(fns ...func()) *Async {
	for _, fn := range fns {
		ac.wg.Add(1)
		ac.fnCount++
		go func(*Async) {
			fn()
			ac.fnDoneCount++
			if ac.fnDoneCount < ac.fnCount {
				ac.wg.Done()
			}
		}(ac)
	}
	return ac
}

// Run 运行一个异步方法
// interval:任务运行的间隔时间
// taskFn:要运行的任务
func New() *Async {
	asyncStruct := &Async{}
	wg := sync.WaitGroup{}
	asyncStruct.wg = &wg
	return asyncStruct
}

func (ac *Async) Run() *Async {
	ac.wg.Wait()
	return ac
}
