package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
    var err error
    if _, err = template.New("index").Parse(Index); err != nil {
        log.Printf("%s\n", err)
    }
    if err != nil {
        log.Fatalf("%s\n", err)
    }

    http.HandleFunc("/", getHandler)

    log.Printf("Listening on port 8080\n")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

type indexparameters struct {
    Configurations []string
}
func getHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        t, err := template.New("index").Parse(Index);
        if (err != nil) {
            fmt.Fprintf(w, "template error %s\n", err)
        }
        idxParams := indexparameters{ Configurations: []string{"A","B"}}
        t.Execute(w, idxParams)
    } else if r.Method == "PUT" {
        log.Printf("Got put\n")
    }
}



var Index = `<!DOCTYPE html>
<html lang="en">
    <head>
        <title>SCIM Client Configurator</title>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <script src="https://unpkg.com/htmx.org@1.9.4" integrity="sha384-zUfuhFKKZCbHTY6aRR46gxiqszMk5tcHjsVFxnUo8VMus4kHGVdIYVbOYYNlKmHV" crossorigin="anonymous"></script>
        <link href="css/style.css" rel="stylesheet">
    </head>
    <body>
        <div class="container">
            <h1>Hello World</h1>
            <select name="make" hx-get="/models" hx-target="#models" hx-indicator=".htmx-indicator">
                {{range .Configurations}}
                <option value="{{.}}">{{.}}</option>
                {{end}}
            </select>
        </div>
    </body>
</html>`
