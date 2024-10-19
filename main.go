package main

import "net/http"

func main() {
	// show the index.html file
	http.Handle("/", http.FileServer(http.Dir("./")))
	// show the index.html file
	http.Handle("/index.html", http.FileServer(http.Dir("./")))

	// show the index.html file
	http.Handle("/index", http.FileServer(http.Dir("./")))

	http.ListenAndServe(":8080", nil)

}
