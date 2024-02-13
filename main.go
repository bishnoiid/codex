package main

import (
	"log"
	"net/http"
)

// Home Handler func writing bite slice "hello from backend" as response body
func home(w http.ResponseWriter, r *http.Request) {
	//log.Println(r)
	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}
	w.Write([]byte("Hello from backend"))
}

func viewSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display snippet"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow","POST")
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}
	w.Write([]byte("Write snippet"))
}


func main () {
	//fmt.Println("Hello Sid")
	// Use the http.NewServeMux() func to initialize a new servemux and register it to home directory "/"
	mux := http.NewServeMux()
	mux.HandleFunc("/",home)
	mux.HandleFunc("/snippet/view",viewSnippet)
	mux.HandleFunc("/snippet/create",createSnippet)

	// start a new server using http.ListenAndServe() we pass two parameters TCP connection address to listen to and the servemux we just created if theres and error we use log.Fatal() to log the error 


	log.Println("starting server on 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
