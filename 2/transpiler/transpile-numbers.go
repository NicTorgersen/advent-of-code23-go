package transpiler

import (
	"aoc/2/file"
	"bufio"
	"strconv"
	"strings"
)

func Transpile(filepath string) string {
	var transpiledFilepath string = strings.Replace(filepath, ".txt", "-numberified.txt", 1)

	stringifiedNumberFile := file.OpenOrCreateFile(filepath)
	transpiledFile := file.OpenOrCreateFile(transpiledFilepath)

	scanner := bufio.NewScanner(stringifiedNumberFile)
	writer := bufio.NewWriter(transpiledFile)

	searches := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for scanner.Scan() {
		line := scanner.Text()

		numerified := translateStringifiedNumbers(searches, line)

		writer.WriteString(numerified + "\n")
		writer.Flush()
	}

	stringifiedNumberFile.Close()
	transpiledFile.Close()

	return transpiledFilepath
}

/**
 * Takes a string like "onetwothree" and translates it to "123"
 */
func translateStringifiedNumbers(searches []string, stringified string) string {
	var translatedString string = stringified

	for index := len(searches) - 1; index >= 0; index-- {
		translatedString = strings.Replace(
			translatedString,
			searches[index],
			strconv.Itoa(index+1),
			-1,
		)
	}

	return translatedString
}
