package main

import (
	"fmt"
	"sync"
	"time"
)

// ⽤用 closed channel 发出退出通知

func main() {
	var wg sync.WaitGroup
	quit := make(chan bool)

	for i := 0; i < 2; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			task :=
				func() {
					fmt.Println(id, time.Now().Nanosecond())
					time.Sleep(time.Second)
				}

			for {
				select {
				case <-quit: // closed channel 不会阻塞, 因此可以作为退出通知
					return
				default:
					task()
				}

			}
		}(i)
	}

	time.Sleep(time.Second * 5)
	close(quit) // 关闭channel 发出退出通知
	wg.Wait()
}
