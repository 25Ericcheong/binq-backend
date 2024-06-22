package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting Binq backend server")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "Welcome to the index")

		if err != nil {
			return
		}
	})

	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		fmt.Println("Error occurred while trying to run server")
		fmt.Println(err.Error())
	}

	fmt.Println("Successful start up. Waiting for request... ")
}
