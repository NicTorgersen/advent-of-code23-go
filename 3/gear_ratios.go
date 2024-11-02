package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type Number struct {
	value int
	x     int
	y     int
	width int
}

func (number Number) String() string {
	return fmt.Sprintf("num %d - x: %d (end: %d), y: %d, w: %d", number.value, number.x, number.x+number.width, number.y, number.width)
}

type Symbol struct {
	value string
	x     int
	y     int
}

func (symbol Symbol) HasCollidedWith(number Number) bool {
	return (symbol.y-1 == number.y || symbol.y == number.y || symbol.y+1 == number.y) && // under / over / same line constraint
		(symbol.x-1 <= number.x+number.width && symbol.x+1 >= number.x ||
			symbol.x-1 <= number.x && symbol.x+1 >= number.x+number.width)
}

func (symbol Symbol) String() string {
	return fmt.Sprintf("sym %s - x: %d, y: %d", symbol.value, symbol.x, symbol.y)
}

func main() {
	if len(os.Args) < 2 {
		panic("Needs file as input")
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		panic(err)
	}

	var numbers []Number
	var symbols []Symbol

	scanner := bufio.NewScanner(file)

	for y := 0; scanner.Scan(); y++ {
		numbersInLine, symbolsInLine := ParseLine(scanner.Text(), y)

		numbers = append(numbers, numbersInLine...)
		symbols = append(symbols, symbolsInLine...)
	}

	var relevantNumbers = map[int]Number{}

	for _, number := range numbers {
		log.Println(number)
	}

	for _, symbol := range symbols {
		log.Println(symbol)
		for _, number := range numbers {
			if symbol.HasCollidedWith(number) {
				relevantNumbers[number.value] = number
			}

		}
	}

	for _, number := range relevantNumbers {
		log.Println(number.value)
	}
}

func ParseLine(line string, y int) ([]Number, []Symbol) {
	var buffer string
	var numbers []Number
	var symbols []Symbol
	var lineAsRunes = []rune(line)

	for i := 0; i < len(lineAsRunes); i++ {
		if unicode.IsDigit(lineAsRunes[i]) {
			// increase buffer because it's a number
			buffer += string(lineAsRunes[i])
			continue
		}

		if len(buffer) > 0 {
			number, err := strconv.Atoi(buffer)

			if err != nil {
				panic(err)
			}

			numbers = append(numbers, Number{
				value: number,
				y:     y,
				x:     i - len(buffer),
				width: len(buffer) - 1,
			})

			buffer = ""
		}

		if lineAsRunes[i] != 46 {
			symbols = append(symbols, Symbol{
				value: string(lineAsRunes[i]),
				x:     i,
				y:     y,
			})

			continue
		}
	}

	return numbers, symbols
}
