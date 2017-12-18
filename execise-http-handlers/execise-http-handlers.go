// http route

package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

func (s String) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, string(s))
}

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s Struct) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, fmt.Sprintf("{\"ok\": \"%v\"}", s))
}

func main() {

	http.Handle("/string", String("I'm Dash, Learning Golang~"))
	http.Handle("/struct", Struct{"Hello", ":", "Dash"})

	err := http.ListenAndServe("127.0.0.1:8888", nil)
	if err != nil {
		log.Fatal(err)
	}
}
