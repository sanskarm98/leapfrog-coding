package handlers

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test data for CSV file
const testCSV = `1,2,3
4,5,6
7,8,9`

const formatedmatrix = "1,4,7\n2,5,8\n3,6,9\n"

const matrixfilename = "matrix.csv"

// MockMatrixOps implements the MatrixOperations interface for testing
type MockMatrixOps struct{}

func (m *MockMatrixOps) FormatMatrix(matrix [][]int) string {
	return formatedmatrix
}

func (m *MockMatrixOps) InvertMatrix(matrix [][]int) [][]int {
	return [][]int{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	}
}

func (m *MockMatrixOps) FlattenMatrix(matrix [][]int) []string {
	return []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
}

func (m *MockMatrixOps) SumMatrix(matrix [][]int) int {
	return 45
}

func (m *MockMatrixOps) MultiplyMatrix(matrix [][]int) int {
	return 362880
}

// Test the EchoHandler
func TestEchoHandler(t *testing.T) {
	mockOps := &MockMatrixOps{}

	req, err := newFileUploadRequest("/echo", "file", matrixfilename, testCSV)
	if err != nil {
		t.Fatalf(formattestError(err))
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		EchoHandler(w, r, mockOps)
	})
	handler.ServeHTTP(rr, req)

	expected := formatedmatrix // Update this based on the mock implementation
	if rr.Body.String() != expected {
		t.Errorf("EchoHandler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// Test the InvertHandler
func TestInvertHandler(t *testing.T) {
	mockOps := &MockMatrixOps{}

	req, err := newFileUploadRequest("/invert", "file", matrixfilename, testCSV)
	if err != nil {
		t.Fatalf(formattestError(err))
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		InvertHandler(w, r, mockOps)
	})
	handler.ServeHTTP(rr, req)

	expected := formatedmatrix
	if rr.Body.String() != expected {
		t.Errorf("InvertHandler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// Test the FlattenHandler
func TestFlattenHandler(t *testing.T) {
	mockOps := &MockMatrixOps{}

	req, err := newFileUploadRequest("/flatten", "file", matrixfilename, testCSV)
	if err != nil {
		t.Fatalf(formattestError(err))
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		FlattenHandler(w, r, mockOps)
	})
	handler.ServeHTTP(rr, req)

	expected := "1,2,3,4,5,6,7,8,9"
	if rr.Body.String() != expected {
		t.Errorf("FlattenHandler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// Test the SumHandler
func TestSumHandler(t *testing.T) {
	mockOps := &MockMatrixOps{}

	req, err := newFileUploadRequest("/sum", "file", matrixfilename, testCSV)
	if err != nil {
		t.Fatalf(formattestError(err))
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		SumHandler(w, r, mockOps)
	})
	handler.ServeHTTP(rr, req)

	expected := "45"
	if rr.Body.String() != expected {
		t.Errorf("SumHandler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// Test the MultiplyHandler
func TestMultiplyHandler(t *testing.T) {
	mockOps := &MockMatrixOps{}

	req, err := newFileUploadRequest("/multiply", "file", matrixfilename, testCSV)
	if err != nil {
		t.Fatalf(formattestError(err))
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		MultiplyHandler(w, r, mockOps)
	})
	handler.ServeHTTP(rr, req)

	expected := "362880"
	if rr.Body.String() != expected {
		t.Errorf("MultiplyHandler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// Helper function to create a multipart file upload request
func newFileUploadRequest(uri, paramName, fileName, fileContents string) (*http.Request, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, fileName)
	if err != nil {
		return nil, err
	}
	_, err = part.Write([]byte(fileContents))
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, nil
}
func formattestError(err error) string {
	return fmt.Sprintf("Failed to create request: %v", err)
}
