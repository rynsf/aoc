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

func parseItemsList(sc *bufio.Scanner) []int {
	items := []int{}

	for sc.Scan() {
		line := sc.Text()
		items = append(items, atoi(line))
	}
	if err := sc.Err(); err != nil {
		log.Fatal()
	}
	return items
}

func isInRange(i int, freshItems [][]int) bool {
	for _, f := range freshItems {
		if i >= f[0] && i <= f[1] {
			fmt.Println(i, f)
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(file)

	freshItems := parseFresh(sc)
	itemsList := parseItemsList(sc)

	fmt.Println(freshItems)
	fmt.Println(itemsList)

	count := 0

	for _, i := range itemsList {
		if isInRange(i, freshItems) {
			count += 1
		}
	}

	fmt.Println(count)
}
