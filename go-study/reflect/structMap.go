package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type AlertEvent struct {
	TaskName  string `json:"taskname"`
	Timestamp int64  `json:"timestamp"`
	AppName   string `json:"appname"`
	ClusterId string `json:"clusterid"`
	Instance  string `json:"instance"`
}

func ConvertEventToMap(event *AlertEvent) map[string]string {
	val := reflect.ValueOf(event).Elem()

	eventMap := make(map[string]string)
	var str string
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		jsonTag := typeField.Tag.Get("json")
		if jsonTag == "" {
			continue
		}

		valueFiled := val.Field(i).Interface()
		switch value := valueFiled.(type) {
		case int64:
			str = strconv.FormatInt(value, 10)
		case string:
			str = value
		default:
			str = ""
		}

		eventMap[jsonTag] = str
	}

	return eventMap
}

func main() {
	event := AlertEvent{
		TaskName:  "aaaa",
		Timestamp: 1231231231231,
		AppName:   "app123",
		ClusterId: "123",
		Instance:  "123.2.1.3",
	}

	cMap := ConvertEventToMap(&event)

	fmt.Printf("%+v", cMap)
}
