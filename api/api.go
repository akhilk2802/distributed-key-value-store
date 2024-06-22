package api

import (
	"encoding/json"
	"key-value-store/cluster"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	router  *mux.Router
	cluster *cluster.Cluster
}

func NewAPI(cluster *cluster.Cluster) *API {
	api := &API{
		router:  mux.NewRouter(),
		cluster: cluster,
	}
	api.setupRoutes()
	return api
}

func (api *API) setupRoutes() {
	api.router.HandleFunc("/set/{key}/{value}", api.setHandler).Methods("POST")
	api.router.HandleFunc("/get/{key}", api.getHandler).Methods("GET")
	api.router.HandleFunc("/delete/{key}", api.deleteHandler).Methods("DELETE")
}

func (api *API) setHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	value := vars["value"]

	api.cluster.Store.Set(key, value)
	api.cluster.BroadcastSet(key, value)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Key-Value pair set successfully"))

}

func (api *API) getHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, ok := api.cluster.Store.Get(key)
	if !ok {
		http.Error(w, "Key not found", http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(value)
}

func (api *API) deleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	api.cluster.Store.Delete(key)
}

func (api *API) Run(addr string) {
	log.Println("Here ")
	http.ListenAndServe(addr, api.router)
}
