package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", HomeHandler)

	// Posts collection
	r.GET("/posts", PostsIndexHandler)
	r.POST("/posts", PostsCreateHandler)

	// Posts singular
	r.GET("/posts/:id", PostShowHandler)
	r.PUT("/posts/:id", PostUpdateHandler)
	r.GET("/posts/:id/edit", PostEditHandler)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}

// HomeHandler handles home page request
func HomeHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Home")
}

// PostsIndexHandler handle get posts request
func PostsIndexHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "posts index")
}

// PostsCreateHandler handles post create request
func PostsCreateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "posts create")
}

// PostShowHandler handles single post show request
func PostShowHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	fmt.Fprintln(rw, "showing post", id)
}

func PostUpdateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "post update")
}

func PostEditHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "post edit")
}
