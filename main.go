package main

import (
"dipta_15116772_pert4/handler" //sesuaikan dengan nama folder (case sensitive)
"log"
"net/http"
)

func main() {
http.HandleFunc("/api/", handler.API)
//Ganti 2 digit akhir port dengan 2 digit akhir NPM anda
log.Println("localhost : 8029")
http.ListenAndServe(":8029", nil)
}