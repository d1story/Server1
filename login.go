package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type user struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DB struct {
	Users []user `json:"Users"`
}

func getUsers() []user {
	db, err := os.Open("userDataBase.json")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	data, err := ioutil.ReadAll(db)
	if err != nil {
		log.Println(err)
	}
	var v DB
	json.Unmarshal(data, &v)
	return v.Users
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	u := getUsers()
	log.Println(u)
	iemail := r.FormValue("email")
	ipassword := r.FormValue("password")
	fmt.Println(iemail)
	fmt.Println(ipassword)
	found := false
	fid := -1
	for _, r := range u {
		fmt.Println(r.Email, r.Password)
		if r.Email == iemail && r.Password == ipassword {
			fmt.Println("TRUEEEE")
			found = true
			fid = r.Id
			break
		}
	}
	if found {
		fmt.Println("in")
		w.Write([]byte(fmt.Sprintf("%d", fid)))
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
