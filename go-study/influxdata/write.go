package main

import (
	"log"
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

func GetInfluxClient() (client.Client, error) {
	return client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://192.168.199.222:5008",
		Username: "root",
		Password: "root",
	})
}

func WriteInflux() {
	c, err := GetInfluxClient()
	if err != nil {
		log.Println("Create influx client got err: ", err)
		return
	}
	defer c.Close()

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "shurenyun",
		Precision: "s",
	})

	if err != nil {
		log.Println(err)
	}

	// Create a point and add to batch
	tags := map[string]string{"cpu": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
	}
	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		log.Println("Error: ", err.Error())
	}
	bp.AddPoint(pt)

	// Write the batch
	c.Write(bp)
	return
}

func WriteInflux2() {
	c, err := GetInfluxClient()
	if err != nil {
		log.Println("Create influx client got err: ", err)
		return
	}
	defer c.Close()

	// Create a new point batch
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "shurenyun",
		Precision: "s",
	})

	// Create a point and add to batch
	tags := map[string]string{"cpu": "cpu-total", "disk": "disk-uesd"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
		"disk":   46.6,
		"net":    46.6,
	}
	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		log.Println("Error: ", err.Error())
	}
	bp.AddPoint(pt)

	// Write the batch
	c.Write(bp)
	return
}

func main() {
	WriteInflux()
	WriteInflux2()
}
