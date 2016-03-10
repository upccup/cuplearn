package main

import (
	"log"

	"github.com/shirou/gopsutil/disk"
)

func main() {
	diskIOCounters, err := disk.DiskIOCounters()
	if err != nil {
		log.Panicln(err)
	}

	for name, IOInfo := range diskIOCounters {
		log.Println("Name: ", name)
		log.Println("IOInfo: ", IOInfo.String())
	}
}
