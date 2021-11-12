package main

import (
	"encoding/json"
	"html/template"
	"log"
	"myapp/rps"
	"net/http"
	"strconv"
)

//return hoePage
func homePage(w http.ResponseWriter, r *http.Request) {
	// html := `<strong>Hello, world </strong>`
	// w.Header().Set("Content-Type", "text/html") //le avisamos q va a recibir un html
	// fmt.Fprint(w, html)

	renderTemplate(w, "index.html")

}

func playRound(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice)
	//convertir el type struct en type l json
	out, err := json.MarshalIndent(result, "", "    ")

	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}

func main() {
	http.HandleFunc("/play", playRound)
	//start a web server
	http.HandleFunc("/", homePage)

	log.Println("starting web server ")
	http.ListenAndServe(":8000", nil)
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
