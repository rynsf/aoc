package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func evolve(stones []int) []int {
	newStones := make([]int, 0)
	for _, s := range stones {
		if s == 0 {
			newStones = append(newStones, 1)
		} else if len(strconv.Itoa(s))%2 == 0 {
			str := strconv.Itoa(s)
			n1, _ := strconv.Atoi(str[:len(str)/2])
			n2, _ := strconv.Atoi(str[len(str)/2:])
			newStones = append(newStones, n1, n2)
		} else {
			newStones = append(newStones, s*2024)
		}
	}
	return newStones
}

func main() {
	f, err := os.Open("./demoinput2")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	stones := make([]int, 0)
	for sc.Scan() {
		s := strings.Fields(sc.Text())
		for _, c := range s {
			n, _ := strconv.Atoi(c)
			stones = append(stones, n)
		}
	}

	fmt.Println(stones)
	for i := 0; i < 12; i++ {
		stones = evolve(stones)
		fmt.Println(stones)
	}
	fmt.Println(len(stones))
}
