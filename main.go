package main

import (
	"fmt"
	"log"
	"net/http"
)

type Task struct {
	Description string
	Done        bool
}

func list(w http.ResponseWriter, r *http.Request) {
	fmt.Println("default")
	return
}
func done(w http.ResponseWriter, r *http.Request) {
	fmt.Println("done")
	return
}
func add(w http.ResponseWriter, r *http.Request) {
	fmt.Println("add")
	return
}
func main() {
	//var task []Task
	
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/done", doneHandler)
	http.HandleFunc("/add", addHandler)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
