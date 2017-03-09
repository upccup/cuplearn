package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		if val.Type().Kind() == reflect.String && structFieldType.Kind() == reflect.Uint32 {
			intValue, err := strconv.ParseUint(value.(string), 10, 64)
			if err != nil {
				return err
			}

			structFieldValue.Set(reflect.ValueOf(uint32(intValue)))
			fmt.Println("xxxxxxxxxxxxxxxxxxxx")
			return nil
		}
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}

type MyStruct struct {
	Name string
	Age  uint32
}

func (s *MyStruct) FillStruct(m map[string]interface{}) error {
	for k, v := range m {
		err := SetField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	myData := make(map[string]interface{})
	myData["Name"] = "Tony"
	myData["Age"] = uint32(23)

	result := &MyStruct{}
	err := result.FillStruct(myData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	result2 := &MyStruct{}
	myData["Age"] = "23"

	err = result2.FillStruct(myData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result2)
}
