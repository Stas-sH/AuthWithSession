package handlers

import "net/http"

func WaagHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("WAAAAGH"))
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
