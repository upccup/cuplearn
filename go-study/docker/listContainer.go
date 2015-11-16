package main

import (
	"fmt"
	docker "github.com/fsouza/go-dockerclient"
)

func main() {
	lient, err := docker.NewVersionedClient("unix:///var/run/docker.sock", "1.17")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	/*
		Parameters:

		all – 1/True/true or 0/False/false, Show all containers. Only running containers are shown by default (i.e., this defaults to false)
		limit – Show limit last created containers, include non-running ones.
		since – Show only containers created since Id, include non-running ones.
		before – Show only containers created before Id, include non-running ones.
		size – 1/True/true or 0/False/false, Show the containers sizes
		filters - a JSON encoded value of the filters (a map[string][]string) to process on the containers list. Available filters:
		exited=<int>; – containers with exit code of <int> ;
		status=(created|restarting|running|paused|exited)
		label=key or label="key=value" of a container label
	*/

	/*
		type ListContainersOptions struct {
		    All     bool
		    Size    bool
		    Limit   int
		    Since   string
		    Before  string
		    Filters map[string][]string
		}
	*/

	listOption := docker.ListContainersOptions{All: true, Filters: map[string][]string{"label": {"dataman"}}}

	containers, err := client.ListContainers(listOption)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(len(containers))

	fmt.Println(containers[0].Names[0])
}
