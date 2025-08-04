package main

import "net/http"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := "hello from port 8001"
		w.Write([]byte(response))
	})

	println("listening at port 8001")
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		panic(err)
	}

}
