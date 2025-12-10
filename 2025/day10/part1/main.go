package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	return n
}

func sameState(starting, final []bool) bool {
	for i, b := range starting {
		if b != final[i] {
			return false
		}
	}
	return true
}

type Queue [][]bool

// don't modify the original, return a new array
func applyButton(lstate []bool, b []int) []bool {
	dst := make([]bool, len(lstate))
	copy(dst, lstate)

	for _, i := range b {
		dst[i] = !dst[i]
	}

	return dst
}

// takes current state and final state of lights and buttons, returns the least number of steps needed.
func bfs(starting, final []bool, buttons [][]int) int {
	queue := Queue{starting}

	level := 0
	for len(queue) != 0 {

		levelSize := len(queue)
		for range levelSize {
			lstate := queue[0]
			queue = queue[1:]

			if sameState(lstate, final) {
				return level
			}

			for _, b := range buttons {
				nextState := applyButton(lstate, b)
				queue = append(queue, nextState)
			}
		}
		level += 1
	}

	log.Fatal("can't find a solution")
	return -1
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(file)

	indicatorLights := [][]bool{}
	buttonWirings := [][][]int{}

	for sc.Scan() {
		line := sc.Text()

		fields := strings.FieldsSeq(line)
		bWiring := [][]int{}

		for f := range fields {
			switch f[0] {
			case '[':
				lightsStr := f[1 : len(f)-1]
				lightIndicator := []bool{}
				for _, b := range lightsStr {
					if b == '.' {
						lightIndicator = append(lightIndicator, false)
					} else {
						lightIndicator = append(lightIndicator, true)
					}
				}
				indicatorLights = append(indicatorLights, lightIndicator)
			case '(':
				bw := []int{}
				buttonStr := strings.Split(f[1:len(f)-1], ",")
				for _, b := range buttonStr {
					bw = append(bw, atoi(string(b)))
				}
				bWiring = append(bWiring, bw)
			}
		}
		buttonWirings = append(buttonWirings, bWiring)
	}

	totalSteps := 0

	for i, buttons := range buttonWirings {
		l := indicatorLights[i]

		startingState := make([]bool, len(l))
		totalSteps += bfs(startingState, l, buttons)
	}

	fmt.Println(totalSteps)
}
