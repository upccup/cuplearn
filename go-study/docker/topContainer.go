package main

import (
	"fmt"
	docker "github.com/fsouza/go-dockerclient"
)

func main() {
	client, err := docker.NewVersionedClient("unix:///var/run/docker.sock", "1.17")
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := client.TopContainer("omega-cadvisor", "aux")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result.Titles)

	for _, process := range result.Processes {
		for i, processInfo := range process {
			fmt.Println(result.Titles[i], ":   ", processInfo)
		}
	}

	/*
		[USER PID %CPU %MEM VSZ RSS TTY STAT START TIME COMMAND]
		USER :    root
		PID :    13677
		%CPU :    1.7
		%MEM :    0.7
		VSZ :    408232
		RSS :    31728
		TTY :    ?
		STAT :    Ssl
		START :    Nov13
		TIME :    64:25
		COMMAND :    /usr/bin/cadvisor -logtostderr -port=9007
	*/
}
