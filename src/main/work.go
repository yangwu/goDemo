package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("%b%v\n", i,":" + arg)
	}
	fmt.Println(os.Args[0])
}
