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

}
