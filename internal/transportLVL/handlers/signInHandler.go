package handlers

import (
	"Stas-sH/authWithSessions/internal/business/session"
	signinusersdata "Stas-sH/authWithSessions/internal/business/signUPsignInUsersData"
	"Stas-sH/authWithSessions/internal/db"
	"Stas-sH/authWithSessions/pkg/hash"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		signIn(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func signIn(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("signIn - ReadAll:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var inp signinusersdata.SignInUserInput
	if err = json.Unmarshal(body, &inp); err != nil {
		log.Println("signIn - Unmarshal:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = inp.Validate(); err != nil {
		log.Println("signIn - Validate:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hasher, err := hash.NewSHA1Hasher()
	if err != nil {
		log.Println("signIn - NewSHA1Hasher:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	inp.Password, err = hasher.Hash(inp.Password)
	if err != nil {
		log.Println("signIn - Hash:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userInfo, err := db.GetUserFromDB(inp)
	if err != nil {
		log.Println("signIn - GetUserFromDB:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if userInfo.Id < 0 {
		log.Println("user with such credentials not found")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	/////////////////////////////////////////////////////////////
	session.InMemorySession = session.NewSession()

	sessionId := session.InMemorySession.Init(userInfo.UserName)
	cookie := &http.Cookie{
		Name:    "sessionId",
		Value:   sessionId,
		Expires: time.Now().Add(time.Minute * 5),

		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
}
