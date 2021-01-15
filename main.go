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
var tasks []Task

func list(rw http.ResponseWriter, _ *http.Request) {
	fmt.Println("list")
}
func done(w http.ResponseWriter, r *http.Request) {
	fmt.Println("done")
}
func add(w http.ResponseWriter, r *http.Request) {
	fmt.Println("add")
}
func main() {
	
	http.HandleFunc("/", list)
	http.HandleFunc("/done", done)
	http.HandleFunc("/add", add)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
