package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/welcome", welcome)
	http.ListenAndServe(":8080", nil)
}

func welcome(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Welcome, "+ r.FormValue("name")))
}
