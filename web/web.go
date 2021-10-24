package main

import (
	"go-study/web/handler"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", handler.NewHttpRouter())
}
