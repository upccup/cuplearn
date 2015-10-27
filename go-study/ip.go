package main

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

func GetIp() (ip string, err error) {
	ifaces, err := net.Interfaces()

	if err != nil {
		return "", err
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}

		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			ip = ip.To4()

			if ip == nil {
				continue
			}

			return ip.String(), nil
		}
	}
	return "", nil
}

func GetIpByEnName(flagEnName string) (ip string, err error) {
	if len(flagEnName) > 0 {
		ifaceByName, err := net.InterfaceByName(flagEnName)
		if err != nil {
			return "", err
		}
		if !validInterface(ifaceByName) {
			return "", errors.New("invalid interface by EnName")
		}
		return getIpAddr(ifaceByName)
	}

	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if !validInterface(&iface) {
			continue
		}
		return getIpAddr(&iface)
	}
	return "", errors.New("are you connected to the network?")
}

func validInterface(iface *net.Interface) bool {
	if iface.Flags&net.FlagUp == 0 {
		return false // interface down
	}
	if iface.Flags&net.FlagLoopback != 0 {
		return false // loopback interface
	}
	return true
}

func getIpAddr(iface *net.Interface) (ip string, err error) {
	addrs, err := iface.Addrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}
		if ip == nil || ip.IsLoopback() {
			continue
		}
		ip = ip.To4()
		if ip == nil {
			continue // not an ipv4 address
		}
		return ip.String(), nil
	}
	return "", errors.New("are you connected to the network?")
}

func GetIpByClient() {
	ip2, err := GetIp()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(ip2)

	conn, err := net.Dial("udp", "baidu:80")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	fmt.Println(strings.Split(conn.LocalAddr().String(), ":")[0])

	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err.Error())
	}

	ips, err := net.LookupIP("localhost")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ips)

	var ip string

	for _, inter := range interfaces {
		iplist, err := inter.Addrs()
		if err != nil {
			fmt.Println(err.Error())
		}

		if inter.Name == "eth0" {
			if len(iplist) == 1 {
				ip = iplist[0].String()
				if err == nil {
					break
				}
			} else if len(iplist) == 2 {
				ip = iplist[1].String()
				if err == nil {
					break
				}
			}
		}
	}
	start := strings.Index(ip, "/")
	ip = ip[0:start]
	fmt.Println(ip)
}

func main() {
	ip, err := GetIpByEnName("awdl0")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("IP:  ", ip)
}
