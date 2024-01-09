package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {

	http.HandleFunc("/greetings", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var request Person
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
			http.Error(w, "Error: ", http.StatusBadRequest)
			return
		}

		response := fmt.Sprintf("Hello %s %s", request.FirstName, request.LastName)
		w.Write([]byte(response))

	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
