package ascii

import (
	"bufio"
	"log"
	"os"
)

func Convert(file string) []string {
	f, err := os.Open(file)

	if err != nil {
		log.Fatalf("Error: failed to open file %s", err)
	}

	defer f.Close()

	var lines []string

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: failed to read file: %s", err)
	}
	return lines
}
