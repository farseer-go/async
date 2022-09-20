/*
作者：黄山鸣
联系：1581400575@qq.com
创建时间：2022/09/18
修订时间：2022/09/18
*/
package async

import (
	"sync"
)

// Async 异步结构体
type Async struct {
	wg          *sync.WaitGroup // sync.WaitGroup
	fnCount     int             // 所有需要执行的异步方法
	fnDoneCount int             // 当前已执行完成的异步方法
	Err         any             // 返回错误
	CallBack    []func()        // 回调方法
}

// ContinueWith 回调方法，Add的所有方法执行结束后执行回调方法
func (ac *Async) ContinueWith(callbacks ...func()) {
	go ac.wg.Wait()
	ac.CallBack = callbacks
}

// Add 添加异步执行的方法
func (ac *Async) Add(fns ...func()) *Async {
	for _, fn := range fns {
		ac.wg.Add(1)
		ac.fnCount++
		go func(nfn func()) {
			defer func() {
				if err := recover(); err != nil {
					// 打印异常，关闭资源，退出此函数
					ac.Err = err
				}
				ac.fnDoneCount++
				if ac.fnDoneCount < ac.fnCount {
					defer ac.wg.Done()
				} else if ac.CallBack != nil {
					callBackFns := ac.CallBack
					go func() {
						defer func() {
							if err := recover(); err != nil {
								// 打印异常，关闭资源，退出此函数
								ac.Err = err
							}
							defer ac.wg.Done()
						}()
						for _, fn := range callBackFns {
							fn()
						}
					}()
				} else {
					defer ac.wg.Done()
				}
			}()
			nfn()
		}(fn)
	}
	return ac
}

// New 初始化
func New() *Async {
	return &Async{
		wg: &sync.WaitGroup{},
	}
}

// Wait 等待
func (ac *Async) Wait() {
	ac.wg.Wait()
}
