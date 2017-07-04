package main

import "net/http"

func main() {
	http.HandleFunc("/", hoge)
	http.ListenAndServe(":8080", nil)
}

func hoge(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hogehogehogehoge!\n"))
}
