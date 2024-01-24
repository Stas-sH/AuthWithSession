package midlewares

import (
	"fmt"
	"log"
	"net/http"

	"Stas-sH/authWithSessions/internal/business/session"
)

func SirchSession(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("sessionId")
		if err != nil {
			log.Println("SirchSession - Cookie:", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		sessionInfo := session.InMemorySession.GetInfo(cookie.Value)
		if sessionInfo == "" {
			log.Println("session is`t found")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		fmt.Println("sessionInfo: UserName - ", sessionInfo)

		next(w, r)
	}
}
