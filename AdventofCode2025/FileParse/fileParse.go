package fileparse

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func FileParse(InputFile string) []string {
	file, err := os.Open(InputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func Grid(InputFile string, seperator string) [][]string {
	file, err := os.Open(InputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		lines = append(lines, strings.Split(scanner.Text(), seperator))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines

}
