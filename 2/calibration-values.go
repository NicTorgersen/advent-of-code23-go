package main

import (
	"aoc/2/file"
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"time"
	"unicode"
)

/**
 *   This program (day 2 of advent of code 2023) reads a file (calibration-values.txt by default), transpiles it and sums the first and last digit present in it.
 */
func main() {
	start := time.Now()

	var sum int
	var filepath string = "calibration-values.txt"

	if len(os.Args) > 1 {
		filepath = os.Args[1]
		fmt.Println(os.Args[1])
	}

	file := file.OpenOrCreateFile(filepath)

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		sum += addendFromCalibrationValue(line)
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()

	fmt.Printf("Sum %d\n", sum)
	fmt.Printf("Execution time: %s", time.Since(start))
}

func addendFromCalibrationValue(stringContainingNumbers string) int {
	var reversedStringContainingNumbers string = reverseString(stringContainingNumbers)

	fmt.Println(stringContainingNumbers)

	firstNumber, firstNumberAtIndex := findDigit(stringContainingNumbers)
	firstStringNum, firstStringNumAtIndex := findWordedNumber(0, stringContainingNumbers)

	fmt.Println(reversedStringContainingNumbers)

	lastNumber, lastNumberAtIndex := findDigit(reversedStringContainingNumbers)
	lastStringNum, lastStringNumAtIndex := findWordedNumber(0, reversedStringContainingNumbers)

	if (firstStringNumAtIndex <= firstNumberAtIndex) && firstStringNumAtIndex > -1 {
		firstNumber = firstStringNum
	}

	if (lastStringNumAtIndex <= lastNumberAtIndex) && lastStringNumAtIndex > -1 {
		lastNumber = lastStringNum
	}

	// fmt.Println("\t", firstNumberAtIndex, firstNumber)
	// fmt.Println("\t", firstStringNumAtIndex, firstStringNum)

	// fmt.Println("\t", lastNumberAtIndex, lastNumber)
	// fmt.Println("\t", lastStringNumAtIndex, lastStringNum)

	addend := firstNumber*10 + lastNumber

	fmt.Println("Addend: ", addend)

	return addend
}

func findDigit(input string) (int, int) {
	inputAsRunes := []rune(input)

	var foundDigit int
	var foundDigitAt int

	for index := 0; index < len(inputAsRunes); index++ {
		if unicode.IsDigit(inputAsRunes[index]) {
			foundDigit = int(inputAsRunes[index] - '0') // if the character in question is '1', it's underlying uint32 value is 49. '0' underlying value is 48, so when we '1' - '0' (49 - 48) we get 1
			foundDigitAt = index
			break
		}
	}

	return foundDigit, foundDigitAt
}

/**
 * Returns the "worded" int (e.g. "one" is 1)
 * as well as it's index position in the string
 * (the index of the number 1 in string "twoone" is 3, while the index of number 2 is 0)
 */
func findWordedNumber(startIndex int, text string) (int, int) {
	textToNum := map[string]int{
		"one":   1,
		"two":   2,
		"six":   6,
		"eno":   1,
		"owt":   2,
		"xis":   6,
		"four":  4,
		"five":  5,
		"nine":  9,
		"ruof":  4,
		"evif":  5,
		"enin":  9,
		"three": 3,
		"seven": 7,
		"eight": 8,
		"eerht": 3,
		"neves": 7,
		"thgie": 8,
	}

	var buffer string

	if startIndex >= len(text)-2 {
		return -1, -1
	}

	for index := startIndex; index < len(text); index++ {
		buffer += string(text[index])

		if len(buffer) > 5 {
			buffer = ""
			break
		}

		for phrase, num := range textToNum {
			if strings.EqualFold(phrase, buffer) {
				return num, startIndex
			}
		}
	}

	return findWordedNumber(startIndex+1, text)
}

func reverseString(text string) string {
	textAsRunes := []rune(text)
	slices.Reverse(textAsRunes)

	return string(textAsRunes)
}
