package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	return n
}

func parseFresh(sc *bufio.Scanner) [][]int {
	freshItems := [][]int{}

	for sc.Scan() {
		line := sc.Text()

		if line == "" {
			return freshItems
		}

		ranges := strings.Split(line, "-")
		low := atoi(ranges[0])
		high := atoi(ranges[1])

		freshItems = append(freshItems, []int{low, high})
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	return [][]int{}
}

func InRangeOf(i int, selfIndex int, freshItems [][]int) []int {
	rof := []int{}
	for n, f := range freshItems {
		if n != selfIndex {
			if i >= f[0] && i <= f[1] {
				rof = append(rof, n)
			}
		}
	}
	return rof
}

func inBoth(x, y []int) int {
	for _, i := range x {
		for _, y := range y {
			if i == y {
				return i
			}
		}
	}
	return -1
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(file)

	freshItems := parseFresh(sc)

	count := 0

	fmt.Println(freshItems)

	for i := 0; i < len(freshItems); i++ {
		indexLow := InRangeOf(freshItems[i][0], i, freshItems)
		indexHigh := InRangeOf(freshItems[i][1], i, freshItems)

		if len(indexLow) != 0 && len(indexHigh) != 0 && inBoth(indexLow, indexHigh) != -1 {
			freshItems[i][0] = -1
			freshItems[i][1] = -1
		} else {
			if len(indexLow) != 0 {
				freshItems[i][0] = freshItems[indexLow[0]][1] + 1
			}
			if len(indexHigh) != 0 {
				freshItems[i][1] = freshItems[indexHigh[0]][0] - 1
			}
		}
	}

	for _, f := range freshItems {
		if f[0] != -1 {
			c := f[1] - f[0] + 1
			if c > 0 {
				count += c
			}
		}
	}
	fmt.Println(count)
}
