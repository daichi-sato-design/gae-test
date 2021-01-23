package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/daichi-sato-desin/gae-test/conf"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Response return json
type Response struct {
    Status  string `json:"status"`
    Message string `json:"message"`
}

// Handlers this app handler functions
func Handlers(){
	r := mux.NewRouter()

	r.HandleFunc("/", handle).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = conf.Config.Port
	}
	handler := cors.AllowAll().Handler(r)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

func handle(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(Response{Status: "ok", Message: "Hello world."})
}