package handlers

import (
	"Backend_Challenge/operations"
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// InvertHandler handles the /invert endpoint (invert the matrix)
func InvertHandler(w http.ResponseWriter, r *http.Request, ops operations.MatrixOperations) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, fmt.Sprintf("error %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		http.Error(w, fmt.Sprintf("error %s", err.Error()), http.StatusBadRequest)
		return
	}

	matrix := parseCSV(records)
	inverted := ops.InvertMatrix(matrix)
	response := ops.FormatMatrix(inverted)
	fmt.Fprint(w, response)
}

// FlattenHandler handles the /flatten endpoint (flatten the matrix)
func FlattenHandler(w http.ResponseWriter, r *http.Request, ops operations.MatrixOperations) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, fmt.Sprintf("error %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		http.Error(w, fmt.Sprintf("error %s", err.Error()), http.StatusBadRequest)
		return
	}

	matrix := parseCSV(records)
	flattened := ops.FlattenMatrix(matrix)
	response := strings.Join(flattened, ",")
	fmt.Fprint(w, response)
}

// SumHandler handles the /sum endpoint  (sum of elemets of the matrix)
func SumHandler(w http.ResponseWriter, r *http.Request, ops operations.MatrixOperations) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, fmt.Sprintf("error %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		http.Error(w, fmt.Sprintf("error %s", err.Error()), http.StatusBadRequest)
		return
	}

	matrix := parseCSV(records)
	sum := ops.SumMatrix(matrix)
	fmt.Fprint(w, strconv.Itoa(sum))
}

// MultiplyHandler handles the /multiply endpoint  (product of elemets of the matrix)
func MultiplyHandler(w http.ResponseWriter, r *http.Request, ops operations.MatrixOperations) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, fmt.Sprintf("error %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		http.Error(w, fmt.Sprintf("error %s", err.Error()), http.StatusBadRequest)
		return
	}

	matrix := parseCSV(records)
	product := ops.MultiplyMatrix(matrix)
	fmt.Fprint(w, strconv.Itoa(product))
}

// EchoHandler handles the /echo endpoint  (display the matrix)
func EchoHandler(w http.ResponseWriter, r *http.Request, ops operations.MatrixOperations) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, fmt.Sprintf("error %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		http.Error(w, fmt.Sprintf("error %s", err.Error()), http.StatusBadRequest)
		return
	}

	matrix := parseCSV(records)
	response := ops.FormatMatrix(matrix)
	fmt.Fprint(w, response)
}

// Helper function to parse CSV records into a matrix of integers
func parseCSV(records [][]string) [][]int {
	matrix := make([][]int, len(records))
	for i, row := range records {
		matrix[i] = make([]int, len(row))
		for j, value := range row {
			val, _ := strconv.Atoi(value)
			matrix[i][j] = val
		}
	}
	return matrix
}
