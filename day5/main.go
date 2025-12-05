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

func part1(stdin *bufio.Scanner) string {
	result := 0
	ranges := [][2]int{}
	for stdin.Scan() && stdin.Text() != "" {
		line := stdin.Text()
		parts := strings.Split(line, "-")
		ranges = append(ranges, [2]int{lib.Atoi(parts[0]), lib.Atoi(parts[1])})
	}

	ingredients := []int{}
	for stdin.Scan() {
		line := stdin.Text()
		ingredients = append(ingredients, lib.Atoi(line))
	}

	for _, ingredient := range ingredients {
		for _, r := range ranges {
			if ingredient >= r[0] && ingredient <= r[1] {
				// fmt.Println(ingredient)
				result++
				break
			}
		}
	}

	return fmt.Sprint(result)
}

func part2(stdin *bufio.Scanner) string {
	result := 0
	ranges := [][2]int{}
	for stdin.Scan() && stdin.Text() != "" {
		line := stdin.Text()
		parts := strings.Split(line, "-")
		ranges = append(ranges, [2]int{lib.Atoi(parts[0]), lib.Atoi(parts[1])})
	}

	for i := 0; i < len(ranges); i++ {
		for j := i + 1; j < len(ranges); {
			// Check for overlap
			start := lib.Max(ranges[i][0], ranges[j][0])
			end := lib.Min(ranges[i][1], ranges[j][1])
			if start <= end {
				// Ranges overlap - merge them and remove the second one
				min := lib.Min(ranges[i][0], ranges[j][0])
				max := lib.Max(ranges[i][1], ranges[j][1])
				ranges[i][0] = min
				ranges[i][1] = max
				ranges = append(ranges[:j], ranges[j+1:]...)
				// Reset the search, previous ranges may now overlap
				j = i + 1
				continue
			}
			j++
		}

		result += ranges[i][1] - ranges[i][0] + 1
	}

	return fmt.Sprint(result)
}
