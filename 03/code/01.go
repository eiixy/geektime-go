package code

import (
	"fmt"
	"log"
	"net/http"
)

func RunServers() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("request: %s\r\n", r.URL.Path)
		fmt.Fprintln(w, "hello, Gopher"+r.URL.Query().Get("a"))
	})
	go serveApp()
	go serveDebug()
}

func serveApp() {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func serveDebug() {
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
