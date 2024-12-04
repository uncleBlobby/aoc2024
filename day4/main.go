package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"slices"
	"strings"

	daily "github.com/uncleBlobby/aoc2024/internal"
)

// NOTE: in 2D array syntax, puzzle locations are written:
// [Y][X]PuzzleLocation and puzzle[Y][X]

type PuzzleLocation struct {
	X int
	Y int
}

type Direction struct {
	X int
	Y int
}

type Neighbour struct {
	Char      string
	Direction Direction
	Location  PuzzleLocation
}

func FindAllLocations(puzzle [][]string, targetChar string) []PuzzleLocation {
	var locs []PuzzleLocation

	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[i]); j++ {
			if puzzle[i][j] == targetChar {
				newLoc := PuzzleLocation{X: j, Y: i}
				locs = append(locs, newLoc)
			}
		}
	}

	return locs
}

func main() {
	fmt.Println("hello, aoc2024 day4!")

	input := daily.GetInputFromFile("input")

	scanner := bufio.NewScanner(input)

	lineCounter := 0

	puzzle := [][]string{}

	for scanner.Scan() {

		line := scanner.Text()

		chars := strings.Split(line, "")

		puzzle = append(puzzle, make([]string, len(chars)))
		puzzle[lineCounter] = chars

		lineCounter += 1
	}

	PartOne(puzzle)
	PartTwo(puzzle)

}

func PartTwo(puzzle [][]string) {

	allA := FindAllLocations(puzzle, "A")

	counter := 0
	for _, a := range allA {
		diagnbs := FindDiagonalNeighboursOfLocation(puzzle, a)

		diagChars := []string{}

		Mlocs := []PuzzleLocation{}
		Slocs := []PuzzleLocation{}

		for _, nb := range diagnbs {
			diagChars = append(diagChars, nb.Char)
			if nb.Char == "M" {
				Mlocs = append(Mlocs, nb.Location)
			}
			if nb.Char == "S" {
				Slocs = append(Slocs, nb.Location)
			}
		}

		if len(Mlocs) == 2 && len(Slocs) == 2 {

			if Mlocs[0].X == Mlocs[1].X || Mlocs[0].Y == Mlocs[1].Y {
				counter += 1
			}

		}
	}

	fmt.Printf("found %d total X-MAS", counter)
}

func PartTwo1(puzzle [][]string) {
	allM := FindAllLocations(puzzle, "M")

	var count = 0

	for _, m := range allM {
		diagNbs := FindDiagonalNeighboursOfLocation(puzzle, m)
		for _, nb := range diagNbs {
			if nb.Char == "A" {
				nextChar, err := FindNeighbourAtGivenDirection(puzzle, nb.Location, nb.Direction)
				if err != nil {
					log.Printf("oob: %s", err)
					continue
				}
				if nextChar.Char == "S" {
					// we found MAS -- now go back to A and check the other directions for
					aDiagNeighbs, err := FindBothNeighboursInOppositeDiagonalDirection(puzzle, nb.Location, nb.Direction)
					if err != nil {
						log.Printf("diag neighbs: %s", err)
						continue
					}

					var neighbChars []string

					for _, anb := range aDiagNeighbs {
						neighbChars = append(neighbChars, anb.Char)
					}

					if slices.Contains(neighbChars, "M") && slices.Contains(neighbChars, "S") {
						// we found a diagonal MAS
						count += 1
					}
				}

			}
		}
	}

	fmt.Printf("found %d X-MAS", count)
}

func PartOne(puzzle [][]string) {
	allXes := FindAllLocations(puzzle, "X")

	var count = 0

	for _, x := range allXes {
		xneighbs := FindAllNeighboursOfGivenLocation(puzzle, x)
		for _, nb := range xneighbs {
			if nb.Char == "M" {
				nextChar, err := FindNeighbourAtGivenDirection(puzzle, nb.Location, nb.Direction)
				if err != nil {
					log.Printf("out of bounds: %s", err)
					continue
				}
				if nextChar.Char == "A" {
					nextNextChar, err := FindNeighbourAtGivenDirection(puzzle, nextChar.Location, nextChar.Direction)
					if err != nil {
						log.Printf("out of bounds: %s", err)
						continue
					}
					if nextNextChar.Char == "S" {
						count += 1
					}
				}
			}
		}
	}

	fmt.Printf("Found %d total XMAS", count)
}

func FindBothNeighboursInOppositeDiagonalDirection(puzzle [][]string, pl PuzzleLocation, origDr Direction) ([]Neighbour, error) {

	oppDiagonalNeighbs := []Neighbour{}

	flippedX := Direction{origDr.X * -1, origDr.Y}
	flippedY := Direction{origDr.X, origDr.Y * -1}

	flippedXLocation := PuzzleLocation{
		Y: pl.Y + flippedX.Y, X: pl.X + flippedX.X,
	}

	flippedYLocation := PuzzleLocation{
		Y: pl.Y + flippedY.Y, X: pl.X + flippedY.X,
	}

	neighb1 := Neighbour{
		Char:      puzzle[flippedXLocation.Y][flippedXLocation.X],
		Location:  flippedXLocation,
		Direction: flippedX,
	}

	neighb2 := Neighbour{
		Char:      puzzle[flippedYLocation.Y][flippedYLocation.X],
		Location:  flippedYLocation,
		Direction: flippedY,
	}

	if TargetLocationIsInBounds(puzzle, neighb1.Location) && TargetLocationIsInBounds(puzzle, neighb2.Location) {
		oppDiagonalNeighbs = append(oppDiagonalNeighbs, neighb1)
		oppDiagonalNeighbs = append(oppDiagonalNeighbs, neighb2)
		return oppDiagonalNeighbs, nil
	}

	return nil, errors.New("one or more neighbours oob")
}

func FindNeighbourAtGivenDirection(puzzle [][]string, pl PuzzleLocation, dr Direction) (*Neighbour, error) {

	target := PuzzleLocation{X: pl.X + dr.X, Y: pl.Y + dr.Y}

	if TargetLocationIsInBounds(puzzle, target) {
		newNb := &Neighbour{
			Char:      puzzle[target.Y][target.X],
			Direction: dr,
			Location:  target,
		}
		return newNb, nil
	}

	return nil, errors.New("target is out of bounds")
}

func FindDiagonalNeighboursOfLocation(puzzle [][]string, pl PuzzleLocation) []Neighbour {
	var nbs []Neighbour

	for i := -1; i <= 1; i += 2 {
		for j := -1; j <= 1; j += 2 {
			target := PuzzleLocation{X: pl.X + j, Y: pl.Y + i}
			if TargetLocationIsInBounds(puzzle, target) {
				newNb := Neighbour{
					Char:      puzzle[target.Y][target.X],
					Direction: Direction{j, i},
					Location:  target,
				}
				nbs = append(nbs, newNb)
			}
		}
	}

	return nbs
}

func FindAllNeighboursOfGivenLocation(puzzle [][]string, pl PuzzleLocation) []Neighbour {
	var nbs []Neighbour

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			target := PuzzleLocation{X: pl.X + j, Y: pl.Y + i}
			if TargetLocationIsInBounds(puzzle, target) {
				newNb := Neighbour{
					Char:      puzzle[target.Y][target.X],
					Direction: Direction{j, i},
					Location:  target,
				}
				nbs = append(nbs, newNb)
			}
		}
	}

	return nbs
}

func TargetLocationIsInBounds(puzzle [][]string, target PuzzleLocation) bool {
	leftBound := 0
	UpperBound := 0
	RightBound := len(puzzle[0]) - 1
	LowerBound := len(puzzle) - 1

	if target.X < leftBound || target.X > RightBound || target.Y < UpperBound || target.Y > LowerBound {
		return false
	}

	return true
}
