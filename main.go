package main

import (
	"net/http"
)

func hanldler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func main() {
	http.HandleFunc("/health/", hanldler)
	http.ListenAndServe(":8000", nil)
}
