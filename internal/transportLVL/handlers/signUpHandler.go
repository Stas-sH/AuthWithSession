package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	signupusersdata "Stas-sH/authWithSessions/internal/business/signUPsignInUsersData"
	"Stas-sH/authWithSessions/internal/db"
	"Stas-sH/authWithSessions/pkg/hash"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		signUp(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func signUp(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("signUp - ReadAll:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var inp signupusersdata.SignUpUserInput

	//fmt.Println(string(body))

	if err = json.Unmarshal(body, &inp); err != nil {
		log.Println("signUp - Unmarshal:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = inp.Validate(); err != nil {
		log.Println("signUp - Validate:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hasher, err := hash.NewSHA1Hasher()
	if err != nil {
		log.Println("signUp - NewSHA1Hasher:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	inp.Password, err = hasher.Hash(inp.Password)
	if err != nil {
		log.Println("signUp - Hash:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var isUser signupusersdata.SignInUserInput = signupusersdata.SignInUserInput{
		Mail:     inp.Mail,
		Password: inp.Password,
	}
	user, err := db.GetUserFromDB(isUser)
	if user.Id != -1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("there is already such a user"))
		return
	}

	if err = db.CreateUserInDB(inp); err != nil {
		log.Println("signUp - CreateUserInDB:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
