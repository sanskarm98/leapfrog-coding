package operations

import (
	"strconv"
)

// MatrixOperations defines the methods for matrix operations
type MatrixOperations interface {
	FormatMatrix(matrix [][]int) string
	InvertMatrix(matrix [][]int) [][]int
	FlattenMatrix(matrix [][]int) []string
	SumMatrix(matrix [][]int) int
	MultiplyMatrix(matrix [][]int) int
}

// MatrixOps provides implementations for matrix operations
type MatrixOps struct{}

// NewMatrixOps returns a new instance of MatrixOps
func NewMatrixOps() *MatrixOps {
	return &MatrixOps{}
}

// FormatMatrix formats the matrix as a string with rows joined by new lines
func (m *MatrixOps) FormatMatrix(matrix [][]int) string {
	var result string
	for _, row := range matrix {
		for i, value := range row {
			if i > 0 {
				result += ","
			}
			result += strconv.Itoa(value)
		}
		result += "\n"
	}
	return result
}

// InvertMatrix returns the transpose of the matrix
func (m *MatrixOps) InvertMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return [][]int{}
	}

	rowCount := len(matrix)
	colCount := len(matrix[0])
	inverted := make([][]int, colCount)
	for i := range inverted {
		inverted[i] = make([]int, rowCount)
	}

	for r := 0; r < rowCount; r++ {
		for c := 0; c < colCount; c++ {
			inverted[c][r] = matrix[r][c]
		}
	}

	return inverted
}

// FlattenMatrix returns a flattened slice of the matrix values
func (m *MatrixOps) FlattenMatrix(matrix [][]int) []string {
	var result []string
	for _, row := range matrix {
		for _, value := range row {
			result = append(result, strconv.Itoa(value))
		}
	}
	return result
}

// SumMatrix returns the sum of all matrix values
func (m *MatrixOps) SumMatrix(matrix [][]int) int {
	sum := 0
	for _, row := range matrix {
		for _, value := range row {
			sum += value
		}
	}
	return sum
}

// MultiplyMatrix returns the product of all matrix values
func (m *MatrixOps) MultiplyMatrix(matrix [][]int) int {
	product := 1
	for _, row := range matrix {
		for _, value := range row {
			product *= value
		}
	}
	return product
}
