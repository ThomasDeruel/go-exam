package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"io/ioutil"
)

type Task struct {
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type List struct {
	ID   string `json:"id"`
	Task string `json:"task"`
}

var tasks []Task

func list(rw http.ResponseWriter, _ *http.Request) {
	list := []List{}
	for id, el := range tasks {
		if !el.Done {
			list = append(list, List{strconv.Itoa(id), el.Description})
		}
	}
	rw.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(list)

	rw.Write(data)
}
func done(w http.ResponseWriter, r *http.Request) {
	fmt.Println("done")
}
func add(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Errorreadingbody:%v", err)
		http.Error(rw,"can'treadbody", http.StatusBadRequest,)
		return
	}
	description := string(body)

	fmt.Println(description)

	tasks = append(tasks,Task{description, false})
	fmt.Println(tasks)
}
func main() {

	http.HandleFunc("/", list)
	http.HandleFunc("/add", add)
	http.HandleFunc("/done", done)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
