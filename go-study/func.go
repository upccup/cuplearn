package main

import "fmt"

type Presion struct {
	Name string
}

func (p *Presion) setName(name string) {
	p.Name = name
}

func (p Presion) setName2(name string) {
	p.Name = name
}

type Dict map[string]int

func (d Dict) set(key string, value int) {
	d[key] = value
}

func main() {
	p := Presion{"aaa"}
	fmt.Println("original name: ", p.Name)
	p.setName2("bbb")
	fmt.Println("after change name by value: ", p.Name)
	p.setName("ccc")
	fmt.Println("after change name by reference: ", p.Name)

	dict := Dict{"foo": 1}
	dict.set("foo", 2)
	fmt.Println(dict)
}

/** output
original name:  aaa
after change name by value:  aaa
after change name by reference:  ccc
map[foo:2]
**/
