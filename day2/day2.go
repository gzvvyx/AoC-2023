package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RuleSet struct {
	red int
	green int
	blue int
}

// var Rule = RuleSet {12, 13, 14}

func main() {
	file, err := os.Open("in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// possible := 0
	sumofpower := 0
	
	for scanner.Scan() {
		line := scanner.Text()
		lineSlice := strings.Split(line, ": ")
		ID := lineSlice[0]
		IDnum, _ := strconv.Atoi(strings.Split(ID, " ")[1])
		fmt.Println(IDnum)

		lineSlice = strings.Split(lineSlice[1], "; ")

		var Max = RuleSet {0, 0, 0}

		for _, turn := range lineSlice[0:] {
			colors := strings.Split(turn, ", ")

			for _, color := range colors {

				colorSlice := strings.Split(color, " ")
				number, _ := strconv.Atoi(colorSlice[0])

				// part1
				// if colorSlice[1] == "red" {
				// 	satisfy = Rule.red >= number
				// } else if colorSlice[1] == "blue" {
				// 	satisfy = Rule.blue >= number
				// } else if colorSlice[1] == "green" {
				// 	satisfy = Rule.green >= number
				// }

				// part2
				if colorSlice[1] == "red" {
					if number > Max.red { Max.red = number }
				} else if colorSlice[1] == "blue" {
					if number > Max.blue { Max.blue = number }
				} else if colorSlice[1] == "green" {
					if number > Max.green { Max.green = number }
				}

				// if !satisfy {break}
			}

			// if !satisfy {break}
		}
		// fmt.Println(satisfy)

		// if satisfy {
		// 	possible += IDnum
		// }
		sumofpower += Max.red * Max.blue * Max.green


	}

	// fmt.Println(possible)
	fmt.Println(sumofpower)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
	}
}
