package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	return n
}

// returns maximum and its index
func max(arr []int) (int, int) {
	maximum, index := 0, 0
	for i, n := range arr {
		if n > maximum {
			index = i
			maximum = n
		}
	}
	return maximum, index
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(file)

	var banks [][]int

	for sc.Scan() {
		line := sc.Text()

		b := []int{}
		for _, n := range line {
			b = append(b, atoi(string(n)))
		}

		banks = append(banks, b)
	}

	sum := 0

	for _, b := range banks {
		digit := 0
		start := 0
		for i := 11; i >= 0; i-- {
			digitTail, dIndex := max(b[start : len(b)-i])
			start += dIndex + 1
			digit = digit*10 + digitTail
		}
		sum += digit
	}

	fmt.Println(sum)
}
