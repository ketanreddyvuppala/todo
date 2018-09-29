package main

import (
	"fmt"
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	list = append(list, task{Name: "Task for the day", Priority: "High", Label: "personal"})
	list = append(list, task{Name: "Research phase", Priority: "Low", Label: "work"})
	list = append(list, task{Name: "Weekly team meeting-skype", Priority: "High", Label: "group"})
	list = append(list, task{Name: "First Draft", Priority: "Medium", Label: "work"})
	list = append(list, task{Name: "Grocery", Priority: "Medium", Label: "Personal"})
	list = append(list, task{Name: "Weekly team meeting-skype", Priority: "High", Label: "group"})
	list = append(list, task{Name: "Documentation", Priority: "Low", Label: "work"})
	list = append(list, task{Name: "Second Draft", Priority: "Low", Label: "work"})
	list = append(list, task{Name: "Friday night Football game", Priority: "High", Label: "Personal"})
	list = append(list, task{Name: "Final Presentation", Priority: "High", Label: "work"})
	router.HandleFunc("/status",StatusHandler).Methods("GET")
	router.HandleFunc("/task", getlist).Methods("GET")
	router.HandleFunc("/task", createtask).Methods("POST")
	router.HandleFunc("/task", deletetask).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Status: OK")
}
