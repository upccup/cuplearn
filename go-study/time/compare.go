package main

import (
	"fmt"
	"time"
)

// compare two time

func main() {
	now := time.Now()
	after := now.Add(1 * time.Minute)
	befroe := now.Add(-1 * time.Minute)

	fmt.Println(befroe, "is bfefroe ", after, " ?", befroe.Before(after))
	// 2015-12-09 22:31:25.361393411 +0800 CST is bfefroe  2015-12-09 22:33:25.361393411 +0800 CST  ? true
	fmt.Println(befroe, "is after ", after, " ?", befroe.After(after))
	// 2015-12-09 22:31:25.361393411 +0800 CST is after  2015-12-09 22:33:25.361393411 +0800 CST  ? false
}
