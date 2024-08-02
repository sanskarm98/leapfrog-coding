package operations

import (
	"testing"
)

func TestFormatMatrix(t *testing.T) {
	ops := NewMatrixOps()
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	expected := "1,2,3\n4,5,6\n7,8,9\n"
	result := ops.FormatMatrix(matrix)
	if result != expected {
		t.Errorf("FormatMatrix returned unexpected result: got %v want %v", result, expected)
	}
}

func TestInvertMatrix(t *testing.T) {
	ops := NewMatrixOps()
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	expected := [][]int{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	}
	result := ops.InvertMatrix(matrix)
	if !equalMatrix(result, expected) {
		t.Errorf("InvertMatrix returned unexpected result: got %v want %v", result, expected)
	}
}

func TestFlattenMatrix(t *testing.T) {
	ops := NewMatrixOps()
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	expected := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	result := ops.FlattenMatrix(matrix)
	if !equalStringSlice(result, expected) {
		t.Errorf("FlattenMatrix returned unexpected result: got %v want %v", result, expected)
	}
}

func TestSumMatrix(t *testing.T) {
	ops := NewMatrixOps()
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	expected := 45
	result := ops.SumMatrix(matrix)
	if result != expected {
		t.Errorf("SumMatrix returned unexpected result: got %v want %v", result, expected)
	}
}

func TestMultiplyMatrix(t *testing.T) {
	ops := NewMatrixOps()
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	expected := 362880
	result := ops.MultiplyMatrix(matrix)
	if result != expected {
		t.Errorf("MultiplyMatrix returned unexpected result: got %v want %v", result, expected)
	}
}

// Helper function to compare 2D slices
func equalMatrix(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

// Helper function to compare string slices
func equalStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
