package main

import "net/http"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := "hello from port 3000"
		w.Write([]byte(response))
	})

	println("listening at port 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}

}
