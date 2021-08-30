package rest_api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
	cors "github.com/rs/cors"
)

type Server struct {
	Router *mux.Router
}

func (serverConfiguration *Server) InitHTTPServer(databasepath string) {
	serverConfiguration.Router = mux.NewRouter()
	serverConfiguration.InitRouters()

	serverConfiguration.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		template, _ := route.GetPathTemplate()
		methods, _ := route.GetMethods()
		fmt.Printf("routes %s %s", methods, template)
		fmt.Println()
		return nil
	})
}

func (serverConfiguration *Server) Run(host string) {
	fmt.Println("Listening to:", host)
	handler := cors.Default().Handler(serverConfiguration.Router)
	log.Fatal(http.ListenAndServe(host, handler))
}

func (serverConfiguration *Server) InitRouters() {
}

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}

}

func JSONResponse(w http.ResponseWriter, statusCode int, response interface{}) {

	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}

}
