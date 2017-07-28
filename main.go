package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func newSplitScanner(r io.Reader, split bufio.SplitFunc) *bufio.Scanner {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	return scanner
}

func main() {
	scanner := newSplitScanner(os.Stdin, bufio.ScanWords)
	count := make(map[string]int)
	for scanner.Scan() {
		word := scanner.Text()
		count[word]++
	}
	for k, v := range count {
		fmt.Printf("%s: %d\n", k, v)
	}
}
