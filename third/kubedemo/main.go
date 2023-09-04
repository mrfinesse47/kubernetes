package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main(){
	started := time.Now()

	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		duration := time.Now().Sub(started)

		if duration.Seconds() > 10 {
			log.Println("timeout triggered")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`i am timed out`))
			// this gets the app to fail after 10 sec
		}else{
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`hello gopher`))
		}
		
		// fmt.Fprintf(w, "Hello Gopher")
	}) 
	http.HandleFunc("/aligator",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello Mr Aligator")
	}) 
	log.Fatal(http.ListenAndServe(":8080",nil))
}