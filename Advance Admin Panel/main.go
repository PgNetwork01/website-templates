package main

import (
	"fmt"
	"html/template"
	"net/http"
	"log"
	"os"

	"github.com/wellington/go-libsass"
)


func compileSCSS() (string, error) {
	// SCSS file path
	scssFile := "style.scss"

	// Read the SCSS file
	scssData, err := os.ReadFile(scssFile)
	if err != nil {
		return "", fmt.Errorf("unable to read SCSS file: %w", err)
	}

	// Compile SCSS to CSS
	result, err := libsass.Compile(string(scssData))
	if err != nil {
		return "", fmt.Errorf("failed to compile SCSS: %w", err)
	}

	return result, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Compile SCSS to CSS
	css, err := compileSCSS()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Serve the compiled CSS
	w.Header().Set("Content-Type", "text/css")
	w.Write([]byte(css))
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	// Parse the index.html template
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template
	tmpl.Execute(w, nil)
}

func main() {
	// Handle requests for SCSS-to-CSS
	http.HandleFunc("/style.css", handler)

	// Handle requests for the index.html page
	http.HandleFunc("/", serveHTML)

	// Start the server
	log.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed: ", err)
	}
}
