package filemanager

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type FileManager struct{}

func checkParam[T any](param T, err error) T {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return param
}

func trimSpacesInSplitedFile(strs []string) []string {
	trimmedStr := []string{}
	for _, str := range strs {
		trimmedStr = append(trimmedStr, strings.TrimSpace(str))
	}
	return trimmedStr
}

func readFileContent(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanLines)
	lines := make([]string, 0)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	splited := make([][]string, 0)
	for _, line := range lines {

		splited = append(splited, trimSpacesInSplitedFile(strings.Split(line, ";")))
	}
	return splited
}
