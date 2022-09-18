/*
作者：黄山鸣
联系：1581400575@qq.com
创建时间：2022/09/18
修订时间：2022/09/18
*/
package async

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var TestVal int = 0

// Async 异步结构体，wg：sync.WaitGroup，fnCount为所有需要执行的异步方法，fnDoneCount为当前已执行完成的异步方法
func callback() {
	TestVal = 10
}
func add() {
	TestVal = 1
}
func TestAsync_Run(t *testing.T) {
	New().Add(func() {
		add()
	}, func() {
		add()
	}).Callback(func() {
		callback()
	}).Run()
	assert.Equal(t, TestVal, 10)
	New().Add(func() {
		callback()
	}, func() {
		callback()
	}).Callback(func() {
		add()
	}).Run()
	assert.Equal(t, TestVal, 1)
}
