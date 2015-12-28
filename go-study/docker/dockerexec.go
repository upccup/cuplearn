package main

import (
	"bytes"
	"fmt"
	docker "github.com/fsouza/go-dockerclient"
	"log"
	"time"
)

func main() {
	client, err := docker.NewVersionedClient("unix:///var/run/docker.sock", "1.17")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	execOptions := docker.CreateExecOptions{
		AttachStdout: true,
		AttachStdin:  false,
		Cmd:          []string{"date"},
		Container:    "omega-abc",
	}

	exec, err := client.CreateExec(execOptions)
	if err != nil {
		log.Fatalln("create exec has error: ", err)
	}

	var squidLogs bytes.Buffer
	err = client.StartExec(exec.ID, docker.StartExecOptions{
		OutputStream: &squidLogs,
	})
	if err != nil {
		log.Fatal(err)
	}
	for {
		tmp, _ := client.InspectExec(exec.ID)
		if !tmp.Running {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	log.Println(squidLogs.String())

}
