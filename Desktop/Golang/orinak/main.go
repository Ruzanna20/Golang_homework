package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println(req.Response.Status)
			fmt.Println("REDIRECRT")
			return nil
		},
	}

	resp, err := http.DefaultClient.Get("https://go.dev/doc/tutorial/web-service-gin")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Println("Respone status:",resp.StatusCode)

	body,err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Body:",string(body))
}