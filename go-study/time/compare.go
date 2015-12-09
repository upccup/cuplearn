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
	fmt.Println(befroe, "is after ", after, " ?", befroe.After(after))
}
