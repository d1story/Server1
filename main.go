package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

var logFile *os.File

func main() {
	var err error
	logFile, err = os.Open("./logFile.txt")
	if err != nil {
		log.Print("admin/metric/index.temp file is missing", err)
		panic(err)
	}
	log.SetOutput(logFile)
	log.Print("Main")

	r := chi.NewRouter()
	Api := apiStat{hits: 0}
	r.Mount("/", Api.middlewareStats(http.FileServer(http.Dir("web"))))

	apiRouter := setUpApiRouter(&Api)
	r.Mount("/api", apiRouter)

	adminRouter := setUpAdminRouter(&Api)
	r.Mount("/admin", adminRouter)

	http.ListenAndServe(":8080", r)
}
