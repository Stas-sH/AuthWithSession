package main

import (
	"Stas-sH/authWithSessions/internal/transportLVL/handlers"
	"Stas-sH/authWithSessions/internal/transportLVL/midlewares"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users/signUp", handlers.SignUpHandler)
	http.HandleFunc("/users/signIn", handlers.SignInHandler)

	http.HandleFunc("/users/waagh", midlewares.SirchSession(handlers.WaagHandler))

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
