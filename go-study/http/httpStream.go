package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func test() {
	var client *http.Client
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client = &http.Client{tr, nil, nil, 0 * time.Second}
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		fmt.Println("response: ", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("response2: ", err)
		return
	}

	reader := bufio.NewReader(resp.Body)
	defer resp.Body.Close()
	//delim := []byte("\n")
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("response3: ", err)
			return
		} else {
			fmt.Printf("resp:%s\n", line)
		}
	}

}

func main() {
	listenToMarathonEventStream()
}

func listenToMarathonEventStream() {
	client := &http.Client{}
	client.Timeout = 0 * time.Second
	ticker := time.NewTicker(1 * time.Second)
	eventsURL := "http://10.3.10.32:8080" + "/v2/events"
	for _ = range ticker.C {
		req, err := http.NewRequest("GET", eventsURL, nil)
		req.Header.Set("Accept", "text/event-stream")
		if err != nil {
			errorMsg := "An error occurred while creating request to Marathon events system: %s\n"
			log.Printf(errorMsg, err)
			continue
		}
		fmt.Println("aaasdaaaaa")
		resp, err := client.Do(req)
		if err != nil {
			errorMsg := "An error occurred while making a request to Marathon events system: %s\n"
			log.Printf(errorMsg, err)
			continue
		}

		defer resp.Body.Close()

		reader := bufio.NewReader(resp.Body)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					errorMsg := "An error occurred while reading Marathon event: %s\n"
					log.Printf(errorMsg, err)
				}
				break
			}

			if len(strings.TrimSpace(line)) == 0 {
				continue
			}

			if !strings.HasPrefix(line, "data: ") {
				errorMsg := "Wrong event format: %s\n"
				log.Printf(errorMsg, line)
				continue
			}

			line = line[6:]
			fmt.Println(line)
			// sub.Notify([]byte(line))
		}

		log.Println("Event stream connection was closed. Re-opening...")
	}

}
