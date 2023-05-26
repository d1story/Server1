package main

import (
	"net/http"
)

type file struct {
	data []byte
}

func (htmlFile file) adminMetricHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("charset", "utf-8")
	w.Write(htmlFile.data)
}
