package main

import (
	"bufio"
	"fmt"
	"os"
)

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



		fmt.Println(line)
	}

	fmt.Println(sum)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
	}
}