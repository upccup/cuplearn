package main

import (
	"fmt"
	"log"
	"os"

	docker "github.com/fsouza/go-dockerclient"
)

func main() {
	client, err := docker.NewVersionedClient("unix:///var/run/docker.sock", "1.17")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fi, err := os.OpenFile("/data/test", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0420)
	if err != nil {
		log.Panicln(err)
	}

	cpOptions := docker.CopyFromContainerOptions{
		OutputStream: fi,
		Container:    "omega-abc",
		Resource:     "/data/test/",
	}

	err = client.CopyFromContainer(cpOptions)
	if err != nil {
		log.Println("copy failed")
		log.Panicln(err)
	}

	options := docker.DownloadFromContainerOptions{
		OutputStream: fi,
		Path:         "/data/test/",
	}

	err = client.DownloadFromContainer("omega-abc", options)
	if err != nil {
		log.Panicln(err)
	}
}
