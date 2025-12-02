package main

import (
	"adventofcode2025/lib"
	"bufio"
	"fmt"
	"math"
	"strings"
)

func main() {
	result := lib.Run(part1, part2)
	fmt.Println(result)
}

func parseRanges(stdin *bufio.Scanner) [][2]int {
	ranges := [][2]int{}

	for stdin.Scan() {
		line := stdin.Text()
		r := strings.Split(line, ",")
		for _, v := range r {
			bounds := strings.Split(v, "-")
			ranges = append(ranges, [2]int{
				lib.Atoi(bounds[0]),
				lib.Atoi(bounds[1]),
			})
		}
	}
	return ranges
}

func isRepeatedSequences(v int, seqCount int) bool {
	digits := int(math.Log10(float64(v))) + 1
	valid := digits%seqCount == 0
	for d := 0; d < digits/seqCount && valid; d++ {
		check := []int{}
		for s := 0; s < seqCount && valid; s++ {
			val := (v / int(math.Pow10(d+s*digits/seqCount))) % 10
			check = append(check, val)
			if val != check[0] {
				valid = false
			}
		}
	}
	return valid
}

func part1(stdin *bufio.Scanner) string {
	result := 0
	ranges := parseRanges(stdin)

	for i := range ranges {
		for v := ranges[i][0]; v <= ranges[i][1]; v++ {
			valid := isRepeatedSequences(v, 2)
			if valid {
				// fmt.Println(v)
				result += v
			}
		}
	}

	return fmt.Sprint(result)
}

func part2(stdin *bufio.Scanner) string {
	result := 0
	ranges := parseRanges(stdin)

	for i := range ranges {
		for v := ranges[i][0]; v <= ranges[i][1]; v++ {
			digits := int(math.Log10(float64(v))) + 1
			for seqCount := 2; seqCount <= digits; seqCount++ {
				if isRepeatedSequences(v, seqCount) {
					// fmt.Println(v)
					result += v
					break
				}
			}
		}
	}

	return fmt.Sprint(result)
}
