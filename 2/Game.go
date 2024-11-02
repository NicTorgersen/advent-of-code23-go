package main

var TotalCubes map[string]int = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

type Set struct {
	cubes map[string]int
}

func (set Set) Power() uint {
	var coefficient uint = 1

	for _, count := range set.cubes {
		coefficient *= uint(count)
	}

	return coefficient
}

type Game struct {
	id   uint8
	sets []Set
}

func (game Game) IsPossible() bool {
	var possible bool = true

	for _, set := range game.sets {
		for color, max := range TotalCubes {
			if set.cubes[color] > max {
				return false
			}
		}
	}

	return possible
}

func (game Game) FewestNeeded() Set {
	var fewestNeededSet Set = Set{
		cubes: map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		},
	}

	for _, set := range game.sets {
		for color, countSeenInSet := range set.cubes {
			if countSeenInSet > fewestNeededSet.cubes[color] {
				fewestNeededSet.cubes[color] = countSeenInSet
			}
		}
	}

	return fewestNeededSet
}

func NewGame(id uint8, sets []Set) Game {
	return Game{
		id:   id,
		sets: sets,
	}
}
