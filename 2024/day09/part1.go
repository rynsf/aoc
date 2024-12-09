package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type memBlock struct {
	filled bool
	id     int
}

func findGap(start int, mem []memBlock) int {
	for i := start; i < len(mem); i++ {
		if !mem[i].filled {
			return i
		}
	}
	return -1
}

func findBlock(start int, mem []memBlock) int {
	for i := start; i >= 0; i-- {
		if mem[i].filled {
			return i
		}
	}
	return -1
}

func main() {
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(f)
	sc.Scan()
	line := sc.Text()

	memReq := 0
	var memStructure []int
	for _, c := range line {
		n, err := strconv.Atoi(string(c))
		if err != nil {
			panic(err)
		}
		memReq += n
		memStructure = append(memStructure, n)
	}

	mem := make([]memBlock, memReq)

	fileId := 0
	memEndPtr := 0
	for i, n := range memStructure {
		if i%2 == 0 {
			for a := 0; a < n; a++ {
				mem[memEndPtr] = memBlock{true, fileId}
				memEndPtr++
			}
			fileId++
		} else {
			for a := 0; a < n; a++ {
				mem[memEndPtr] = memBlock{false, 0}
				memEndPtr++
			}
		}
	}

	gap := findGap(0, mem)
	block := findBlock(len(mem)-1, mem)
	for gap < block {
		mem[gap] = mem[block]
		mem[block] = memBlock{false, 0}
		gap = findGap(gap, mem)
		block = findBlock(block, mem)
	}

	checksum := 0
	for i, n := range mem {
		if n.filled {
			checksum += i * n.id
		}
	}
	fmt.Println(checksum)
}
