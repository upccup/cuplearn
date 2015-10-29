package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")

	numPtr := flag.Int("num", 42, "an int")

	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word: ", *wordPtr)
	fmt.Println("numb: ", *numPtr)
	fmt.Println("fork: ", *boolPtr)
	fmt.Println("svar: ", *svar)
	fmt.Println("tail: ", flag.Args())
}

/*

$ ./flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

$ ./flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

$ ./flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]


$ ./flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
trailing: [a1 a2 a3 -numb=7]

$ ./flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

 $ ./flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...

*/
