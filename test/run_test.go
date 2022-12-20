/*
作者：黄山鸣
联系：1581400575@qq.com
创建时间：2022/09/18
修订时间：2022/09/18
*/
package test

import (
	"github.com/farseer-go/async"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAsync_ContinueWith(t *testing.T) {
	var count = 0
	async.Parallel(func() {
		count += 1
	}, func() {
		count += 2
	}).Add(func() {
		count += 3
	}, func() {
		count += 4
	}).ContinueWith(func() {
		count += 5
	})
	count = 10
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, 25, count)
}

func TestAsync_Wait(t *testing.T) {
	var count = 0
	_ = async.Parallel(func() {
		count += 1
	}, func() {
		count += 2
	}).Add(func() {
		count += 3
	}, func() {
		count += 4
	}).Wait()
	count *= 2
	assert.Equal(t, 20, count)
}
