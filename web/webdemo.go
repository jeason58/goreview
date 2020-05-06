package web

import (
	"fmt"
	"log"
	"net/http"
)

func StartWebServer()  {
	http.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<html><head><title>web demo</title></head><body><h1>url is: <link><i>%s</i><link></h1></body></html>", r.RequestURI)
	})

	log.Printf("server started on port: %s", "80")
	log.Fatal(http.ListenAndServe(":80", nil))
}
