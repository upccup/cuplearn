package main

import (
	"log"

	"github.com/shirou/gopsutil/disk"
)

func main() {
	diskList, err := disk.DiskPartitions(true)
	if err != nil {
		log.Panicln(err)
	}

	for _, diskPartition := range diskList {
		log.Println(diskPartition)
		useInfo, err := disk.DiskUsage(diskPartition.Mountpoint)
		if err != nil {
			log.Panicln(err)
		}
		if useInfo.Total == 0 {
			continue
		}
		log.Printf("%+v", useInfo)
	}
}
