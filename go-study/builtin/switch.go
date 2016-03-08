package main

import (
	"log"
)

func main() {
	integer := 6
	switch integer {
	case 4:
		log.Println("The integer was <= 4")
		fallthrough
	case 5:
		log.Println("The integer was <= 5")
		fallthrough
	case 6:
		log.Println("The integer was <= 6")
		fallthrough
	case 7:
		log.Println("The integer was <= 7")
		fallthrough
	case 8:
		log.Println("The integer was <= 8")
		fallthrough
	default:
		log.Println("default case")
	}

	switch integer {
	case 4:
		log.Println("The integer was <= 4")
		fallthrough
	case 5:
		log.Println("The integer was <= 5")
		fallthrough
	case 6:
		log.Println("The integer was <= 6")
		fallthrough
	case 7:
		log.Println("The integer was <= 7")
	case 8:
		log.Println("The integer was <= 8")
	default:
		log.Println("default case")
	}
}

// 2016/03/08 21:55:54 The integer was <= 6
// 2016/03/08 21:55:54 The integer was <= 7
// 2016/03/08 21:55:54 The integer was <= 8
// 2016/03/08 21:55:54 default case
// 2016/03/08 21:55:54 The integer was <= 6
// 2016/03/08 21:55:54 The integer was <= 7
