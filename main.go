package main

import (
	"encoding/json"
	"net/http"

	"github.com/SidVermaS/Ethereum-Consensus/pkg/helpers"
)

func GreetHandler(w http.ResponseWriter, r *http.Request)	{
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"data":"hello "})
}

func main() {
	helpers.InitializeAll()
	http.HandleFunc("/",GreetHandler)
	err:=http.ListenAndServe(":8080",nil)
	panic(err)
}
