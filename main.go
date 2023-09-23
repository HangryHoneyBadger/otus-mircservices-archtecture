package main

import (
	"net/http"
)

var statusOk = []byte(`{"status": "OK"}`)

func hanldler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(statusOk)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func main() {
	http.HandleFunc("/health/", hanldler)
	http.ListenAndServe(":8000", nil)
}
