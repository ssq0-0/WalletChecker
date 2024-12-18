package utils

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
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

func FileWriter(filename string, data map[string]float64) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0664)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл %s: %v", filename, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"Address", "Amount"}
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("не удалось записать заголовки: %w", err)
	}

	for addr, amount := range data {
		record := []string{
			addr,
			strconv.FormatFloat(amount, 'f', 2, 64),
		}

		if err := writer.Write(record); err != nil {
			return fmt.Errorf("не удалось записать %v: %v", record, err)
		}
	}

	return nil
}
