package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 说明: 调⽤用 runtime.Goexit 将⽴立即终⽌止当前 goroutine 执⾏行,调度器确保所有已注册 defer 延迟调⽤用被执⾏行

func main() {
	runtime.GOMAXPROCS(1)
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		for i := 0; i < 15; i++ {
			if i > 5 {
				runtime.Goexit()
			}
			fmt.Println(i)
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		defer wg.Done()
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			runtime.Goexit() // 立即终止当前线程
			fmt.Println("B") // 不会执行
		}()

		fmt.Println("A") // 不会执行
	}()
	wg.Wait()
	time.Sleep(time.Second * 10)
	fmt.Println("main")
}

/*
B.defer
A.defer
0
1
2
3
4
5
main
*/
