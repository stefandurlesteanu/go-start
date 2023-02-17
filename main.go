package main

import (
	"fmt"
	"strconv"
)

var (
	actorName string = "Cool Actor"
	companion string = "Cool Companion"
	season    int    = 3
	episode   int    = 10
)

func main() {
	fmt.Println(actorName)

	//Conversion
	var i int = 42
	var j float32
	j = float32(i)
	fmt.Println(j)

	var x string
	x = strconv.Itoa(i)
	fmt.Println(x)
}
