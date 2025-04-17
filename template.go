package main

import(
	"fmt"
	"html/template"
	"net/http"
)

/*	template.ParseFiles = digunakan untuk membaca dan memparsing file tersebut untuk mengecek apakah ada template tag {{.v}}
	Execute = membuat hasil parsing template ditampilkan di client
 */

func templ() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Name" : "John Wick",
			"Message" : "Good Luck",
		}

		p,err := template.ParseFiles("template.html")

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		p.Execute(w, data)
	})

	fmt.Println("Starting web server at http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		fmt.Errorf(err.Error())
	}
}