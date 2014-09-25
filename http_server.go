package main

import (
	"fmt"
	"net/http"
)

type handle struct {
	count int
}

func (o *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//	fmt.Println(o.count)
	fmt.Fprintf(w, "hello %d", o.count)
	o.count++
}
func main() {
	h := handle{}
	http.Handle("/count", &h)
	http.ListenAndServe(":8001", nil)
}

