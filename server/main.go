package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter,r *http.Request){
	if r.Method == "GET" && r.URL.Path == "/hello" {
		fmt.Fprintln(w,"Barev.Es Ruzannan em.")
	}
}

func byeHandler(w http.ResponseWriter,r *http.Request) {
	if r.Method == "GET" && r.URL.Path == "/"  {
		fmt.Fprintln(w,"Hajoxutyun!!!")
	}
}
func main() {
	http.HandleFunc("/hello",helloHandler)
	http.HandleFunc("/",byeHandler)

	fmt.Println("Servery ashxatuum e http://localhost:8080:")
	http.ListenAndServe(":8080",nil)
}