package main

import (
	"bufio"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	daily "github.com/uncleBlobby/aoc2024/internal"
)

type Results struct {
	ValidUpdates   []Update
	InvalidUpdates []Update
}

type PrintQueue struct {
	Rules       []PageOrderingRule
	Updates     []Update
	SortedRules []int
}

type PageOrderingRule struct {
	First  int
	Second int
}

type Update struct {
	Pages []int
}

func main() {
	fmt.Println("hello, aoc2024 day5!")

	pq := ParsePrintQueue("input")

	results := pq.PartOne()

	results = pq.PartTwo(results)

	var totalSum = 0

	for _, validUpdate := range results.ValidUpdates {
		middleIndex := len(validUpdate.Pages) / 2
		middleValue := validUpdate.Pages[middleIndex]
		totalSum += middleValue
	}

	fmt.Printf("%d total sum after part2\n", totalSum)
}

func (pq *PrintQueue) PartTwo(results Results) Results {
	// invalidUpdates := []Update{}
	// for _, update := range pq.Updates {
	// 	if !slices.Contains(validUpdates, update) {
	// 		invalidUpdates = append(invalidUpdates, update)
	// 	}
	// }

	for _, update := range results.InvalidUpdates {

		// clone := Update{}
		// _ = copy(clone, update)

		for !pq.UpdateIsValid(update) {
			for i := 0; i < len(update.Pages)-1; i++ {
				firstPage := update.Pages[i]
				secondPage := update.Pages[i+1]
				thisPair := PageOrderingRule{
					First:  firstPage,
					Second: secondPage,
				}

				for _, rule := range pq.Rules {
					if thisPair.First == rule.First && thisPair.Second == rule.Second {
						//ordering is correct

						continue
					}
					if thisPair.First == rule.Second && thisPair.Second == rule.First {
						//ordering is incorrect, entire page list is invalid
						//break middle

						update.Pages[i], update.Pages[i+1] = update.Pages[i+1], update.Pages[i]
					}
				}
			}
		}

		if pq.UpdateIsValid(update) {
			results.ValidUpdates = append(results.ValidUpdates, update)
		}

	}

	fmt.Printf("%d valid after part2\n", len(results.ValidUpdates))
	return results
}

func (pq *PrintQueue) PartOne() Results {
	validUpdates := []Update{}
	invalidUpdates := []Update{}

	for _, update := range pq.Updates {

		updateIsValid := true
		//middle:
		for i := 0; i <= len(update.Pages)-2; i++ {
			firstPage := update.Pages[i]
			secondPage := update.Pages[i+1]
			thisPair := PageOrderingRule{
				First:  firstPage,
				Second: secondPage,
			}

			for _, rule := range pq.Rules {
				if thisPair.First == rule.First && thisPair.Second == rule.Second {
					//ordering is correct

					continue
				}
				if thisPair.First == rule.Second && thisPair.Second == rule.First {
					//ordering is incorrect, entire page list is invalid
					updateIsValid = false
					//break middle
				}
			}

		}
		if updateIsValid {
			validUpdates = append(validUpdates, update)
		} else {
			invalidUpdates = append(invalidUpdates, update)
		}

	}

	for _, validUpdate := range validUpdates {
		log.Printf("%v is a valid set\n", validUpdate)
	}

	var sumOfMiddles = 0
	for _, validUpdate := range validUpdates {
		middleIndex := len(validUpdate.Pages) / 2
		sumOfMiddles += validUpdate.Pages[middleIndex]
	}

	fmt.Printf("total sum: %d\n", sumOfMiddles)
	fmt.Printf("%d valid after part1", len(validUpdates))

	return Results{
		ValidUpdates:   validUpdates,
		InvalidUpdates: invalidUpdates,
	}
}

func (p *PrintQueue) SortRules() {

	for _, rule := range p.Rules {
		if !slices.Contains(p.SortedRules, rule.First) {
			p.SortedRules = append(p.SortedRules, rule.First)
		}
		if !slices.Contains(p.SortedRules, rule.Second) {
			p.SortedRules = append(p.SortedRules, rule.Second)
		}
	}

	for _, rule := range p.Rules {
		firstInd := slices.Index(p.SortedRules, rule.First)

		secondInd := slices.Index(p.SortedRules, rule.First)

		if firstInd > secondInd {
			p.SortedRules[firstInd], p.SortedRules[secondInd] = p.SortedRules[secondInd], p.SortedRules[firstInd]
		}
	}
}

func ParsePrintQueue(inputFile string) PrintQueue {
	input := daily.GetInputFromFile(inputFile)

	scanner := bufio.NewScanner(input)

	rules := []PageOrderingRule{}
	updates := []Update{}

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		line := scanner.Text()

		if !strings.Contains(line, ",") {
			if len(line) < 2 {
				continue
			}
			// any string without a comma is a page ordering rule
			pageNumbers := strings.Split(line, "|")
			first, err := strconv.Atoi(pageNumbers[0])
			if err != nil {
				log.Printf("error parsing int: %s", err)
				continue
			}
			second, err := strconv.Atoi(pageNumbers[1])
			if err != nil {
				log.Printf("error parsing int: %s", err)
				continue
			}
			newRule := PageOrderingRule{
				First:  first,
				Second: second,
			}
			rules = append(rules, newRule)
		}

		// any string with a comma is an update list
		if strings.Contains(line, ",") {
			newUpdate := Update{}

			pageNumbers := strings.Split(line, ",")
			for _, pageNum := range pageNumbers {
				pn, err := strconv.Atoi(pageNum)
				if err != nil {
					log.Printf("error parsing int: %s", err)
					continue
				}
				newUpdate.Pages = append(newUpdate.Pages, pn)
			}
			updates = append(updates, newUpdate)
		}
	}

	fmt.Println("Page Rules:")
	for _, rule := range rules {
		fmt.Println(rule)
	}

	fmt.Println("Updates:")
	for _, update := range updates {
		fmt.Println(update)
	}

	return PrintQueue{
		Rules:   rules,
		Updates: updates,
	}
}

func (pq *PrintQueue) UpdateIsValid(update Update) bool {
	updateIsValid := true

	for i := 0; i <= len(update.Pages)-2; i++ {
		firstPage := update.Pages[i]
		secondPage := update.Pages[i+1]
		thisPair := PageOrderingRule{
			First:  firstPage,
			Second: secondPage,
		}

		for _, rule := range pq.Rules {
			if thisPair.First == rule.First && thisPair.Second == rule.Second {
				//ordering is correct

				continue
			}
			if thisPair.First == rule.Second && thisPair.Second == rule.First {
				//ordering is incorrect, entire page list is invalid
				updateIsValid = false
				//break middle
			}
		}
	}

	return updateIsValid
}
