package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data := readInputFile("./data.txt")

	solutionPart1 := part1(data)
	fmt.Println(solutionPart1)

	solutionPart2 := part2(data)
	fmt.Println(solutionPart2)
}

func part1(input string) int {
	results := extractMulInstructions(input)

	total := 0
	for _, result := range results {
		total += result
	}

	return total
}

func extractMulInstructions(input string) []int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	var results []int
	for _, match := range matches {
		if len(match) == 3 {
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			results = append(results, x*y)
		}
	}

	return results
}

func part2(input string) int {
	results := extractMulInstructionsWithConditions(input)

	total := 0
	for _, result := range results {
		total += result
	}

	return total
}

func extractMulInstructionsWithConditions(input string) []int {
	mulRe := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	doRe := regexp.MustCompile(`do\(\)`)
	dontRe := regexp.MustCompile(`don't\(\)`)

	var results []int
	isEnabled := true

	doMatches := doRe.FindAllStringIndex(input, -1)
	dontMatches := dontRe.FindAllStringIndex(input, -1)

	// Combine and sort all instruction positions
	instructions := [][]int{}
	for _, match := range doMatches {
		instructions = append(instructions, append(match, 0)) // 0 for do
	}
	for _, match := range dontMatches {
		instructions = append(instructions, append(match, 1)) // 1 for don't
	}
	for _, match := range mulRe.FindAllStringSubmatchIndex(input, -1) {
		instructions = append(instructions, append(match[2:], 2)) // 2 for mul
	}

	// Sort instructions by their position in the string
	sortInstructions(instructions)

	for _, inst := range instructions {
		switch inst[len(inst)-1] {
		case 0: // do()
			isEnabled = true
		case 1: // don't()
			isEnabled = false
		case 2: // mul()
			if isEnabled {
				x, _ := strconv.Atoi(input[inst[0]:inst[1]])
				y, _ := strconv.Atoi(input[inst[2]:inst[3]])
				results = append(results, x*y)
			}
		}
	}

	return results
}

func sortInstructions(instructions [][]int) {
	for i := 0; i < len(instructions); i++ {
		for j := i + 1; j < len(instructions); j++ {
			if instructions[i][0] > instructions[j][0] {
				instructions[i], instructions[j] = instructions[j], instructions[i]
			}
		}
	}
}

func readInputFile(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return strings.Join(lines, "")
}
