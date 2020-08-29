package main

import (
	"github.com/hamologist/dice-roll/pkg/app"
	"net/http"
)

func main() {
	router := app.NewRouter()
	http.ListenAndServe(":3000", router)
}
