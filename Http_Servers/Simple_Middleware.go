package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type logger struct {
	Inner http.Handler
}

func (l *logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	log.Printf("start %d\n", start)
 	l.Inner.ServeHTTP(w, r)
	log.Printf("finish %d\n",time.Since(start))
}

func helloo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello\n")
}

func main() {
	f := http.HandlerFunc(helloo)
	l := logger{Inner: f}
	http.ListenAndServe(":8000", &l)
}