package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func match (match string) string {
	switch match {
		case "one":
			return "1"
		case "two":
			return "2"
		case "three":
			return "3"
		case "four":
			return "4"
		case "five":
			return "5"
		case "six":
			return "6"
		case "seven":
			return "7"
		case "eight":
			return "8"
		case "nine":
			return "9"
		default:
			return match
	}
}

func getLastIndex(s string) int {
	index := -1
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for _, number := range numbers {
			if (strings.LastIndex(s, number) > index) {
				index = strings.LastIndex(s, number)
			}
	}
	return index
}

func replace(s string) string {
	re := regexp.MustCompile("one|two|three|four|five|six|seven|eight|nine")
	sep := 0
	firstHalf := ""
	secondHalf := ""

	firstIndex := re.FindStringIndex(s)
	if firstIndex != nil {
		firstIndexStart, firstIndexEnd := firstIndex[0], firstIndex[1]
		sep = firstIndexStart+1
		sAfterFirst := s[:firstIndexStart] + re.ReplaceAllStringFunc(s[firstIndexStart:firstIndexEnd], match) + s[firstIndexEnd:]
		firstHalf = s[:firstIndexStart] + sAfterFirst[firstIndexStart:sep]
	}

	secondHalf = s[sep:]

	x := getLastIndex(secondHalf)
	if (x != -1) {
		lastIndex := re.FindStringIndex(secondHalf[x:])
		lastIndexStart, lastIndexEnd := x, lastIndex[1] + x
		secondHalf = secondHalf[:lastIndexStart] + re.ReplaceAllStringFunc(secondHalf[lastIndexStart:lastIndexEnd], match) + secondHalf[lastIndexEnd:]
	}

	return firstHalf + secondHalf
}

func main() {
	// Open the file
	file, err := os.Open("in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read from the file
	scanner := bufio.NewScanner(file)

	sum := 0
	
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)

		char1 := ""
		char2 := ""

		converted := replace(line)

		// for part1 replace "converted" with "line" in loops
		for _, char := range converted {
			if unicode.IsDigit(char) {
				char1 = string(char)
				break
			}
		}

		for i := len(converted) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(converted[i])) {
				char2 = string(converted[i])
				break
			} else {continue}
		}

		number := 0
		number, _ = strconv.Atoi(char1 + char2)

		sum += number

	}

	fmt.Println(sum)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
	}
}
