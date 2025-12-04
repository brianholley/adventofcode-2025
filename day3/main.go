package main

import (
	"adventofcode2025/lib"
	"bufio"
	"fmt"
	"math"
)

func main() {
	result := lib.Run(part1, part2)
	fmt.Println(result)
}

func part1(stdin *bufio.Scanner) string {
	result := 0
	for stdin.Scan() {
		line := stdin.Text()
		bank := lib.ParseStringOfIntsDelimited(line, "")

		tens := 0
		tens_index := 0
		for i := 0; i < len(bank)-1; i++ {
			if bank[i] > tens {
				tens = bank[i]
				tens_index = i
			}
		}

		ones := 0
		for j := tens_index + 1; j < len(bank); j++ {
			if bank[j] > ones {
				ones = bank[j]
			}
		}

		battery := tens*10 + ones
		// fmt.Println(battery)
		result += battery
	}
	return fmt.Sprint(result)
}

func part2(stdin *bufio.Scanner) string {
	result := 0
	for stdin.Scan() {
		line := stdin.Text()
		bank := lib.ParseStringOfIntsDelimited(line, "")

		on := []int{}
		index := 0
		for len(on) < 12 {
			digit := bank[index]
			digit_index := index
			for i := index + 1; i < len(bank)-(11-len(on)); i++ {
				if bank[i] > digit {
					digit = bank[i]
					digit_index = i
				}
			}
			on = append(on, digit)
			index = digit_index + 1
		}

		battery := 0
		for i := 0; i < len(on); i++ {
			battery = battery + on[i]*int(math.Pow10(11-i))
		}
		// fmt.Println(battery)
		result += battery
	}
	return fmt.Sprint(result)
}
