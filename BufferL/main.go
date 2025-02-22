package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	b := bytes.NewBuffer(nil)

	var name string
	fmt.Scanln(&name)

	b.Write([]byte("Hello "))
	b.WriteString(name)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.Copy(w, b)
	})
	log.Fatalln(http.ListenAndServe(":8080", nil))

	fmt.Println(b.String())
}
