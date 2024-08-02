# question given 
# League Backend Challenge

In main.go you will find a basic web server written in GoLang. It accepts a single request _/echo_. Extend the webservice with the ability to perform the following operations

Given an uploaded csv file
```
1,2,3
4,5,6
7,8,9
```

1. Echo (given)
    - Return the matrix as a string in matrix format.
    
    ```
    // Expected output
    1,2,3
    4,5,6
    7,8,9
    ``` 
2. Invert
    - Return the matrix as a string in matrix format where the columns and rows are inverted
    ```
    // Expected output
    1,4,7
    2,5,8
    3,6,9
    ``` 
3. Flatten
    - Return the matrix as a 1 line string, with values separated by commas.
    ```
    // Expected output
    1,2,3,4,5,6,7,8,9
    ``` 
4. Sum
    - Return the sum of the integers in the matrix
    ```
    // Expected output
    45
    ``` 
5. Multiply
    - Return the product of the integers in the matrix
    ```
    // Expected output
    362880
    ``` 

The input file to these functions is a matrix, of any dimension where the number of rows are equal to the number of columns (square). Each value is an integer, and there is no header row. matrix.csv is example valid input.  


# main.go file given 
package leapfrog-coding

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}
		defer file.Close()
		records, err := csv.NewReader(file).ReadAll()
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}
		var response string
		for _, row := range records {
			response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
		}
		fmt.Fprint(w, response)
	})
	http.ListenAndServe(":8080", nil)
}

# matrix.csv given same as in test folder


Run web server
```
go run .
```

Send request
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
```

## What we're looking for

- The solution runs
- The solution performs all cases correctly
- The code is easy to read
- The code is reasonably documented
- The code is tested
- The code is robust and handles invalid input and provides helpful error messages

# Prerequisites

- Go 1.16 or higher

# Setup

- unzip the leapfrog-coding.zip once downloaded from the google drive

# ways to run 

-use "go run ." to Run the server

# curls commands to use in command promt
"pathtoyourfolder" replace it with actual path

-for echo
curl -F "file=@pathtoyourfolder\\leapfrog-coding\\testdata\\matrix.csv" "localhost:8080/echo"

-for invert
curl -F "file=@pathtoyourfolder\\leapfrog-coding\\testdata\\matrix.csv" "localhost:8080/invert"

-for flatten
curl -F "file=@pathtoyourfolder\\leapfrog-coding\\testdata\\matrix.csv" "localhost:8080/flatten"

-for sum
curl -F "file=@pathtoyourfolder\\leapfrog-coding\\testdata\\matrix.csv" "localhost:8080/sum"

-for multiply
curl -F "file=@pathtoyourfolder\\leapfrog-coding\\testdata\\matrix.csv" "localhost:8080/multiply"

# Testing
-run the test using from the folder
go test ./...