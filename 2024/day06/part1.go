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

func outOfBound(pos vector, lenx, leny int) bool {
	if pos.x >= lenx || pos.y >= leny || pos.x < 0 || pos.y < 0 {
		return true
	}
	return false
}

func moveInDirec(pos vector, dir int) vector {
	directions := []vector{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}
	return vector{pos.x + directions[dir].x, pos.y + directions[dir].y}
}

func main() {
	f, err := os.Open("./input")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(f)
	area := []string{}
	for sc.Scan() {
		line := sc.Text()
		area = append(area, line)
	}

	pos := vector{0, 0}
	facingDir := 0
	areaMove := make([][]int, len(area))
	for i := range areaMove {
		areaMove[i] = make([]int, len(area[0]))
	}

	for y, s := range area {
		for x, c := range s {
			if c == '^' {
				pos.x = x
				pos.y = y
			}
		}
	}

	for {
		newPos := moveInDirec(pos, facingDir)
		if outOfBound(newPos, len(area[0]), len(area)) {
			areaMove[pos.y][pos.x] = 1
			break
		} else if area[newPos.y][newPos.x] == '#' {
			facingDir = (facingDir + 1) % 4
		} else {
			areaMove[pos.y][pos.x] = 1
			pos = newPos
		}
		//check the square inform of current possition.
		//if it is out of bound, then mark the current square and break
		//if the character in front is '.', move the pos in the current direction
		//if the chracter in front is '#', turn right.

	}

	count := 0
	for _, a := range areaMove {
		for _, c := range a {
			if c == 1 {
				count += 1
			}
		}
	}
	fmt.Println(count)
}
