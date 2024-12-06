package main

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
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

type StepInformation struct {
	Position  Position
	Direction Direction
}

type Guard struct {
	Position Position
	History  []StepInformation
	Facing   Direction
}

// type WorldMap struct {
// 	Positions map[Position]string
// }

type Game struct {
	WorldMap      map[Position]string
	Guard         Guard
	Clone         Guard
	worldBoundary Position
	currentFrame  int
	DrawnPath     []Position
}

func (g *Game) Update() error {
	g.currentFrame += 1
	time.Sleep(66 * time.Millisecond)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, world!")

	vector.DrawFilledRect(screen, 100, 100, 10, 10, color.RGBA{255, 0, 0, 255}, false)

	for position, char := range g.WorldMap {
		if char == "." {
			vector.DrawFilledRect(screen, float32(position.X)*2, float32(position.Y)*2, 2, 2, color.RGBA{239, 239, 240, 1}, false)
		}
		if char == "#" {
			vector.DrawFilledRect(screen, float32(position.X)*2, float32(position.Y)*2, 2, 2, color.RGBA{255, 148, 112, 1}, false)
		}
	}

	vector.DrawFilledRect(screen, float32(g.Guard.History[g.currentFrame].Position.X)*2, float32(g.Guard.History[g.currentFrame].Position.Y)*2, 2, 2, color.RGBA{45, 85, 255, 255}, false)

	g.DrawnPath = append(g.DrawnPath, g.Guard.History[g.currentFrame].Position)

	for _, position := range g.DrawnPath {
		vector.DrawFilledRect(screen, float32(position.X)*2, float32(position.Y)*2, 2, 2, color.RGBA{45, 85, 255, 150}, false)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	fmt.Println("hello, aoc2024 day6!")

	input := daily.GetInputFromFile("input")

	defer input.Close()

	scanner := bufio.NewScanner(input)

	lineCount := 0

	worldBoundary := Position{}

	game := &Game{
		currentFrame: 1,
	}

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
				guard.History = append(guard.History, StepInformation{
					Position:  guard.Position,
					Direction: guard.Facing,
				})
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
		game.worldBoundary = worldBoundary
	}

	game.WorldMap = world
	game.Guard = guard

	newObstacles := []Position{}

	for game.Guard.InBounds(game.worldBoundary) {
		log.Printf("%d percent complete...", len(game.Guard.History)+1/5067)
		game.Clone = game.Guard
		anchor := StepInformation{
			guard.Position,
			guard.Facing,
		}
		game.Clone.Facing = guard.Facing.Rotate90()
		for game.Clone.InBounds(game.worldBoundary) {
			//fmt.Printf("clone guard %v\n", guard)
			game.Clone.Travel(game.WorldMap)
		}

		for _, historyStep := range game.Clone.History {
			//log.Printf("checking path steps %v\n", historyStep)
			if historyStep == anchor {
				// found a loop if we put obstacle in facing direction of anchor position
				log.Printf("found new obstacle position")
				newObstacles = append(newObstacles, Position{anchor.Position.X + anchor.Direction.X, anchor.Position.Y + anchor.Direction.Y})
				break
			}
		}

		game.Guard.Travel(game.WorldMap)
		//time.Sleep(1 * time.Second)
	}

	path := game.Guard.FindUniquePathPositions()

	fmt.Printf("Guard visits %d unique positions before leaving map\n", len(path)+1)

	for i := 0; i < len(newObstacles); i++ {
		if contains(newObstacles[i+1:], newObstacles[i]) {
			newObstacles = remove(newObstacles, i)
			i--
		}
	}

	fmt.Printf("Can add obstacles at %d positions", len(newObstacles))

	ebiten.SetWindowSize(1000, 1000)
	ebiten.SetWindowTitle("Hello, World")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}

func contains(p []Position, t Position) bool {
	for _, a := range p {
		if a == t {
			return true
		}
	}
	return false
}

func remove(slice []Position, idx int) []Position {
	return append(slice[:idx], slice[idx+1:]...)
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
		if path[visit.Position] == 0 {
			path[visit.Position] += 1
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

	// // before we do anything, pretend there is an obstacle in front of us and we have to turn right now -- then check if that path leads us back to our path history
	// anchor := StepInformation{
	// 	Position:  g.Position,
	// 	Direction: g.Facing,
	// }

	// cloneHist := []StepInformation{}
	// _ = copy(g.History, cloneHist)
	// clone := Guard{
	// 	Position: g.Position,
	// 	History:  cloneHist,
	// 	Facing:   g.Facing,
	// }

	// clone.Facing.Rotate90()
	// clone.Travel(world)

	// if next position is clear ("."), move into that position
	if nextPositionStatus == "." {
		//log.Printf("guard moving to %v", nextPosition)
		g.History = append(g.History, StepInformation{
			Position:  g.Position,
			Direction: g.Facing,
		})
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
