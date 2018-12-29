package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	port := 8000
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/status", Status)

	fmt.Println(fmt.Sprintf("Starting web server on port %d", port))

	sPort := fmt.Sprintf(":%d", port)
	log.Fatal(http.ListenAndServe(sPort, router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the test app!")
}

func Status(w http.ResponseWriter, r *http.Request) {
	if _, err := os.Stat("/tmp/testapp_status"); os.IsNotExist(err) {
		fmt.Fprint(w, "200")
	} else {
		if err != nil {
			log.Fatal(err)
		}
		bytes, err2 := ioutil.ReadFile("/tmp/testapp_status")
		if err2 != nil {
			log.Fatal(err2)
		}
		status := string(bytes)
		fmt.Fprint(w, status)
	}
}
