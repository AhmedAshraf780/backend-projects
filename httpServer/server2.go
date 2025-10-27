package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Simple HTML upload form
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head><title>Upload File</title></head>
		<body>
			<h1>Upload File to Laptop</h1>
			<form enctype="multipart/form-data" action="/upload" method="post">
				<input type="file" name="uploadfile" />
				<input type="submit" value="Upload" />
			</form>
		</body>
		</html>
		`)
	})

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		// Parse uploaded file
		file, header, err := r.FormFile("uploadfile")
		if err != nil {
			http.Error(w, "Failed to read uploaded file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Save to current directory
		dstPath := filepath.Join("/home/mada/Pictures", header.Filename)
		dst, err := os.Create(dstPath)
		if err != nil {
			http.Error(w, "Failed to save file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			http.Error(w, "Failed to write file", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "File uploaded successfully: %s", header.Filename)
		fmt.Println("Received file:", header.Filename)
	})

	fmt.Println("Server started at http://localhost:4000")
	fmt.Println("Open this from your phone (using your laptop’s IP) to upload files.")

	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
