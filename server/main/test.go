package main

import (
	"fmt"
	"net/http"
	"strings"
)

func handleShell(w http.ResponseWriter, r *http.Request) {
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		ips := strings.Split(forwarded, ", ")
		fmt.Fprintln(w, ips[0])
	}
	fmt.Fprintln(w, strings.Split(r.RemoteAddr, ":")[0])
}
func main() {
	http.HandleFunc("/", handleShell)
	http.ListenAndServe(":8080", nil)
}
