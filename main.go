package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	// ListenAndServe ホストを起動する
	err := http.ListenAndServe(":18080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world %s", r.URL.Path[1:])
	}),
	)

	if err != nil {
		fmt.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}

}
