package main

import (
	"log"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func main() {
	cpuPercent, err := cpu.CPUPercent(time.Second*20, true)
	if err != nil {
		log.Panicln(err)
	}

	for _, cpu := range cpuPercent {
		log.Println(cpu)
	}
}

// out put
/*
27.03648175912044
1.4000000000000001
24.18790604697651
1.5492253873063468
23.45
1.5
23.976023976023978
1.7500000000000002
*/
