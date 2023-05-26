package main

import (
	"fmt"
	"net/http"
)

type apiStat struct {
	hits int
}

func (api *apiStat) middlewareStats(ret http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			api.hits++
			ret.ServeHTTP(w, r)
		})
}

func (api *apiStat) displayStats(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprint("hits:", api.hits)
	w.Write([]byte(response))
}
