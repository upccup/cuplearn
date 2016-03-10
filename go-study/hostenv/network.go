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

// out put
/*
{"name":"eth0","bytes_sent":93368314,"bytes_recv":534584077,"packets_sent":244595,"packets_recv":591808,"errin":0,"errout":0,"dropin":2,"dropout":0}
{"name":"docker0","bytes_sent":4063,"bytes_recv":356,"packets_sent":24,"packets_recv":5,"errin":0,"errout":0,"dropin":0,"dropout":0}
{"name":"lo","bytes_sent":82727242,"bytes_recv":82727242,"packets_sent":823762,"packets_recv":823762,"errin":0,"errout":0,"dropin":0,"dropout":0}
{"protocol":"ip","stats":{"DefaultTTL":64,"ForwDatagrams":0,"Forwarding":1,"FragCreates":0,"FragFails":0,"FragOKs":0,"InAddrErrors":36,"InDelivers":1233602,"InDiscards":0,"InHdrErrors":0,"InReceives":1233697,"InUnknownProtos":0,"OutDiscards":48,"OutNoRoutes":712,"OutRequests":1039802,"ReasmFails":0,"ReasmOKs":0,"ReasmReqds":0,"ReasmTimeout":0}}
{"protocol":"icmp","stats":{"InAddrMaskReps":0,"InAddrMasks":0,"InCsumErrors":0,"InDestUnreachs":1293,"InEchoReps":4,"InEchos":3,"InErrors":596,"InMsgs":1300,"InParmProbs":0,"InRedirects":0,"InSrcQuenchs":0,"InTimeExcds":0,"InTimestampReps":0,"InTimestamps":0,"OutAddrMaskReps":0,"OutAddrMasks":0,"OutDestUnreachs":1298,"OutEchoReps":3,"OutEchos":4,"OutErrors":0,"OutMsgs":1305,"OutParmProbs":0,"OutRedirects":0,"OutSrcQuenchs":0,"OutTimeExcds":0,"OutTimestampReps":0,"OutTimestamps":0}}
{"protocol":"icmpmsg","stats":{"InType0":4,"InType3":1293,"InType8":3,"OutType0":3,"OutType3":1298,"OutType8":4}}
{"protocol":"tcp","stats":{"ActiveOpens":9003,"AttemptFails":1173,"CurrEstab":43,"EstabResets":82,"InCsumErrors":0,"InErrs":13,"InSegs":1151516,"MaxConn":-1,"OutRsts":1374,"OutSegs":1046364,"PassiveOpens":4024,"RetransSegs":3525,"RtoAlgorithm":1,"RtoMax":120000,"RtoMin":200}}
{"protocol":"udp","stats":{"IgnoredMulti":44791,"InCsumErrors":0,"InDatagrams":37035,"InErrors":0,"NoPorts":105,"OutDatagrams":16000,"RcvbufErrors":0,"SndbufErrors":0}}
{"protocol":"udplite","stats":{"IgnoredMulti":0,"InCsumErrors":0,"InDatagrams":0,"InErrors":0,"NoPorts":0,"OutDatagrams":0,"RcvbufErrors":0,"SndbufErrors":0}}
{"fd":12,"family":2,"type":2,"localaddr":{"ip":"*","port":5353},"remoteaddr":{"ip":"","port":0},"status":"","pid":587}
{"fd":13,"family":10,"type":2,"localaddr":{"ip":"*","port":5353},"remoteaddr":{"ip":"","port":0},"status":"","pid":587}
{"fd":14,"family":2,"type":2,"localaddr":{"ip":"*","port":36280},"remoteaddr":{"ip":"","port":0},"status":"","pid":587}
...
*/
