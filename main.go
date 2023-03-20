package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
)

func bodyClose(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		log.Println(err)
	}
}

//func renderJSON(w http.ResponseWriter, v interface{}) {
//	js, err := json.Marshal(v)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	w.Header().Set("Content-Type", "application/json")
//	w.Write(js)
//}

func mainPage(writer http.ResponseWriter, request *http.Request) {
	defer bodyClose(request.Body)
	//tmpl := template.New("/templates/main.html")
	//err := tmpl.Execute(writer, data)
	//if err != nil {
	//	http.Error(writer, err.Error(), http.StatusInternalServerError)
	//}
	http.ServeFile(writer, request, "./templates/main.html")
}

func gamesList(writer http.ResponseWriter, request *http.Request) {
	defer bodyClose(request.Body)
	http.ServeFile(writer, request, "./templates/games.html")
}

func main() {
	router := mux.NewRouter()
	router.StrictSlash(true)
	//uri := "localhost:27017"
	//if len(os.Args) != 1 && (os.Args[1] == "" || os.Args[1] == "_") {
	//	uri = os.Args[1]
	//}
	target := "localhost:7777"
	if len(os.Args) > 2 && os.Args[2] != "" {
		target = os.Args[2]
	}

	router.HandleFunc("/", mainPage).Methods("GET")
	router.HandleFunc("/games", gamesList).Methods("GET")
	http.Handle("/", router)

	err := http.ListenAndServe(target, nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		log.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
