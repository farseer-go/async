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
	"sync"
	"testing"
	"time"
)

func TestAsync_ContinueWith(t *testing.T) {
	var lock sync.Mutex
	var count = 0
	lock.Lock()
	async.Parallel(
		func() {
			lock.Lock()
			defer lock.Unlock()
			count += 1
		},
		func() {
			lock.Lock()
			defer lock.Unlock()
			count += 2
		}).
		Add(func() {
			lock.Lock()
			defer lock.Unlock()
			count += 3
		}, func() {
			lock.Lock()
			defer lock.Unlock()
			count += 4
		}).
		ContinueWith(func() {
			lock.Lock()
			defer lock.Unlock()
			count += 5
		})

	count = 10
	lock.Unlock()
	time.Sleep(10 * time.Millisecond)

	lock.Lock()
	defer lock.Unlock()
	assert.Equal(t, 25, count)
}

func TestAsync_Wait(t *testing.T) {
	var lock sync.Mutex
	var count = 0
	_ = async.Parallel(func() {
		lock.Lock()
		defer lock.Unlock()
		count += 1
	}, func() {
		lock.Lock()
		defer lock.Unlock()
		count += 2
	}).Add(func() {
		lock.Lock()
		defer lock.Unlock()
		count += 3
	}, func() {
		lock.Lock()
		defer lock.Unlock()
		count += 4
	}).Wait()
	count *= 2
	assert.Equal(t, 20, count)
}

func TestAsync_Error(t *testing.T) {
	var count = 0
	var num = 0
	err := async.Parallel().Add(func() {
		count = count / num
	}).Wait()
	assert.NotEqual(t, err, nil)

	err = async.Parallel().Add(func() {
		panic("error")
	}).Wait()

	assert.NotEqual(t, err, nil)
}

func TestAsync_Parallel(t *testing.T) {
	async := async.Parallel()
	assert.NotEqual(t, async, nil)

}
