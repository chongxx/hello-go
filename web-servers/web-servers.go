// Web 服务器

package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Dash!")
}

func main() {
	var h Hello
	err := http.ListenAndServe("127.0.0.1:8888", h)
	if err != nil {
		log.Fatal(err)
	}
}
