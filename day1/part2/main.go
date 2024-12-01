package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	list1 := []int{}
	occurence := make(map[int]int, 100)
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

		_, ok := occurence[num2]
		if !ok {
			occurence[num2] = 0
		}
		occurence[num2] += 1
	}

	similarity := 0
	for i := range len(list1) {
		// i fucking hate this language
		similarity += occurence[list1[i]] * list1[i]
	}

	fmt.Printf("%d", similarity)

}
