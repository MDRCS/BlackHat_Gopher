package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	var r = mux.NewRouter()
	//Using Pattern checking with regex to validate the user name -> (bob1 -> 404 not found) ; (med -> Hello med)
	r.HandleFunc("/foo/{user:[a-z]+}",func(w http.ResponseWriter,r *http.Request){
		user := mux.Vars(r)["user"]
		fmt.Fprintf(w,"Hello %s \n",user)
	}).Methods("GET")

	http.ListenAndServe(":8000",r)

}
