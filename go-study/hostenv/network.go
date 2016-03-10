package main

import (
	"log"

	"github.com/shirou/gopsutil/net"
)

func main() {
	netIOCounterSt, err := net.NetIOCounters(true)

	if err != nil {
		log.Panicln(err)
	}

	for _, netIOSt := range netIOCounterSt {
		log.Println(netIOSt)
	}

	netProtoCounters, err := net.NetProtoCounters(nil)
	if err != nil {
		log.Panicln(err)
	}

	for _, netProtoSt := range netProtoCounters {
		log.Println(netProtoSt)
	}

	netConnectionStat, err := net.NetConnections("all")
	if err != nil {
		log.Panicln(err)
	}

	for _, netConnSt := range netConnectionStat {
		log.Println(netConnSt)
	}

}
