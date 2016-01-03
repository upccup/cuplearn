package main

import (
	"fmt"
	"reflect"
)

type EventInfo struct {
	Id       string `json:"id"`
	Describe string `json:"describe"`
	Status   int    `json:"status"`
}

func main() {
	event := &EventInfo{}
	fmt.Println("before", event)
	val := reflect.ValueOf(event).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag
		switch valueField.Type().Kind() {
		case reflect.String:
			fmt.Println("string ===")
		case reflect.Int:
			fmt.Println("int =====")
		default:
			fmt.Println("aaaaa")
		}
		fmt.Printf("%+v", typeField.Type)
		if tag.Get("json") == "id" {
			valueField.SetString("111111")
		} else if tag.Get("json") == "describe" {
			valueField.SetString("aaaaaa")
		} else if tag.Get("json") == "status" {
			valueField.SetInt(10)
		}
	}

	fmt.Println("after: ", event)
}
