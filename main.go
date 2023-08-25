package main

import "net/http"

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("<h2 style='color: red'>Welcome - My Argo CD testing project</h2>"))
	})

	http.ListenAndServe(":8080", nil)
}
