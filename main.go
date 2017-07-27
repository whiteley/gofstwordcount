package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	count := make(map[string]int)
	for scanner.Scan() {
		word := scanner.Text()
		count[word]++
	}
	for k, v := range count {
		fmt.Printf("%s: %d\n", k, v)
	}
}
