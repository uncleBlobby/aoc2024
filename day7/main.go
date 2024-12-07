package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand/v2"
	"strconv"
	"strings"

	daily "github.com/uncleBlobby/aoc2024/internal"
)

// # of Possible Expressions
// if 2 operands:
//	[x] + [y]
// 	[x] * [y]
// if 3 operands:
//	[x] + [y] + [z]
//  [x] + [y] * [z]
//  [x] * [y] * [z]
//  [x] * [y] + [z]
// if 4 operands:
//  [x] + [y] + [z] + [a]
//  [x] + [y] + [z] * [a]
//  [x] + [y] * [z] * [a]
//  [x] * [y] * [z] * [a]
//  [x] * [y] * [z] + [a]
//  [x] * [y] + [z] + [a]
//  [x] + [y] * [z] + [a]
//  [x] * [y] + [z] * [a]

//945409439768
//948778272862
//975649463097
//975671981569

type Equation struct {
	Result   int
	Operands []int
}

func main() {
	fmt.Println("hello, aoc2024 day 7!")

	input := daily.GetInputFromFile("input")

	ops := map[string]func(int, int) int{}
	ops["*"] = func(a, b int) int {
		return a * b
	}
	ops["+"] = func(a, b int) int {
		return a + b
	}

	scanner := bufio.NewScanner(input)

	equations := []Equation{}

	for scanner.Scan() {
		fmt.Println(scanner.Text())

		line := scanner.Text()

		result := GetResult(line)

		fmt.Println(result)

		operands := GetOperands(line)

		fmt.Println(operands)

		equations = append(equations, Equation{
			Result:   result,
			Operands: operands,
		})

	}

	calibrationResult := 0

	for ind, eq := range equations {
		fmt.Printf("%d", ind)
		_, adder := Foo(eq, ops)
		calibrationResult += adder
	}

	log.Printf("calibration result: %d", calibrationResult)
}

func GetRandomOperatorFunc(ops map[string]func(a, b int) int) func(a, b int) int {
	r := rand.IntN(2)
	if r == 1 {
		return ops["*"]
	} else {
		return ops["+"]
	}
}

func Foo(equation Equation, ops map[string]func(a, b int) int) (bool, int) {
	var res int
	maxTries := 1000000
	tries := 0
	for res != equation.Result && tries < maxTries {
		for i := 0; i < len(equation.Operands)-1; i++ {
			if i == 0 {
				res = GetRandomOperatorFunc(ops)(equation.Operands[i], equation.Operands[i+1])
			} else {
				res = GetRandomOperatorFunc(ops)(res, equation.Operands[i+1])
			}
		}
		tries += 1
	}
	if tries < maxTries {
		log.Printf("found matching result: %v is a valid equation", equation)
		return true, res
	} else {
		log.Printf("too many tries, giving up")
		return false, 0
	}
}

func GetResult(line string) int {
	parts := strings.Split(line, ":")

	result, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Printf("error parsing result: %s", err)
	}

	return result
}

func GetOperands(line string) []int {
	parts := strings.Split(line, ":")

	nums := strings.Split(parts[1][1:], " ")

	operands := []int{}

	for _, num := range nums {
		n, err := strconv.Atoi(num)
		if err != nil {
			log.Printf("error parsing operand: %s", err)
		}
		operands = append(operands, n)
	}

	return operands
}
