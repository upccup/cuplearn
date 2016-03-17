package main

import (
	"log"

	"github.com/shirou/gopsutil/mem"
)

func main() {
	memSt, err := mem.VirtualMemory()
	if err != nil {
		log.Panicln(err)
	}

	log.Println("Total: ", memSt.Total)
	log.Println("Free: ", memSt.Free)
	log.Println("Used:", memSt.Used)
}
