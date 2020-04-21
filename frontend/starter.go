package main

import (
	"log"
	"net/http"
	"spider/frontend/controller"
)

func main() {
	http.Handle("/search", controller.CreateSearchResultHandler("frontend/view/template.html"))
	log.Fatal(http.ListenAndServe(":8888", nil))
}
