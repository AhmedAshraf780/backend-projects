package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

// TODO: The process
// 1. We will Validate the args
// 2. We will see if this file exists or not
// 3. We will handle the get request and set the headers
// 4. We will serve this file
// 5. We will listen

func main() {
	// 1. Validations
	if len(os.Args) != 2 {
		fmt.Println("This command must have exactly one argument")
		return
	}

	// 1.1 Extract the argument
	filePath := os.Args[1]

	// 2. check existing file
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("the file path is not correct", filePath)
		return
	}

	// 3. We will handle the get request

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fileName := filepath.Base(filePath)

		w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
		w.Header().Set("Content-Type", "application/octet-stream")

		// here we need to get the file and serve it

		http.ServeFile(w, r, filePath)
	})

	// 4. We will be listenning

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Something Went wrong during listenning", err)
		return
	}
}
