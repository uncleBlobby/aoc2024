package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	daily "github.com/uncleBlobby/aoc2024/internal"
)

type Position struct {
	Height int
	X      int
	Y      int
}

type TrailHead struct {
	Position Position
	Paths    []Path
	Score    int
}

type Path struct {
	Steps    []Position
	Complete bool
}

func main() {
	fmt.Println("hello, aoc2024 day 10!")

	input := daily.GetInputFromFile("testdata")

	scanner := bufio.NewScanner(input)

	//					[y][x]
	topology := [][]Position{}

	lineCount := 0
	for scanner.Scan() {
		//fmt.Println(scanner.Text())

		line := scanner.Text()

		chars := strings.Split(line, "")

		row := []Position{}

		for ind, char := range chars {
			num, err := strconv.Atoi(char)
			if err != nil {
				log.Printf("error converting to int: %s", err)
				continue
			}
			row = append(row, Position{
				Height: num,
				X:      ind,
				Y:      lineCount,
			})
		}

		topology = append(topology, row)
		lineCount += 1
	}

	for _, row := range topology {
		fmt.Println(row)
	}

	//fmt.Println("All Trailheads:")
	//fmt.Printf("%#v\n", FindAllTrailheads(topology))

	//trailHeads := FindAllTrailheads(topology)

}

// DFS\

func DFS(topo [][]Position) []Path {
	startingNodes := FindAllTrailheads(topo)

	stack := []Position{}

	for _, node := range startingNodes {
		nbs, err := Get1HeightNeighbours(node.Position, topo)
		if err != nil {
			continue
		}
		for _, nb := range nbs {
			stack = append(stack, nb)
		}
	}

	// pop a node from the stack to visit next

	for len(stack) > 0 {
		next := stack[len(stack)-1]

		nbs, err := Get1HeightNeighbours(next, topo)
		if err != nil {
			return nil
		}
		for _, nb := range nbs {
			stack = append(stack, nb)
		}
		stack = stack[:len(stack)-1]
	}

	return nil
}

// 1. Pick a starting node and push all its adjacent nodes onto a stack
// 			starting node := any trailhead
//			adjacent nodes := all 1 height neighbours

// 2. Pop a node from the stack <- next node to visit
// 3. Push all adjacent nodes of the selected node into the stack
// 4. Repeat until stack is empty
// 5. Mark visited nodes to prevent visiting the same node more than once

func FindAllTrailheads(topo [][]Position) []TrailHead {
	heads := []Position{}

	for y := 0; y < len(topo); y++ {
		for x := 0; x < len(topo[y]); x++ {
			if topo[y][x].Height == 0 {
				heads = append(heads, topo[y][x])
			}
		}
	}

	trailHeads := []TrailHead{}

	for _, thp := range heads {
		t := TrailHead{
			Position: thp,
		}
		trailHeads = append(trailHeads, t)
	}

	return trailHeads
}

func GetAll4Neighbours(start Position, topo [][]Position) ([]Position, error) {
	if !start.IsInBounds(topo) {
		return nil, errors.New(fmt.Sprintf("Cannot get Neighbours: Start Position %v is out of bounds", start))
	}

	nbs := []Position{}

	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			if x == -1 && y == 0 || x == 1 && y == 0 || x == 0 && y == 1 || x == 0 && y == -1 {
				target := Position{
					Height: topo[y][x].Height,
					X:      x,
					Y:      y,
				}
				if target.IsInBounds(topo) {
					nbs = append(nbs, target)
				}
			}
		}
	}

	return nbs, nil
}

func Get1HeightNeighbours(start Position, topo [][]Position) ([]Position, error) {

	allNbs, err := GetAll4Neighbours(start, topo)
	if err != nil {
		return nil, err
	}

	if !start.IsInBounds(topo) {
		return nil, errors.New(fmt.Sprintf("Cannot get Neighbours: Start Position %v is out of bounds", start))
	}

	heightNbs := []Position{}

	for _, neighb := range allNbs {
		if neighb.Height == start.Height+1 {
			heightNbs = append(heightNbs, neighb)
		}
	}

	return heightNbs, nil
}

func (p *Position) IsInBounds(topo [][]Position) bool {
	if p.X < 0 || p.Y < 0 {
		return false
	}
	if p.X >= len(topo[0]) || p.Y >= len(topo) {
		return false
	}

	return true
}
