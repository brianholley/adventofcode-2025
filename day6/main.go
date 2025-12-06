package main

import (
	"adventofcode2025/lib"
	"bufio"
	"fmt"
	"slices"
	"strings"
)

func main() {
	result := lib.Run(part1, part2)
	fmt.Println(result)
}

func part1(stdin *bufio.Scanner) string {
	result := 0

	arguments := [][]string{}

	for stdin.Scan() {
		line := stdin.Text()
		args := slices.DeleteFunc(strings.Split(line, " "), func(s string) bool { return s == "" })
		arguments = append(arguments, args)
	}

	for problem := range arguments[0] {
		operator := arguments[len(arguments)-1][problem]

		answer := lib.Atoi(arguments[0][problem])
		for line := 1; line < len(arguments)-1; line++ {
			if operator == "+" {
				answer += lib.Atoi(arguments[line][problem])
			} else if operator == "*" {
				answer *= lib.Atoi(arguments[line][problem])
			} else {
				fmt.Println("unknown operator", operator)
			}
		}
		result += answer
	}
	return fmt.Sprint(result)
}

func part2(stdin *bufio.Scanner) string {
	result := 0

	// Don't strip whitespace, we need it for alignment
	lines := []string{}
	for stdin.Scan() {
		line := stdin.Text()
		lines = append(lines, line)
	}

	// Find the expected length of each problem group from the separation between operators
	offsets := []int{}
	operators := lines[len(lines)-1]
	for i := 0; i < len(operators); i++ {
		if operators[i] != ' ' {
			offsets = append(offsets, i)
		}
	}
	offsets = append(offsets, len(operators)+1)

	// Split argument lines by problem length with whitespace preserved
	arguments := [][]string{}
	for _, line := range lines {
		args := []string{}
		for o := 1; o < len(offsets); o++ {
			args = append(args, line[offsets[o-1]:offsets[o]-1])
		}
		arguments = append(arguments, args)
	}

	for problem := range arguments[0] {
		operator := arguments[len(arguments)-1][problem][0]
		digits := offsets[problem+1] - offsets[problem] - 1

		// Read the columns of digits for the inputs
		inputs := []int{}
		for d := 0; d < digits; d++ {
			input := ""
			for line := 0; line < len(arguments)-1; line++ {
				input += string(arguments[line][problem][digits-d-1])
			}
			inputs = append(inputs, lib.Atoi(strings.Trim(input, " ")))
		}

		if operator == '+' {
			result += lib.ArraySum(inputs)
		} else if operator == '*' {
			result += lib.ArrayProduct(inputs)
		} else {
			fmt.Println("unknown operator", operator)
		}
	}
	return fmt.Sprint(result)
}
