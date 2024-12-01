package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	list1 := []int{}
	list2 := []int{}
	pwd, _ := os.Getwd()

	file, err := os.Open(pwd + "\\day1\\input\\input.txt")
	if err != nil {
		fmt.Println("Error when opening file")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numstrings := strings.Split(line, "   ")
		num1, _ := strconv.Atoi(numstrings[0])
		num2, _ := strconv.Atoi(numstrings[1])
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})

	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})
	distance := 0.0
	for i := range len(list1) {
		// i fucking hate this language
		distance += math.Abs(float64(list1[i] - list2[i]))
	}

	//fmt.Println(list1, list2)
	fmt.Printf("%d", int(distance))
}
