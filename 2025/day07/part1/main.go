package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

var manifold [][]rune

func printManifold() {
	for _, l := range manifold {
		for _, c := range l {
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
}

func beamTravel(x, y int) int {
	if x < 0 || x >= len(manifold[0]) || y < 0 || y >= len(manifold) {
		return 0
	} else if manifold[y][x] == '|' {
		return 0
	} else if manifold[y][x] == '^' {
		return 1 + beamTravel(x-1, y) + beamTravel(x+1, y)
	} else {
		manifold[y][x] = '|'
		return beamTravel(x, y+1)
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
	fmt.Println(beamTravel(start, 1))
}
