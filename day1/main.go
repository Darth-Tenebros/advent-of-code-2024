package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	left, right := readInputData("./data.txt")
	slices.Sort(left)
	slices.Sort(right)

	solutionPart1 := part1(left, right)
	fmt.Println(solutionPart1)

	solutionPart2 := part2(left, right)
	fmt.Println(solutionPart2)

}

func part1(left, right []int) int {
	var total int = 0

	for i := 0; i < len(left); i++ {
		total += int(math.Abs(float64(left[i]) - float64(right[i])))
	}

	return total
}

func part2(left, right []int) int {

	simScores := []int{}
	for i := 0; i < len(left); i++ {
		var appearances int = 0
		for j := 0; j < len(right); j++ {
			if left[i] == right[j] {
				appearances += 1
			}
		}
		simScores = append(simScores, left[i]*appearances)
	}

	var total int = 0
	for _, num := range simScores {
		total += num
	}

	return total
}

func readInputData(fileName string) ([]int, []int) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("file err ", err)
	}
	defer file.Close()

	bufReader := bufio.NewReader(file)

	a := []int{}
	b := []int{}
	for {
		line, _, err := bufReader.ReadLine()
		if len(line) > 0 {
			lineData := strings.Split(string(line), "   ")
			left, _ := strconv.Atoi(lineData[0])
			right, _ := strconv.Atoi(lineData[1])
			a = append(a, left)
			b = append(b, right)
		}
		if err != nil {
			break
		}
	}

	return a, b
}
