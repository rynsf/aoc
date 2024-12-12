package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var rules = make(map[int][]int)

func contains(arr1 []int, arr2 []int) bool {
	for _, i := range arr1 {
		for _, j := range arr2 {
			if i == j {
				return true
			}
		}
	}
	return false
}

func main() {
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	isRule := true
	var pages [][]int
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			isRule = false
			continue
		}
		if isRule {
			s := strings.Split(line, "|")
			n1, _ := strconv.Atoi(s[0])
			n2, _ := strconv.Atoi(s[1])
			rules[n1] = append(rules[n1], n2)
		} else {
			seprated := strings.Split(line, ",")
			var i []int
			for _, s := range seprated {
				n, _ := strconv.Atoi(s)
				i = append(i, n)
			}
			pages = append(pages, i)
		}
	}

	var unsorted [][]int
	for _, p := range pages {
		valid := true
		for i := 1; i < len(p); i++ {
			if contains(rules[p[i]], p[:i]) {
				valid = false
				break
			}
		}
		if !valid {
			unsorted = append(unsorted, p)
		}
	}

	for a := 0; a < len(unsorted); a++ {
		sort.Slice(unsorted[a], func(i, j int) bool {
			return !contains([]int{unsorted[a][i]}, rules[unsorted[a][j]])
		})
	}

	result := 0
	for _, arr := range unsorted {
		result += arr[len(arr)/2]
	}
	fmt.Println(result)
}