package main

import (
	"fmt"
	"net/http"
)

/*	http.ResponseWriter = mengirim response ke client
	http.Request = menyimpan data permintaan dari client
	http.HandleFunc(route, action) = mendaftarkan handle function tertentu terhadap sebuah route tertentu
	http.ListenAndServe = menjalankan server pada alamat port yang sudah ditentukan
 */

func index(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path != "/index" {
		http.Error(w,"404 Not Found", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Hello World")
}

func net() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, "404 Not Found", http.StatusNotFound)
			return
		}

		fmt.Fprintln(w, "Halo!")
	})
	http.HandleFunc("/index", index)

	fmt.Println("Starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}