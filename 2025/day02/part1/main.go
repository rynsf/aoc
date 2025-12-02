package main

import (
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

func main() {
	bfile, err := os.ReadFile("./input")
	if err != nil {
		log.Fatal(err)
	}

	line := strings.TrimSpace(string(bfile))
	fmt.Println(line)
	strRanges := strings.Split(line, ",")
	fmt.Println(strRanges)

	ranges := [][]string{}

	for _, r := range strRanges {
		strNums := strings.Split(r, "-")
		ranges = append(ranges, strNums)
	}
	fmt.Println(ranges)

	sum := 0

	for _, r := range ranges {
		low := atoi(r[0])
		high := atoi(r[1])

		for i := low; i <= high; i++ {
			s := strconv.Itoa(i)

			if len(s)%2 == 0 {
				if s[:len(s)/2] == s[len(s)/2:] {
					sum += i
				}
			}
		}
	}

	fmt.Println(sum)
}
