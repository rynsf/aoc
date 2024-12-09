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

type file struct {
	index int
	size  int
	id    int
}

func main() {
	f, err := os.Open("./input")
	defer f.Close()
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

	fileId := 0
	memEndPtr := 0
	mem := make([]memBlock, memReq)
	var files []file
	var empty []file
	for i, n := range memStructure {
		if i%2 == 0 {
			if n > 0 {
				files = append(files, file{memEndPtr, n, fileId})
			}
			for a := 0; a < n; a++ {
				mem[memEndPtr] = memBlock{true, fileId}
				memEndPtr++
			}
			fileId++
		} else {
			if n > 0 {
				empty = append(empty, file{memEndPtr, n, 0})
			}
			for a := 0; a < n; a++ {
				mem[memEndPtr] = memBlock{false, 0}
				memEndPtr++
			}
		}
	}

	for f := len(files) - 1; f >= 0; f-- {
		for e := 0; e < len(empty); e++ {
			if files[f].size <= empty[e].size && files[f].index > empty[e].index {
				for i, j := files[f].index, empty[e].index; i < files[f].index+files[f].size; i, j = i+1, j+1 {
					mem[j] = mem[i]
					mem[i] = memBlock{false, 0}
				}
				empty[e] = file{empty[e].index + files[f].size, empty[e].size - files[f].size, 0}
				break
			}
		}
	}

	checksum := 0
	for i, n := range mem {
		if n.filled {
			checksum += i * n.id
		}
	}
	fmt.Println(checksum)
}
