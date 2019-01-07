package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//readFromStdin()
	readFromFile()
}

func readFromStdin() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
		//fmt.Printf("%s : %d",input.Text(),counts[input.Text()])
	}

	//ctrl+d Stop os.Stdin
	for line, _ := range counts {
		if counts[line] > 1 {
			fmt.Println(line)
		}
	}
}

func readFromFile() {
	counts := make(map[string]int)
	files := os.Args[1:]
	tempvalue := "start"
	if len(files) == 0 {
		countLines(os.Stdin, counts,tempvalue)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Printf("%v\n", err)
				continue
			}
			fmt.Println(f.Name())
			countLines(f, counts,tempvalue)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 0 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	fmt.Println(tempvalue)
}

func countLines(f *os.File, counts map[string]int, tempvalue string) {

	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
	}
	tempvalue += "a"
}
