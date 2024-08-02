package utils

import (
	"encoding/csv"
	"errors"
	"net/http"
	"strconv"
)

// ReadCSV reads the CSV file from the HTTP request and parses it into a 2D slice of integers.
func ReadCSV(r *http.Request) ([][]int, error) {
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var matrix [][]int
	for _, record := range records {
		var row []int
		for _, value := range record {
			num, err := strconv.Atoi(value)
			if err != nil {
				return nil, errors.New("invalid integer in CSV file")
			}
			row = append(row, num)
		}
		matrix = append(matrix, row)
	}

	return matrix, nil
}
