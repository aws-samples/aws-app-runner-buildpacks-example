package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const HTML = `<!DOCTYPE html>
<html>
  <head>
    <title>Golang Sample</title>
    <link rel='stylesheet' href='/stylesheets/style.css'/>
  </head>
  <body>
    <h1>Golang on AWS App Runner with Cloud Native Buildpacks</h1>
    <p>This application is running as a container on AWS App Runner and was built using Cloud Native Buildpacks!</p>
  </body>
</html>
`

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, HTML)
	})

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
