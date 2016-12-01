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

	intCh := make(chan int)
	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			select {
			case value, ok := <-intCh:
				if ok {
					fmt.Println("value: ", value, "ok: ", ok)
				} else {
					fmt.Println("value: ", value, "ok: ", ok, "will return")
					return
				}
			}
		}
	}()

	intCh <- 999
	intCh <- 1000

	close(intCh)
	wg.Wait()
}

// output
// 0 654582000
// 1 654644000
// 1 659220838
// 0 659227862
// 0 664065318
// 1 664065280
// 1 668047368
// 0 668056051
// 0 673026314
// 1 673030808
// value:  999 ok:  true
// value:  1000 ok:  true
// value:  0 ok:  false will return
