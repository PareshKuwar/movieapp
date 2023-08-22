package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"
)

func gadar2(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type","text/html")

	htmlcontent:= `<!DOCTYPE html>
		<html>
		<head>
			<title>Simple Go API</title>
			</head></html>`
	fmt.Fprintf(w, htmlcontent)		
}

func main(){
	r:= mux.NewRouter()
	r.HandleFunc("/",gadar2).Methods("GET")
	fmt.Println("serever startedd")
	log.Fatal(http.ListenAndServe(":7000",r))

}