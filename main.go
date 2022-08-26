package main

import (
	"net/http"

	"github.com/Gimmingyu/GoWebPractice/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler)
}
