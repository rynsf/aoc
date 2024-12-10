package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var topoMap [][]int

type vec struct {
	x, y int
}

func outOfBound(x, y int) bool {
	if x < 0 || y < 0 || x >= len(topoMap[0]) || y >= len(topoMap) {
		return true
	}
	return false
}

func contain(pos vec, spots *[]vec) bool {
	for _, i := range *spots {
		if pos.x == i.x && pos.y == i.y {
			return true
		}
	}
	return false
}

func validAppend(pos vec, spots *[]vec) {
	if pos.x != -1 && pos.y != -1 {
		if !contain(pos, spots) {
			*spots = append(*spots, pos)
		}
	}
}

func score(x, y, last int, spots *[]vec) vec {
	if outOfBound(x, y) {
		return vec{-1, -1}
	}
	if topoMap[y][x]-1 != last {
		return vec{-1, -1}
	}
	if topoMap[y][x] == 9 {
		return vec{x, y}
	}
	right := score(x+1, y, topoMap[y][x], spots)
	left := score(x-1, y, topoMap[y][x], spots)
	top := score(x, y-1, topoMap[y][x], spots)
	bottom := score(x, y+1, topoMap[y][x], spots)
	validAppend(right, spots)
	validAppend(left, spots)
	validAppend(top, spots)
	validAppend(bottom, spots)
	return vec{-1, -1}
}

func main() {
	fmt.Println("Hello World")
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		s := sc.Text()
		l := make([]int, 0)
		for _, c := range s {
			n, _ := strconv.Atoi(string(c))
			l = append(l, n)
		}
		topoMap = append(topoMap, l)
	}

	result := 0
	for y, l := range topoMap {
		for x, spot := range l {
			if spot == 0 {
				spots := make([]vec, 0)
				score(x, y, -1, &spots)
				result += len(spots)
			}
		}
	}
	fmt.Println(result)
}
