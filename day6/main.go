package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"

	daily "github.com/uncleBlobby/aoc2024/internal"
)

type Position struct {
	X int
	Y int
}

type Direction struct {
	X int
	Y int
}

type Guard struct {
	Position Position
	History  []Position
	Facing   Direction
}

// type WorldMap struct {
// 	Positions map[Position]string
// }

func main() {
	fmt.Println("hello, aoc2024 day6!")

	input := daily.GetInputFromFile("input")

	defer input.Close()

	scanner := bufio.NewScanner(input)

	lineCount := 0

	worldBoundary := Position{}

	world := map[Position]string{}
	guard := Guard{}

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		line := scanner.Text()
		chars := strings.Split(line, "")

		for ind, char := range chars {

			if char == "^" {
				guard.Position = Position{
					X: ind,
					Y: lineCount,
				}
				guard.Facing = Direction{
					X: 0,
					Y: -1,
				}
				guard.History = append(guard.History, guard.Position)
				thisPosition := Position{
					X: ind,
					Y: lineCount,
				}
				world[thisPosition] = "."
			} else {
				thisPosition := Position{
					X: ind,
					Y: lineCount,
				}
				world[thisPosition] = char
			}

		}
		lineCount += 1
		worldBoundary.X = len(chars)
		worldBoundary.Y = lineCount
	}

	for guard.InBounds(worldBoundary) {
		guard.Travel(world)
	}

	path := guard.FindUniquePathPositions()

	fmt.Printf("Guard visits %d unique positions before leaving map\n", len(path)+1)
}

func (g *Guard) InBounds(boundary Position) bool {
	if g.Position.X < 0 || g.Position.Y < 0 || g.Position.X >= boundary.X-1 || g.Position.Y >= boundary.Y-1 {

		log.Printf("guard left the arena at position: %v", g.Position)
		return false
	}

	return true
}

func (g *Guard) FindUniquePathPositions() map[Position]int {
	log.Println("computing unique positions in path...")
	path := map[Position]int{}
	for _, visit := range g.History {
		if path[visit] == 0 {
			path[visit] += 1
		}
	}
	return path
}

func (g *Guard) Travel(world map[Position]string) {
	// find the next square based on direction facing
	nextPosition := Position{
		X: g.Position.X + g.Facing.X,
		Y: g.Position.Y + g.Facing.Y,
	}

	// check whether the next position is clear or obstacle
	nextPositionStatus := world[nextPosition]

	// if next position is clear ("."), move into that position
	if nextPositionStatus == "." {
		log.Printf("guard moving to %v", nextPosition)
		g.History = append(g.History, g.Position)
		g.Position = nextPosition
	}
	// if next position is obstacle ("#"), turn right ninety degrees and start over
	if nextPositionStatus == "#" {
		g.Facing = g.Facing.Rotate90()
	}
}

func (d *Direction) Rotate90() Direction {
	// from up to right
	if d.X == 0 && d.Y == -1 {
		return Direction{
			X: 1,
			Y: 0,
		}
	}
	// from right to down
	if d.X == 1 && d.Y == 0 {
		return Direction{
			X: 0,
			Y: 1,
		}
	}
	// from down to left
	if d.X == 0 && d.Y == 1 {
		return Direction{
			X: -1,
			Y: 0,
		}
	}

	// from left to up
	if d.X == -1 && d.Y == 0 {
		return Direction{
			X: 0,
			Y: -1,
		}
	}

	// Should be unreachable code
	fmt.Println("hit panic")
	panic("direction error")
	// return Direction{
	// 	X: 0,
	// 	Y: 0,
	// }
}
