package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

func main() {
	http.HandleFunc("/api/vali", vali)

	log.Fatal(http.ListenAndServe(":80", nil))
}

type CommonRequest struct {
	Aa int `validate:"required,gte=1"`
	Bb int `validate:"required,gte=1,lte=100"`
}

func vali(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	aa, _ := strconv.Atoi(q.Get("aa"))
	bb, _ := strconv.Atoi(q.Get("bb"))

	validate := validator.New()
	req := CommonRequest{Aa: aa, Bb: bb}
	err := validate.Struct(req)

	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
		return
	}
}
