package main

import (
	"adventofcode2025/lib"
	"bufio"
	"fmt"
	"strconv"
)

func main() {
	result := lib.Run(part1, part2)
	fmt.Println(result)
}

func part1(stdin *bufio.Scanner) string {
	result := 0

	dial := 50

	for stdin.Scan() {
		line := stdin.Text()

		turn, _ := strconv.Atoi(line[1:])
		if line[0] == 'L' {
			dial = ((dial - turn) + 100) % 100
		} else {
			dial = (dial + turn) % 100
		}

		if dial == 0 {
			result++
		}
	}

	return fmt.Sprint(result)
}

func part2(stdin *bufio.Scanner) string {
	result := 0

	dial := 50

	for stdin.Scan() {
		line := stdin.Text()

		// fmt.Print(line, ": ", dial, " -> ")
		turn, _ := strconv.Atoi(line[1:])

		result = result + int(turn/100)
		turn = turn % 100

		if line[0] == 'L' {
			if dial != 0 && dial <= turn {
				result++
			}
			dial = ((dial - turn) + 100) % 100
		} else {
			if dial+turn >= 100 {
				result++
			}
			dial = (dial + turn) % 100
		}
		// fmt.Print(dial)
		fmt.Println()
	}

	return fmt.Sprint(result)
}
