package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	data := readInputData("./data.txt")

	solutionPart1 := countXMAS(data)
	fmt.Println(solutionPart1)
}

func countXMAS(grid []string) int {
	xmasCount := 0
	height := len(grid)
	width := len(grid[0])

	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			// Horizontal
			if y+3 < width &&
				grid[x][y] == 'X' &&
				grid[x][y+1] == 'M' &&
				grid[x][y+2] == 'A' &&
				grid[x][y+3] == 'S' {
				xmasCount++
			}

			// Vertical
			if x+3 < height &&
				grid[x][y] == 'X' &&
				grid[x+1][y] == 'M' &&
				grid[x+2][y] == 'A' &&
				grid[x+3][y] == 'S' {
				xmasCount++
			}

			// Diagonal down-right
			if x+3 < height && y+3 < width &&
				grid[x][y] == 'X' &&
				grid[x+1][y+1] == 'M' &&
				grid[x+2][y+2] == 'A' &&
				grid[x+3][y+3] == 'S' {
				xmasCount++
			}

			// Diagonal down-left
			if x+3 < height && y-3 >= 0 &&
				grid[x][y] == 'X' &&
				grid[x+1][y-1] == 'M' &&
				grid[x+2][y-2] == 'A' &&
				grid[x+3][y-3] == 'S' {
				xmasCount++
			}

			// Reverse Horizontal
			if y+3 < width &&
				grid[x][y] == 'S' &&
				grid[x][y+1] == 'A' &&
				grid[x][y+2] == 'M' &&
				grid[x][y+3] == 'X' {
				xmasCount++
			}

			// Reverse Vertical
			if x+3 < height &&
				grid[x][y] == 'S' &&
				grid[x+1][y] == 'A' &&
				grid[x+2][y] == 'M' &&
				grid[x+3][y] == 'X' {
				xmasCount++
			}

			// Reverse Diagonal down-right
			if x+3 < height && y+3 < width &&
				grid[x][y] == 'S' &&
				grid[x+1][y+1] == 'A' &&
				grid[x+2][y+2] == 'M' &&
				grid[x+3][y+3] == 'X' {
				xmasCount++
			}

			// Reverse Diagonal down-left
			if x+3 < height && y-3 >= 0 &&
				grid[x][y] == 'S' &&
				grid[x+1][y-1] == 'A' &&
				grid[x+2][y-2] == 'M' &&
				grid[x+3][y-3] == 'X' {
				xmasCount++
			}
		}
	}

	return xmasCount
}

func readInputData(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("file err ", err)
	}
	defer file.Close()

	bufReader := bufio.NewReader(file)
	var data []string

	for {
		line, _, err := bufReader.ReadLine()
		if len(line) > 0 {
			data = append(data, string(line))
		}
		if err != nil {
			break
		}
	}
	return data
}
