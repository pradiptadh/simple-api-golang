package handler

import (
"dipta_15116772_pert4/model" //sesuaikan dengan nama folder (case sensitive)
"encoding/json"
"io/ioutil"
"net/http"
"fmt"

)

func HandlerMahasiswaGet(w http.ResponseWriter, r *http.Request) {
	var data interface{}
	var err error

	npm := r.URL.Query()["npm"] //localhost:8029/api/mahasiswa?npm=12345678

	if len(npm) != 0 {
	data, err = model.GetMahasiswa(db, npm[0])
	} else {
	data, err = model.GetAllMahasiswa(db)
	}
	if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	jsonData, _ := json.Marshal(data)
	
w.Write(jsonData)
}
func HandlerMahasiswaPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body) 

//	defer r.Body.Close()
	if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return 
	}
	fmt.Println(body)
	var data model.Mahasiswa
	if err = json.Unmarshal([]byte(body), &data); err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
	}

	if err = data.Insert(db); err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	w.Write(jsonData)
}

func HandlerMahasiswaDelete(w http.ResponseWriter, r *http.Request) {
		npm := r.URL.Query()["npm"]
		if len(npm) != 0 {
		data := model.Mahasiswa{NPM: npm[0]}
		if err := data.Delete(db); err != nil {
			http.Error(w, "ID tidak ditemukan", http.StatusBadRequest)
			return
		}
			w.Write([]byte("Data telah dihapus"))
		} else {
			http.Error(w, "ID tidak ditemukan", http.StatusInternalServerError)
		}
}

func HandlerMahasiswaPut(w http.ResponseWriter, r *http.Request) {
	npm := r.URL.Query()["npm"]

	if len(npm) == 0 {
	http.Error(w, "ID tidak ditemukan", http.StatusBadRequest)
	return
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
	}

	jsonmap := make(map[string]interface{})
	err = json.Unmarshal(body, &jsonmap)
	if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
	}

	data := model.Mahasiswa{NPM: npm[0]}
	err = data.Update(db, jsonmap)
	if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
	}

	result, err := model.GetMahasiswa(db, npm[0])
	if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
	}

	w.Write(jsonData)
}