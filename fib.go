package main

import (
	"fmt"
	"os"
	"strconv"
)

var f int
var x int

func main() {
	fmt.Println(fib(getArg()))
}

func getArg() int {
	s := ""
	for _, arg := range os.Args[1:] {
		s += arg
	}
	f, _ := strconv.Atoi(s)
	return f
}

func fib(f int) int {
	x, y := 0, 1
	for i := 0; i < f; i++ {
		x, y = y, x+y
	}
	return x
}
