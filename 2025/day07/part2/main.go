package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

var manifold [][]rune
var memo = make(map[int]int)

func beamTravel(x, y int) int {
	if m, ok := memo[y*len(manifold)+x]; ok {
		return m
	}
	if x < 0 || x >= len(manifold[0]) || y < 0 || y >= len(manifold) {
		return 0
	} else if manifold[y][x] == '^' {
		memo[y*len(manifold)+x] = 1 + beamTravel(x-1, y) + beamTravel(x+1, y)
		return memo[y*len(manifold)+x]
	} else {
		memo[y*len(manifold)+x] = beamTravel(x, y+1)
		return memo[y*len(manifold)+x]
	}
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal()
	}

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		line := sc.Text()
		manifold = append(manifold, []rune(line))
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	start := slices.Index(manifold[0], 'S')
	fmt.Println(beamTravel(start, 1) + 1)
}
