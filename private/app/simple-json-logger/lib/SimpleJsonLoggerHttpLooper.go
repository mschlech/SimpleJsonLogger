package lib

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func StartJsonLogger() error {
	// call the http handler

	log.Print("starting the simple json logger")
	restApiMap()

	return nil
}

func restApiMap() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/getJson", encodePostRequestFromSource).Methods("GET")
	router.HandleFunc("/health", healtCheck).Methods("HEAD")
	log.Fatal(http.ListenAndServe(":5000", router))
}

func encodePostRequestFromSource(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewDecoder(r.Body)
	// response Body
	responseBody, err := ioutil.ReadAll(r.Body)

	data := &map[string]interface{}{}
	json.Unmarshal(responseBody, data)
	if err != nil {
		panic(err)
	}

	fmt.Print(data)
}

func healtCheck(responseWriter http.ResponseWriter, r *http.Request) {

	fmt.Printf("healthCheck=", responseWriter, r)
	return
}
