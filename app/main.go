package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/npex42/timer/components"
)

var startTime time.Time

func main() {

	startTime = time.Now()

	header := components.Header()

	http.Handle("/", templ.Handler(header))
	
	http.HandleFunc("/timer", timer);

	http.HandleFunc("/htmx.min.js", serveHTMX)

	fmt.Println("Server Running On :8080")
	http.ListenAndServe(":8080", nil)
}


func serveHTMX(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./static/htmx.min.js")
}

func timer(w http.ResponseWriter, req *http.Request) {
	components.Timer(startTime.Add(time.Minute * 10), time.Now(), "10s").Render(context.Background(), w)
}

