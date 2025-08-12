package main

import (
	"encoding/json"
	"io"
	"net/http"
)

var quotes = []string{"I am Ruzanna.", "I am 20."}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte("Hello from server!"))
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Method not allowed!"))
}

func quotesHandler(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-type","application/json")
	switch r.Method {
		case "GET":
			Quotes,err := json.Marshal(quotes)
			if err == nil {
				w.Write(Quotes)
				return
			} 
		case "POST":
			body,err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Bad Request!"))
			}
			var newQuote string
			err = json.Unmarshal(body,&newQuote)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Bad Request!"))
				return			
			} 
			quotes = append(quotes, newQuote)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("New Quote added!"))
			return
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed!"))
		}
}

func main() {
	http.HandleFunc("/hello",helloHandler)
	http.HandleFunc("/quotes",quotesHandler)

	http.ListenAndServe(":8080", nil)
}