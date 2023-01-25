package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	fmt.Print("Listening on 4000")
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
