package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
)

type metricData struct {
	hits int
}

func formatAdminMetricPage(hits int) []byte {
	data := metricData{hits: hits}
	tmpl, err := template.ParseFiles("./admin/metric/index.tmpl")
	if err != nil {
		log.Print("./admin/metric/index.tmpl file is missing", err)
		panic(err)
	}
	buf := bytes.NewBuffer(nil)
	a, _ := os.Open("./logFile.txt")
	a.Write([]byte("hihi"))
	tmpl.Execute(a, data)
	a.Close()
	tmpl.Execute(buf, data)
	fmt.Print(buf.Bytes())
	return buf.Bytes()
}
