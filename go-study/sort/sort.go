package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

type ByTime []os.FileInfo

func (t ByTime) Len() int {
	return len(t)
}

func (t ByTime) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t ByTime) Less(i, j int) bool {
	return t[i].ModTime().Before(t[j].ModTime())
}

func main() {
	files, _ := ioutil.ReadDir("../")
	sort.Sort(ByTime(files))
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i].ModTime())
	}
}
