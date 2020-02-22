package main

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni" //negroni is middleware framework that allows you to extend to much of operations linked to middleware
	"net/http"
)


//A middleware is a wrapper that is used when you want to execute some checks or operations on the response of the requests, for example authetification etc ..

func main() {
	r := mux.NewRouter()
	n := negroni.Classic()
	n.UseHandler(r)
	http.ListenAndServe(":8000",n)
}