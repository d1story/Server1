package main

import (
	"bytes"
	"html/template"
	"log"
)

type metricData struct {
	Hits int
}

func formatAdminMetricPage(hits int) []byte {
	data := metricData{Hits: hits}
	tmpl, err := template.ParseFiles("./admin/metric/index.tmpl")
	if err != nil {
		log.Print("./admin/metric/index.tmpl file is missing", err)
		panic(err)
	}
	buf := bytes.NewBuffer(nil)
	err = tmpl.Execute(buf, data)
	if err != nil {
		log.Panic(err)
	}
	return buf.Bytes()
}
