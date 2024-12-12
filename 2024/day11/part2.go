package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var memo = make(map[int][2]int)

func evolveStone(stone int) [2]int {
	m, ok := memo[stone]
	if ok {
		return m
	}
	var newStones [2]int
	if stone == 0 {
		newStones[0] = 1
		newStones[1] = -1
	} else if len(strconv.Itoa(stone))%2 == 0 {
		str := strconv.Itoa(stone)
		newStones[0], _ = strconv.Atoi(str[:len(str)/2])
		newStones[1], _ = strconv.Atoi(str[len(str)/2:])
	} else {
		newStones[0] = stone * 2024
		newStones[1] = -1
	}
	memo[stone] = newStones
	return newStones
}

func evolve(stones map[int]int) map[int]int {
	newMap := make(map[int]int)
	for s, n := range stones {
		e := evolveStone(s)
		for i := 0; i < 2; i++ {
			if e[i] != -1 {
				newMap[e[i]] += n
			}
		}
	}
	return newMap
}

func main() {
	f, err := os.Open("./input")
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

	stonesMap := make(map[int]int)
	for _, s := range stones {
		stonesMap[s] += 1
	}

	for i := 0; i < 75; i++ {
		stonesMap = evolve(stonesMap)
	}

	totalStones := 0
	for _, n := range stonesMap {
		totalStones += n
	}
	fmt.Println(totalStones)
}
