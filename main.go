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
func done(rw http.ResponseWriter, r *http.Request) {
	switch r.Method{
	case http.MethodGet:
		fmt.Println("get")
		rw.WriteHeader(http.StatusOK)
		data, _ := json.Marshal(filterDone())
		rw.Write(data)
	case http.MethodPost:
		myid := r.FormValue("id")
		id,_ := strconv.Atoi(myid)
		if id > len(tasks)-1{
			rw.WriteHeader(http.StatusBadRequest)
		}
		for ind, _ := range tasks {
			if ind == id {
				tasks[ind] = Task{myid,true}
			}
		}
		rw.WriteHeader(http.StatusOK)
	default:
		rw.WriteHeader(http.StatusBadRequest)
	}
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
	tasks = append(tasks,Task{description, false})
}

func filterDone() []List{
	list := []List{}
	for id, el := range tasks {
		if el.Done {
			list = append(list, List{strconv.Itoa(id), el.Description})
		}
	}
	return list
}
func main() {

	http.HandleFunc("/", list)
	http.HandleFunc("/add", add)
	http.HandleFunc("/done", done)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
