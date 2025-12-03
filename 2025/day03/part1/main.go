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

	fmt.Println(banks)

	sum := 0

	for _, b := range banks {
		d1, i1 := max(b[:len(b)-1])
		d2, _ := max(b[i1+1:])

		sum += d1*10 + d2
	}

	fmt.Println(sum)
}
