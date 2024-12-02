package daily

import (
	"log"
	"os"
)

func GetInputFromFile(filename string) *os.File {
	input, err := os.Open(filename)
	if err != nil {
		log.Printf("error opening input file, please check filename and path")
		log.Printf("os.Open: %s", err)
		os.Exit(1)
	}

	return input
}
