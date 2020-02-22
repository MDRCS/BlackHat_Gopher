package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s\n", r.URL.Query().Get("name"))
}

func main(){
	http.HandleFunc("/hello",hello)
	http.ListenAndServe(":8000",nil)
}



//Run the script and enter this command curl -i http://localhost:8000/hello?name=alice