package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var sup int
	var err error

	args := os.Args
	if len(args) == 2 {
		sup, err = strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
	}
	hello(sup)
}

func hello(yeah int) {
	for i := 0; i < yeah; i = i + 1 {
		fmt.Printf("hello!!!!!!!!! %v\n", i)
	}
}
