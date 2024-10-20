package main

import (
	"html/template"
	"log"
	"net/http"
)

// Welcome struct to hold the welcome message
type Welcome struct {
	Message string
}

// Handler function to render the HTML template
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML template
	tmpl, err := template.ParseFiles("welcome.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	// Data to be passed to the template
	welcomeData := Welcome{Message: "Welcome to the Golang Web App with Bootstrap!"}

	// Execute the template with the welcomeData
	err = tmpl.Execute(w, welcomeData)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}

func main() {
	// Set up a handler for the root URL
	http.HandleFunc("/", welcomeHandler)

	// Start the web server
	log.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
