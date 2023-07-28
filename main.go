package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
)

func customFileServer(fs http.FileSystem) http.Handler {
	fileServer := http.FileServer(fs)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fs.Open(path.Clean(r.URL.Path)) // Do not allow path traversals.
		if os.IsNotExist(err) {
			http.ServeFile(w, r, "./static/index.html")
			return
		}
		fileServer.ServeHTTP(w, r)
	})
}

func main() {
	fs := customFileServer(http.Dir("./static"))
	http.Handle("/", fs)
	fmt.Print("Listening on 4000")
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
