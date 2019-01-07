package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
		//fmt.Printf("%s : %d",input.Text(),counts[input.Text()])
	}

	//ctrl+d Stop os.Stdin
	for line,_:= range counts {
		if counts[line]>1 {
			fmt.Println(line)	
		}
	}

}
