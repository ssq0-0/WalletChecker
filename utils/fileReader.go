package utils

import (
	"bufio"
	"os"
)

func FileReader(filename string) ([]string, error) {
	var lines []string
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func GetRootDir() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}
