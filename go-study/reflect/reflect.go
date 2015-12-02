package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type MyType struct {
	i    int
	name string
}

func (mt *MyType) SetI(i int) {
	mt.i = i
}

func (mt *MyType) SetName(name string) {
	mt.name = name
}

func (mt *MyType) String() string {
	return fmt.Sprintf("%p", mt) + "--name:" + mt.name + "i:" + strconv.Itoa(mt.i)
}

func SetMyType(mt *MyType) {
	mt.SetI(1)
	mt.SetName("1111")
}

type Foo struct {
	FirstName string `tag_name:"tag 1"`
	LastName  string `tag_name:"tag 2"`
	Age       int    `tag_name:"tag 3"`
}

func (f *Foo) reflect() {
	val := reflect.ValueOf(f).Elem()

	fmt.Printf("valElem: %+v \n", val)

	// NumField returns a struct type's field count.
	// It panics if the type's Kind is not Struct.
	// NumField() int
	for i := 0; i < val.NumField(); i++ {

		// Field returns a struct type's i'th field.
		// It panics if the type's Kind is not Struct.
		// It panics if i is not in the range [0, NumField()).
		// Field(i int) StructFiel
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", typeField.Name, valueField.Interface(), tag.Get("tag_name"))
	}
}

func main() {
	MyType := &MyType{22, "wowzai"}

	mtV := reflect.ValueOf(&MyType).Elem()

	fmt.Println("Before:", mtV.MethodByName("String").Call(nil)[0])
	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(18)
	mtV.MethodByName("SetI").Call(params)
	params[0] = reflect.ValueOf("reflection test")
	mtV.MethodByName("SetName").Call(params)
	fmt.Println("After:", mtV.MethodByName("String").Call(nil)[0])

	// In returns the type of a function type's i'th input parameter.
	// It panics if the type's Kind is not Func.
	// It panics if i is not in the range [0, NumIn()).
	// In(i int) Type
	firstArg := reflect.TypeOf(SetMyType).In(0)
	fmt.Printf("firstArg: %+v \n", firstArg)

	// Elem returns a type's element type.
	// It panics if the type's Kind is not Array, Chan, Map, Ptr, or Slice.
	// Elem() Type
	firstArgElem := firstArg.Elem()
	fmt.Printf("firstArgElem: %+v \n", firstArgElem)

	// Name returns the type's name within its package.
	// It returns an empty string for unnamed types.
	// Name() string
	firstArgTypeStr := firstArgElem.Name()
	fmt.Printf("firstArgTypeStr %s \n", firstArgTypeStr)

	f := &Foo{
		FirstName: "Drew",
		LastName:  "Olson",
		Age:       30,
	}

	f.reflect()
}
