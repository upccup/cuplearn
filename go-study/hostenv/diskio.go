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

// output
/*
Name:  sda2
IOInfo:  {"read_count":2,"write_count":0,"read_bytes":2048,"write_bytes":0,"read_time":4,"write_time":0,"name":"sda2","io_time":4,"serial_number":"VBOX_HARDDISK_VB36b30180-ee3144aa"}
Name:  sda5
IOInfo:  {"read_count":10064,"write_count":24215,"read_bytes":70799360,"write_bytes":566939648,"read_time":5540,"write_time":11112,"name":"sda5","io_time":9772,"serial_number":"VBOX_HARDDISK_VB36b30180-ee3144aa"}
Name:  sr0
IOInfo:  {"read_count":22,"write_count":0,"read_bytes":71680,"write_bytes":0,"read_time":12,"write_time":0,"name":"sr0","io_time":12,"serial_number":"VBOX_CD-ROM_VB2-01700376"}
Name:  sda
IOInfo:  {"read_count":141221,"write_count":113236,"read_bytes":2258050048,"write_bytes":3677831168,"read_time":83656,"write_time":146684,"name":"sda","io_time":89964,"serial_number":"VBOX_HARDDISK_VB36b30180-ee3144aa"}
Name:  sda1
IOInfo:  {"read_count":130994,"write_count":88064,"read_bytes":2185913344,"write_bytes":3110891520,"read_time":78068,"write_time":134920,"name":"sda1","io_time":80092,"serial_number":"VBOX_HARDDISK_VB36b30180-ee3144aa"}
*/
