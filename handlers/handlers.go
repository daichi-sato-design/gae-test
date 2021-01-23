package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/daichi-sato-design/gae-test/conf"
	"github.com/daichi-sato-design/gae-test/utils"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Response return json
type Response struct {
    Port  string `json:"port"`
		Logfile string `json:"logfile"`
		URL string `json:"api_url"`
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
	utils.Respond(w,http.StatusOK,Response{
		Port: conf.Config.Port, 
		Logfile: conf.Config.LogFile, 
		URL: conf.Config.APIURL, 
		Message:"Hi! this is GAE test app by golang.",
	})
}