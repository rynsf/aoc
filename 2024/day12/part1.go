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

func perimeter(x, y int, farmMap [][]int) int {
	topx, topy := x, y-1
	rightx, righty := x+1, y
	bottomx, bottomy := x, y+1
	leftx, lefty := x-1, y
	w := len(farmMap[0])
	h := len(farmMap)
	p := 0

	if outOfBound(topx, topy, w, h) || farmMap[topy][topx] != farmMap[y][x] {
		p += 1
	}
	if outOfBound(rightx, righty, w, h) || farmMap[righty][rightx] != farmMap[y][x] {
		p += 1
	}
	if outOfBound(bottomx, bottomy, w, h) || farmMap[bottomy][bottomx] != farmMap[y][x] {
		p += 1
	}
	if outOfBound(leftx, lefty, w, h) || farmMap[lefty][leftx] != farmMap[y][x] {
		p += 1
	}
	return p
}

func findAreaPeri(farmMap [][]int) map[int][2]int {
	areaPeriMap := make(map[int][2]int)
	for y, l := range farmMap {
		for x, c := range l {
			arr := areaPeriMap[c]
			arr[0] += 1
			arr[1] += perimeter(x, y, farmMap)
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
	f, err := os.Open("./demoinput2")
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
	areaPeriMap := findAreaPeri(newMap)
	totalCost := 0
	for _, ap := range areaPeriMap {
		totalCost += ap[0] * ap[1]
	}
	fmt.Println(totalCost)
}
