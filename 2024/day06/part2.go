package main

import (
	"bufio"
	"fmt"
	"os"
)

type vector struct {
	x int
	y int
}

type visited struct {
	pos vector
	dir int
}

var directions = []vector{
	{0, -1}, // N
	{1, 0},  // E
	{0, 1},  // S
	{-1, 0}, // W
}

func outOfBound(pos vector, lenx, leny int) bool {
	if pos.x >= lenx || pos.y >= leny || pos.x < 0 || pos.y < 0 {
		return true
	}
	return false
}

func moveInDirec(pos vector, dir int) vector {
	return vector{pos.x + directions[dir].x, pos.y + directions[dir].y}
}

func OnSameSpot(pos vector, dir int, areaVisited []visited) bool {
	for _, v := range areaVisited {
		if v.pos.x == pos.x && v.pos.y == pos.y && v.dir == dir {
			return true
		}
	}
	return false
}

// check if area is looping
func isLooping(area [][]rune) bool {
	pos := vector{0, 0}
	facingDir := 0
	areaVisited := make([]visited, len(area)*len(area[0]))

	// find the starting postion
	for y, s := range area {
		for x, c := range s {
			if c == rune('^') {
				pos.x = x
				pos.y = y
			}
		}
	}

	//loop till out of bound
	// // search if you have been on the same spot, with same direction, if yes, return looping
	// // else move forward and apend current postitoin and direction to arevisited.
	for {
		newPos := moveInDirec(pos, facingDir)
		if outOfBound(newPos, len(area[0]), len(area)) {
			return false
		}

		if OnSameSpot(pos, facingDir, areaVisited) {
			return true
		} else if area[newPos.y][newPos.x] == rune('#') {
			areaVisited = append(areaVisited, visited{pos, facingDir})
			facingDir = (facingDir + 1) % 4
		} else {
			areaVisited = append(areaVisited, visited{pos, facingDir})
			pos = newPos
		}
	}
}

func main() {
	f, err := os.Open("./input")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(f)
	area := [][]rune{}
	for sc.Scan() {
		line := sc.Text()
		area = append(area, []rune(line))
	}

	// loop over entire input with nest loop.
	// check if the current character is empty, marked by '.', if it is place an obstacle on it, the character '#'.
	// // check if this new board is looping, if it is add 1 to count
	// // remove the obstactle the was added at the end of the loop.

	count := 0
	for y := 0; y < len(area); y++ {
		for x := 0; x < len(area[0]); x++ {
			if area[y][x] == rune('.') {
				area[y][x] = rune('#')
				if isLooping(area) {
					count += 1
				}
				area[y][x] = rune('.')
			}
		}
		fmt.Println("line done", y)
	}
	fmt.Println(count)
}
