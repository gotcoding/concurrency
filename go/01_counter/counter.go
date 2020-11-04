package counter

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

// NotSafeCounter 未加锁，多个goroutine同时修改一个值，结果不确定
func NotSafeCounter() {
	var count int64 = 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				count++
			}
		}()
	}
	wg.Wait()
	fmt.Println("NotSafeCounter: ", count)
}

// MutexCounter 使用Mutex管理锁
func MutexCounter() {
	var count int64 = 0
	var mu sync.Mutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println("MutexCounter: ", count)
}

// AtomicCounter 使用atomic原子操作
func AtomicCounter() {
	var count int64 = 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				atomic.AddInt64(&count, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("AtomicCounter: ", count)
}

// ChannelCounter 使用channel管理
func ChannelCounter() {
	ch := make(chan struct{}) // 发送和接收count++"信号"的通道
	var count int64 = 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				ch <- struct{}{}
			}
		}()
	}

	go func() {
		wg.Wait() // 等待上面所有的 goroutine 运行完成
		close(ch) // 关闭ch通道
	}()
	for range ch { // 如果ch通道读取完了(ch是关闭状态), 则for循环结束
		count++
	}
	fmt.Println("ChannelCounter: ", count)
}
