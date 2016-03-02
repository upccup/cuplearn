package main

import (
	"log"
	"os/exec"
)

func main() {

	cmd := exec.Command("tail", "-f", "/var/log/omega/agent.log")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Printf("Error: %s\n", err)
		return
	}
	if err = cmd.Start(); err != nil {
		log.Printf("Error: %s\n", err)
	}
	go func() {
		for {
			bytes := make([]byte, 1024)
			count, err := stdout.Read(bytes)
			//test in ArchLinux
			if err != nil {
				log.Println(err)
			}
			log.Println(count)
			log.Printf("The output is: %s\n", string(bytes)) //The output is: Hello World!
		}

	}()

	if err = cmd.Wait(); err != nil {
		log.Printf("Error: %s\n", err)
	}

}
