package main

import (
	"encoding/json"
	"fmt"

	"github.com/influxdata/influxdb/client/v2"
)

type AlertEvent struct {
	TaskName  string `json:"taskname"`
	AlertTime int64  `json:"alerttime"`
	AppName   string `json:"appname"`
	ClusterId string `json:"clusterid"`
	Instance  string `json:"instance"`
	Level     string `json:"level"`
	Index     int64  `json:"index"`
}

// Make a Query
func ExampleClient_query() {
	// Make client
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://192.168.199.222:5008",
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer c.Close()

	//q := client.NewQuery("SELECT * FROM ALERT_EVENTS ORDER BY time DESC LIMIT 5 OFFSET 40", "shurenyun", "ns")
	q := client.NewQuery("SELECT * FROM ALERT_EVENTS", "shurenyun", "ns")
	if response, err := c.Query(q); err == nil && response.Error() == nil {
		//fmt.Printf("%+v", response.Results)
		bytes, err := json.Marshal(response.Results)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(bytes))

		if len(response.Results) < 1 {
			return
		}

		// alertEvents := []map[string]interface{}{}
		var alertEvents []map[string]interface{}

		series := response.Results[0].Series
		for _, serie := range series {
			alertEvent := make(map[string]interface{})
			columns := serie.Columns
			values := serie.Values
			fieldNum := len(columns)
			if fieldNum < 1 {
				break
			}

			for _, value := range values {
				for i := 0; i < fieldNum; i++ {
					if len(value) != fieldNum {
						break
					}
					alertEvent[columns[i]] = value[i]
				}
				alertEvents = append(alertEvents, alertEvent)

			}
		}
		fmt.Printf("%+v ", alertEvents)
	}
}

func main() {
	ExampleClient_query()
}
