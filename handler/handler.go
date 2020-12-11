package handler

import (
	"dipta_15116772_pert4/model"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

var username, password, host, namaDB, defaultDB string
var db *sql.DB
var err error 

func init() {
	username = "root" //Misal : CPC29
	password = "" 
	host = "localhost:3306" //dbms.lepkom.f4.com
	namaDB = "gunadarma" //Nama DB misal : db_13116429
	defaultDB = "mysql"
}

func API(w http.ResponseWriter, r *http.Request) {
	db, err = model.Connect(username, password, host, namaDB)
	if err != nil {
	return
	}
	defer db.Close()
	w.Header().Set("Content-Type", "text-html; charset=utf-8; application/json")
	dataURL := strings.Split(fmt.Sprintf("%s", r.URL.Path), "/")
		switch dataURL[2] {
			case "mahasiswa":
			switch r.Method {
			case "GET":
			HandlerMahasiswaGet(w, r)
			case "POST":
			HandlerMahasiswaPost(w, r)
			case "PUT":
			HandlerMahasiswaPut(w, r)
			case "DELETE":
			HandlerMahasiswaDelete(w, r)
			default:
				w.Write([]byte("method tidak ditemukan"))
		}
	default:
	w.Write([]byte("request tidak ditemukan"))
	}
}
