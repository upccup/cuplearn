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

// Output
// I:   0
// j:   0
// k:   0
// I:   0
// k:   1
// j:   1
// I:   1
// I:   1
// j:   2
// I:   2
// I:   2
// k:   2
// j:   3
// I:   3
// k:   3
// I:   3
// I:   4
// context ***I*** done:   context canceled
// context ***K*** done:   context canceled
// context ***J*** done:   context canceled
// context parent done:   context canceled
