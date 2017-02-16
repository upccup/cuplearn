package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/net/context"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	ctx1, _ := context.WithCancel(ctx)
	ctx2, _ := context.WithCancel(ctx)
	ctx3, _ := context.WithCancel(ctx)

	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		i := 0
		ticker := time.NewTicker(time.Second * 1)

		go UseContext(ctx, 1)
		go UseContext(ctx, 2)
		for {
			select {
			case <-ticker.C:
				fmt.Println("I:  ", i)
				i++
			case <-ctx.Done():
				fmt.Println("context parent done:  ", ctx.Err())
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		i := 0
		ticker := time.NewTicker(time.Second * 1)
		for {
			select {
			case <-ticker.C:
				fmt.Println("I:  ", i)
				i++
			case <-ctx1.Done():
				fmt.Println("context ***I*** done:  ", ctx1.Err())
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		j := 0
		ticker := time.NewTicker(time.Second * 1)
		for {
			select {
			case <-ticker.C:
				fmt.Println("j:  ", j)
				j++
			case <-ctx2.Done():
				fmt.Println("context ***J*** done:  ", ctx2.Err())
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		k := 0
		ticker := time.NewTicker(time.Second * 1)
		for {
			select {
			case <-ticker.C:
				fmt.Println("k:  ", k)
				k++
			case <-ctx3.Done():
				fmt.Println("context ***K*** done:  ", ctx3.Err())
				return
			}
		}
	}()

	time.Sleep(time.Second * 5)
	cancel()

	wg.Wait()
}

func UseContext(ctx context.Context, i int) error {
	fmt.Printf("use  ctx at  %d time \n", i)
	select {
	case <-ctx.Done():
		fmt.Printf("ctx cancel at  %d time \n", i)
		return ctx.Err()
	}
}

/*
use  ctx at  2 time
use  ctx at  1 time
j:   0
k:   0
I:   0
I:   0
I:   1
j:   1
I:   1
k:   1
j:   2
I:   2
k:   2
I:   2
k:   3
j:   3
I:   3
I:   3
I:   4
k:   4
j:   4
I:   4
context parent done:   context canceled
ctx cancel at  2 time
ctx cancel at  1 time
context ***K*** done:   context canceled
context ***J*** done:   context canceled
context ***I*** done:   context canceled
*/

// 说明: 通过这个程序可以看到当cancel() 被调用后及时一个 ctx 在多个线程中呗使用所有的 ctx.Done()都能收到信号
// 不是只有其中一个能收到信号. 而且所有子Ctx 的 ctx.Done() 也会收到信号.
