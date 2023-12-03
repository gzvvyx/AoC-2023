package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
	// "strconv"
	// "strings"
)

type symbol_t struct {
	posx int
	posy int
	symbol string
}

type number_t struct {
	number int
	posxs int
	posxe int
	// posy int, map key
	used bool
}


var matrix = []string{}

func main() {
	file, err := os.Open("in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	count2 := 0

	var symbols []symbol_t
	linesMap := make(map[int][]number_t)
	
	i := 0
	for scanner.Scan() {

		var numbers []number_t

		line := scanner.Text()
		for j := 0; j < len(line); j++ {
			c := line[j]
			if c == '.' {
				continue
			}
			if !unicode.IsDigit(rune(c)) {
				// is symbol
				symbols = append(symbols, symbol_t{j, i, string(c)})
				continue
			}
			// is number
			x := j + 1
			for x <= len(line) {
				if x == len(line) || !unicode.IsDigit(rune(line[x])) {
					// end of number or line
					num, _ := strconv.Atoi(line[j:x])
					numbers = append(numbers, number_t{num, j, x - 1, false})
					j = x - 1
					break
				}
				x++
			}
		}
		// Increment line
		linesMap[i] = numbers
		i++
	}

	for _, symbol := range symbols {
		stars := 1
		starcount := 0
		// fmt.Println(symbol)
		var check []int
		if symbol.posy == 0 {
			check = []int{0, 1}
		} else  {
			check = []int{symbol.posy - 1, symbol.posy, symbol.posy + 1}
		}

		for _, y := range check {
			// fmt.Println(linesMap[y])

			for j := range linesMap[y] {
				number := &linesMap[y][j]

				if number.used {
					continue
				}

				if symbol.posx+1 >= number.posxs && symbol.posx-1 <= number.posxe {
					number.used = true
					// fmt.Println(number.number)
					count += number.number

					if symbol.symbol == "*" {
						starcount++
						stars *= number.number
					}

				}
			}
		}
		if starcount == 2 {
			count2 += stars
		}
	}

	fmt.Println(count)
	fmt.Println(count2)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
	}
}
