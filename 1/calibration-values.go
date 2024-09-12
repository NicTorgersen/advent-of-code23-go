package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
	"unicode"
)

func main() {
	start := time.Now()

	var sum int
	var filepath string = "calibration-values.txt"

	if len(os.Args) > 1 {
		filepath = os.Args[1]
		fmt.Println(os.Args[1])
	}

	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		fmt.Println(line)

		sum += addendFromCalibrationValue(line)
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sum %d\n", sum)
	fmt.Printf("Execution time: %s", time.Since(start))
}

func addendFromCalibrationValue(stringContainingNumbers string) int {
	var stringAsRunes []rune = []rune(stringContainingNumbers) // convert the string to its underlying rune representation

	var firstNumber int
	var lastNumber int
	var addend int

	for letter := 0; letter < len(stringAsRunes); letter++ {
		if unicode.IsDigit(stringAsRunes[letter]) {
			firstNumber = int(stringAsRunes[letter] - '0') // if the character in question is '1', it's underlying uint32 value is 49. '0' underlying value is 48, so when we '1' - '0' (49 - 48) we get 1
			break
		}
	}

	// reverse the first loop to find the last number in the sequence
	for letter := len(stringAsRunes) - 1; letter >= 0; letter-- {
		if unicode.IsDigit(stringAsRunes[letter]) {
			lastNumber = int(stringAsRunes[letter] - '0')
			break
		}
	}

	addend = firstNumber*10 + lastNumber

	fmt.Println("First number: ", firstNumber)
	fmt.Println("Last number: ", lastNumber)
	fmt.Println("Addend: ", addend)

	return addend
}
