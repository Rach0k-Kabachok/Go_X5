package main

import (
	"fmt"
	"os"
)

func ParseInt64(to_parse string) int64 {
	var res int64
	for _, elem := range to_parse {
		res *= 10
		res += int64(elem - '0')
	}

	return res
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	order := ParseInt64(os.Args[1])
	if order < 0 {
		fmt.Println("Ploho")
	} else if order < 10000 {
		fmt.Println("small")
	} else if order < 50000 {
		fmt.Println("regular")
	} else {
		fmt.Println("large")
	}
}
