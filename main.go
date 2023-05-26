package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	api := apiStat{hits: 0}
	mux.Handle("/", api.middlewareStats(http.FileServer(http.Dir("web"))))
	mux.HandleFunc("/healthz", healthHandler)
	mux.HandleFunc("/metrics", api.displayStats)
	http.ListenAndServe(":8080", mux)
}
