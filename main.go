package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Task struct {
	Description string
	Done        bool
}
type List struct {
	ID   string `json:"id"`
	Task string `json:"task"`
}

var tasks []Task

func list(rw http.ResponseWriter, _ *http.Request) {
	list := []List{}
	tasks = []Task{
		{"Faire l'examen dans les temps", false},
		{"Se poser dans son lit", true},
		{"Faire les cours", false},
	}
	for id, el := range tasks {
		if !el.Done {
			list = append(list, List{strconv.Itoa(id), el.Description})
		}
	}
	rw.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(list)

	rw.Write(data)
	return
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
