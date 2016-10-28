package main

import (
	"fmt"
	"github.com/shirou/gopsutil/host"
	"os/exec"
	"strings"
)

//获取CentOs Ubuntu系统版本信息
func getUname() (uname string, err error) {
	f, err := exec.LookPath("lsb_release")
	if err != nil {
		uname, _ = getMacUname()
		return uname, nil
	}
	argv := []string{"-a"}
	c := exec.Command(f, argv...)
	unames, _ := c.Output()
	items := strings.Split(strings.TrimSpace(string(unames)), "\n")
	m := make(map[string]string)
	for _, item := range items {
		kv := strings.Split(item, ":")
		m[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
	}
	uname = m["Distributor ID"] + " " + m["Release"]
	return uname, nil
}

// 获取Os X的系统版本
func getMacUname() (u string, err error) {
	f, err := exec.LookPath("sw_vers")
	if err != nil {
		return "", err
	}
	argv1 := []string{"-productName"}
	argv2 := []string{"-productVersion"}

	productName, _ := exec.Command(f, argv1...).Output()
	productVersion, _ := exec.Command(f, argv2...).Output()

	u = strings.TrimSpace(string(productName)) + " " + strings.TrimSpace(string(productVersion))
	return u, nil
}

func main() {
	uname, err := getUname()
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println("uname: ", uname)

	hostinfo := host.HostInfo()

	fmt.Println(hostinfo.String())
}
