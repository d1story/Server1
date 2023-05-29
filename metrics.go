package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

func (api *apiStat) adminMetricFormat(w http.ResponseWriter, r *http.Request) {
	file{data: formatAdminMetricPage(api.hits)}.adminMetricHandler(w, r)
}

func (api *apiStat) displayStats(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprint("hits:", api.hits)
	w.Write([]byte(response))
	w.Header()
}

func checkLength(w http.ResponseWriter, r *http.Request) {

}

func checkProfane(w http.ResponseWriter, r *http.Request) {
	User := r.FormValue("User")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()
	sbody := string(body)
	log.Println(sbody)
	if User == "kid" {
		split := strings.Split(sbody, " ")
		for i, str := range split {
			if str == "kerfuffle" || str == "sharbert" || str == "fornax" {
				split[i] = "****"
			}
		}
		sbody = strings.Join(split, " ")
	}
	w.Write([]byte(sbody))
}
