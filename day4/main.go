package main

import (
	"adventofcode2025/lib"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	result := lib.Run(part1, part2)
	fmt.Println(result)
}

func adjacentCounts(grid [][]string, row int, col int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if row+i < 0 || row+i >= len(grid) {
				continue
			}
			if col+j < 0 || col+j >= len(grid[0]) {
				continue
			}
			if grid[row+i][col+j] == "@" {
				count++
			}
		}
	}
	return count
}

func part1(stdin *bufio.Scanner) string {
	result := 0
	grid := [][]string{}
	for stdin.Scan() {
		line := stdin.Text()
		grid = append(grid, strings.Split(line, ""))
	}

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == "@" {
				count := adjacentCounts(grid, r, c)
				if count < 4 {
					result++
				}
			}
		}
	}

	return fmt.Sprint(result)
}

func part2(stdin *bufio.Scanner) string {
	result := 0
	grid := [][]string{}
	for stdin.Scan() {
		line := stdin.Text()
		grid = append(grid, strings.Split(line, ""))
	}

	removed := true
	for removed {
		removed = false

		for r := 0; r < len(grid); r++ {
			for c := 0; c < len(grid[0]); c++ {
				if grid[r][c] == "@" {
					count := adjacentCounts(grid, r, c)
					if count < 4 {
						result++
						removed = true
						grid[r][c] = "."
					}
				}
			}
		}
	}

	return fmt.Sprint(result)
}
