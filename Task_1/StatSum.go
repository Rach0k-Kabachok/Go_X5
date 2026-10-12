package main

import (
	"fmt"
	"math"
)

func PrintResult(minim, maxim, sum int64, avg float64) {
	fmt.Printf("%s%d\n%s%d\n%s%d\n%s%f\n",
		"Min:", minim,
		"Max: ", maxim,
		"Sum: ", sum,
		"Avg: ", avg)
}

func ParseInt64(to_parse string) int64 {
	var res int64
	for _, elem := range to_parse {
		res *= 10
		res += int64(elem - '0')
	}

	return res
}

func main() {
	var minim int64 = math.MaxInt64
	var maxim int64 = math.MinInt64
	var sum int64 = 0
	var count int64 = 0

	for {
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			break
		}

		cur_num := ParseInt64(input)
		minim = min(minim, cur_num)
		maxim = max(maxim, cur_num)
		sum += cur_num
		count++
	}

	if count == 0 {
		PrintResult(0, 0, 0, 0)
		return
	}

	PrintResult(minim, maxim, sum, float64(sum)/float64(count))
}
