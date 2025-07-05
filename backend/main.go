package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)



type GreetRequest struct {
	Name string `json:"name"`
}
type GreetResponse struct {
	Message string `json:"message"`
}
func GreetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") //handle CORS
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	var req GreetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	resp := GreetResponse{Message: fmt.Sprintf("Hello, %s!", req.Name)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}


func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/greet", GreetHandler)

	fmt.Println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}


