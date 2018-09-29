package main

import (
	"encoding/json"
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
)

type task struct {
	Name     string `json:"name,omitempty"`
	Priority string `json:"priority,omitempty"`
	Label    string `json:"label,omitempty"`
}

var list []task

func gettask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, v := range list {
		if v.Priority == params["priority"] {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	json.NewEncoder(w).Encode(&task{})
}

func getlist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// query param:
	r.ParseForm()

	pr := r.Form.Get("priority")
	la := r.Form.Get("label")
	filteredTasks := []task{}

	if pr != "" && la == "" {

		// filter
		for _, v := range list {
			if v.Priority == pr {
				filteredTasks = append(filteredTasks, v)
			}
		}
		json.NewEncoder(w).Encode(filteredTasks)
		return
	}
	if la != "" && pr == "" {

		// filter
		for _, v := range list {
			if v.Label == la {
				filteredTasks = append(filteredTasks, v)
			}
		}
		json.NewEncoder(w).Encode(filteredTasks)
		return
	}
	json.NewEncoder(w).Encode(list)
	return
}

func createtask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var listed []task
	var err error
	err = json.NewDecoder(r.Body).Decode(&listed)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Unable to decode given json:", err)
		return
	}
	list = append(list, listed...)
}

func deletetask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	la := r.Form.Get("label")
	for i := len(list) - 1; i >= 0; i = i - 1 {
		if list[i].Label == la {
			list = append(list[:i], list[i+1:]...)
		}
	}
}