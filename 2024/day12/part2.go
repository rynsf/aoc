package main

import (
	"bufio"
	"fmt"
	"os"
)

func outOfBound(x, y, w, h int) bool {
	if x < 0 || y < 0 || x >= w || y >= h {
		return true
	}
	return false
}

func b2i(a bool) int {
	if a {
		return 1
	}
	return 0
}

func bin2num(a, b, c, d bool) int {
	return (b2i(a) * 8) + (b2i(b) * 4) + (b2i(c) * 2) + b2i(d)
}

var numOfVertex = []int{4, 2, 2, 1, 2, 0, 1, 0, 2, 1, 0, 0, 1, 0, 0, 0}

func vertex(x, y int, farmMap [][]int) int {
	c := farmMap[y][x]

	sameElement := func(x, y int) bool {
		if outOfBound(x, y, len(farmMap[0]), len(farmMap)) {
			return false
		}
		if c != farmMap[y][x] {
			return false
		}
		return true
	}

	topx, topy := x, y-1
	rightx, righty := x+1, y
	bottomx, bottomy := x, y+1
	leftx, lefty := x-1, y

	blockTop := sameElement(topx, topy)
	blockRight := sameElement(rightx, righty)
	blockBottom := sameElement(bottomx, bottomy)
	blockLeft := sameElement(leftx, lefty)

	i := bin2num(blockTop, blockRight, blockBottom, blockLeft)
	num := numOfVertex[i]

	if i == 15 {
		for i := -1; i <= 1; i += 2 {
			for j := -1; j <= 1; j += 2 {
				if !sameElement(x+i, y+j) {
					num += 1
				}
			}
		}
		return num
	}

	if num == 1 {
		offx := b2i(blockRight)*1 + b2i(blockLeft)*-1
		offy := b2i(blockBottom)*1 + b2i(blockTop)*-1
		if !sameElement(x+offx, y+offy) {
			num += 1
		}
	} else if num == 0 {
		offx := b2i(blockRight)*1 + b2i(blockLeft)*-1
		offy := b2i(blockBottom)*1 + b2i(blockTop)*-1
		if offx == 0 && offy == 0 {
			return num
		} else if offx == 0 {
			for i := -1; i <= 1; i += 2 {
				if !sameElement(x+i, y+offy) {
					num += 1
				}
			}
		} else if offy == 0 {
			for i := -1; i <= 1; i += 2 {
				if !sameElement(x+offx, y+i) {
					num += 1
				}
			}
		}
	}
	return num
}

func findAreaLines(farmMap [][]int) map[int][2]int {
	areaPeriMap := make(map[int][2]int)
	for y, l := range farmMap {
		for x, c := range l {
			arr := areaPeriMap[c]
			arr[0] += 1
			arr[1] += vertex(x, y, farmMap)
			areaPeriMap[c] = arr
		}
	}
	return areaPeriMap
}

func watershed(x, y, id int, block rune, farmMap [][]rune, newMap [][]int) {
	if outOfBound(x, y, len(farmMap[0]), len(farmMap)) {
		return
	}
	if newMap[y][x] != 0 {
		return
	}
	if farmMap[y][x] != block {
		return
	}
	newMap[y][x] = id
	watershed(x, y-1, id, block, farmMap, newMap)
	watershed(x+1, y, id, block, farmMap, newMap)
	watershed(x, y+1, id, block, farmMap, newMap)
	watershed(x-1, y, id, block, farmMap, newMap)
}

func separateRegion(farmMap [][]rune) [][]int {
	newMap := make([][]int, len(farmMap))
	for i := range newMap {
		newMap[i] = make([]int, len(farmMap[0]))
	}
	id := 1
	for y, l := range farmMap {
		for x, c := range l {
			if newMap[y][x] == 0 {
				watershed(x, y, id, c, farmMap, newMap)
				id += 1
			}
		}
	}
	return newMap
}

func main() {
	f, err := os.Open("./input")
	if err != nil {
		f.Close()
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	var farmMap [][]rune
	for sc.Scan() {
		farmMap = append(farmMap, []rune(sc.Text()))
	}

	newMap := separateRegion(farmMap)

	areaLineMap := findAreaLines(newMap)
	totalCost := 0
	for _, ap := range areaLineMap {
		totalCost += ap[0] * ap[1]
	}
	fmt.Println(areaLineMap)
	fmt.Println(totalCost)
}
