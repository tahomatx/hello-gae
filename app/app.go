package main

import (
    "fmt"
    "net/http"
	// "github.com/satori/go.uuid"
	"github.com/tahomatx/hello-gae/hello"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    u1 := hello.Hi()+ "okokoko"
    fmt.Fprint(w , u1)
}
