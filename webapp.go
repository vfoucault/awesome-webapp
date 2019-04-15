package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var (
	buildRevision string
	buildTag      string
)

func main() {
	app := new(App).Create()
	app.AddRoutes()
	app.Run("0.0.0.0", 8000)
}


type baseResponse struct {
	//Data map[interface{}]interface{} `json:"data"`
	Data map[string]interface{} `json:"data"`
}


func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	data := make(map[string]interface{})
	response := &baseResponse{}
	data["page"] = "home page"
	response.Data = data
	payload, _ := json.MarshalIndent(response, "", "    ")
	fmt.Fprintf(w, string(payload))
}

type App struct {
	Router *mux.Router
}

func (a *App) Create() *App {
	a.Router = mux.NewRouter()
	return a
}

func (a *App) Run(host string, port int) {
	srv := &http.Server{
		Handler:      a.Router,
		Addr:         fmt.Sprintf("%v:%v", host, port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func (a *App) AddRoutes() {
	a.Router.HandleFunc("/", HomeHandler)
	http.Handle("/", a.Router)
}