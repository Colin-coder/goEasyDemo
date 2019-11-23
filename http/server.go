package main

import (
	"io"
	"net/http"

	_ "net/http/pprof"
)

func helloServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world!!\n")
}

func main() {
	http.HandleFunc("/nktest/", helloServer)
	err := http.ListenAndServe(":8585", nil)
	if err != nil {
		panic(err)
	}
}
