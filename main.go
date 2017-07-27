package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	fooCount, barCount := 0, 0
	for scanner.Scan() {
		word := scanner.Text()
		switch word {
		case "foo":
			fooCount++
		case "bar":
			barCount++
		}
	}
	fmt.Println("foo:", fooCount)
	fmt.Println("bar:", barCount)
}
