package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 说明: 和协程 yield 作⽤用类似,Gosched 让出底层线程,将当前 goroutine 暂停,放回队列等 待下次被调度执⾏行

func main() {
	runtime.GOMAXPROCS(1)
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 6; i++ {
			fmt.Println(i)
			if i == 3 {
				runtime.Gosched()
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			if i == 3 {
				runtime.Gosched()
			}
			fmt.Println("Hell0 World!")
		}
	}()

	wg.Wait()
}

/*
Hell0 World!
Hell0 World!
Hell0 World!
0
1
2
3
Hell0 World!
Hell0 World!
Hell0 World!
Hell0 World!
Hell0 World!
Hell0 World!
Hell0 World!
4
5
*/
