package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Needs input file as argument.")
		return
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	games := ReadGames(file)
	var possibleGamesSum uint
	var sumOfPowers uint

	for _, game := range games {
		if game.IsPossible() {
			possibleGamesSum += uint(game.id)
		}

		sumOfPowers += game.FewestNeeded().Power()
	}

	fmt.Println(strings.Repeat("= ", 16))
	fmt.Printf("Sum of possible games: %d\n", possibleGamesSum)
	fmt.Printf("Sum of powers for fewest needed cubes: %d\n", sumOfPowers)
	fmt.Println(strings.Repeat("= ", 16))

	file.Close()
}

func ReadGames(file *os.File) []Game {
	var games []Game

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		gameRow := scanner.Text()

		gameId, remainder := ParseGameID(gameRow)
		sets := ParseSets(remainder)

		games = append(games, NewGame(gameId, sets))
	}

	return games
}

func ParseSets(gameRow string) []Set {
	var sets []Set = []Set{}

	s := strings.Split(gameRow, ";")

	for _, set := range s {
		setOfCubes := strings.Split(strings.Trim(set, " "), ",")

		var cubes map[string]int = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, cube := range setOfCubes {
			cc := strings.Split(strings.Trim(cube, " "), " ")

			count, err := strconv.Atoi(cc[0])

			if err != nil {
				panic(err)
			}

			color := cc[1]

			cubes[color] += count
		}

		sets = append(sets, Set{cubes})
	}

	return sets
}

func ParseGameID(gameRow string) (uint8, string) {
	after, found := strings.CutPrefix(gameRow, "Game ")

	if !found {
		panic("Could not find prefix Game in gameRow")
	}

	before, remainder, ok := strings.Cut(after, ":")

	if !ok {
		panic("Could not find : in gameRow")
	}

	id64, err := strconv.ParseUint(before, 10, 8)

	if err != nil {
		panic(err)
	}

	return uint8(id64), remainder
}
