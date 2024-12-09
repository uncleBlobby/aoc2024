package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	daily "github.com/uncleBlobby/aoc2024/internal"
)

// Disk Map
//
//	- digits alternate between indicating the length of a file and the length of a free space
//  - each file on the disk has an ID number based on the order of the files in the original list, starting with 0

// a file models the structure for any entity on the disk
// file id represents the order in the initial array
// file id == -1 represents a free space object???

type file struct {
	id     int
	length int
	isFree bool
}

func main() {
	fmt.Println("hello aoc2024 day9!")

	input := daily.GetInputFromFile("input")

	scanner := bufio.NewScanner(input)
	diskMap := []file{}
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		diskMapChars := strings.Split(scanner.Text(), "")

		emptySpaceCount := 0
		for id, char := range diskMapChars {
			fLen, err := strconv.Atoi(char)
			if err != nil {
				log.Printf("error converting char to int: %s", err)
				continue
			}
			f := file{
				id:     id - emptySpaceCount,
				length: fLen,
			}
			if id == 0 || id%2 == 0 {
				f.isFree = false
			} else {
				f.isFree = true
				f.id = -1
				emptySpaceCount += 1
			}
			diskMap = append(diskMap, f)
		}
	}

	fmt.Println(diskMap)

	fragmented := []string{}

	for _, f := range diskMap {
		if !f.isFree {
			//fmt.Printf("%d", f.id)
			for i := 0; i < f.length; i++ {
				fmt.Printf("%d", f.id)
				c := strconv.Itoa(f.id)
				fragmented = append(fragmented, c)
			}
		} else {
			emptySpaces := make([]string, f.length)
			for i := 0; i < f.length; i++ {
				emptySpaces = append(emptySpaces, ".")
				fmt.Printf("%s", ".")
				fragmented = append(fragmented, ".")
			}

			// fmt.Printf("%s", emptySpaces)
		}
	}

	fmt.Println()
	fmt.Println(fragmented)
	fmt.Println()

	// count all empty spaces in fragmented array

	emptyCount := 0
	for _, c := range fragmented {
		if c == "." {
			emptyCount += 1
		}
	}

	// while NOT all spaces are at end of array...
	// loop through array and move non empty entity to first empty spot

	fmt.Printf("LEN OF fragmented: %d", len(fragmented))

	//os.Exit(1)

	iterCount := 0
outer:
	for i := 0; i < len(fragmented); i++ {
		for j := len(fragmented) - 1; j >= 0; j-- {
			iterCount++

			if iterCount%10000 == 0 {
				//fmt.Printf("%d\n", iterCount)
				fmt.Printf("%f", (float32(iterCount)/float32((len(fragmented)*len(fragmented))))*100)
				fmt.Println("percent complete")
			}
			if fragmented[i] == "." && fragmented[j] != "." {
				fmt.Printf("swapping [%d] & [%d] {%s & %s}\n", i, j, fragmented[i], fragmented[j])
				fragmented[i], fragmented[j] = fragmented[j], fragmented[i]
				fString := strings.Join(fragmented, "")
				stringTailArr := []string{}
				for _ = range emptyCount {
					stringTailArr = append(stringTailArr, ".")
				}
				stringTail := strings.Join(stringTailArr, "")
				if strings.HasSuffix(fString, stringTail) {
					break outer
				}
			}
		}
	}

	fmt.Println(fragmented)

	unfragged := []int{}

	for _, char := range fragmented {
		n, err := strconv.Atoi(char)
		if err != nil {
			continue
		}
		unfragged = append(unfragged, n)
	}

	checkSum := 0

	for ind, num := range unfragged {
		checkSum += ind * num
	}

	fmt.Println(checkSum)
}
