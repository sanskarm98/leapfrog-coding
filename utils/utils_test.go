package utils

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"testing"
)

const testCSV = `1,2,3
4,5,6
7,8,9`

func TestReadCSV(t *testing.T) {
	req, err := newFileUploadRequest("/upload", "file", "matrix.csv", testCSV)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	matrix, err := ReadCSV(req)
	if err != nil {
		t.Fatalf("ReadCSV returned an error: %v", err)
	}

	expected := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	if !equalMatrix(matrix, expected) {
		t.Errorf("ReadCSV returned unexpected matrix: got %v want %v", matrix, expected)
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
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

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
