package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := readInputData("./data.txt")

	solutionPart1 := part1(data)
	fmt.Println(solutionPart1)

	solutionPart2 := part2(data)
	fmt.Println(solutionPart2)
}

func part1(data [][]string) int {
	var countSafe int = 0
	for _, report := range data {
		if isSafeReport(report) {
			countSafe += 1
		}
	}
	return countSafe
}

func isSafeReport(data []string) bool {
	increasing := true
	decreasing := true

	for j := 0; j < len(data)-1; j++ {
		a, _ := strconv.Atoi(data[j])
		b, _ := strconv.Atoi(data[j+1])

		if a < b {
			decreasing = false
		} else if a > b {
			increasing = false
		} else {
			increasing = false
			decreasing = false
		}

		difference := int(math.Abs(float64(a) - float64(b)))
		if difference < 1 || difference > 3 {
			return false
		}
	}

	return increasing || decreasing
}

func part2(data [][]string) int {
	var countSafe int = 0
	for _, report := range data {
		if isSafeReportWithDampener(report) {
			countSafe++
		}
	}
	return countSafe
}

func isSafeReportWithDampener(data []string) bool {
	
	if isSafeReport(data) {
		return true
	}

	for i := 0; i < len(data); i++ {
		reducedReport := make([]string, 0, len(data)-1)
		reducedReport = append(reducedReport, data[:i]...)
		reducedReport = append(reducedReport, data[i+1:]...)

		if isSafeReport(reducedReport) {
			return true
		}
	}

	return false
}

func readInputData(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("file err ", err)
	}
	defer file.Close()

	bufReader := bufio.NewReader(file)

	a := [][]string{}
	for {
		line, _, err := bufReader.ReadLine()
		if len(line) > 0 {
			lineData := strings.Split(string(line), " ")
			a = append(a, lineData)
		}
		if err != nil {
			break
		}
	}

	return a
}
