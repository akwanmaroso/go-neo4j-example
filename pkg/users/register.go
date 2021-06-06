package users

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type UserRegistration struct {
	User User `json:"user"`
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type UserHandler struct {
	Path           string
	UserRepository UserRepository
}

func (u *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	respBody, _ := ioutil.ReadAll(r.Body)
	userRegistrationRequestPayload := UserRegistration{}
	_ = json.Unmarshal(respBody, &userRegistrationRequestPayload)

	_ = u.UserRepository.RegisterUser(userRegistrationRequestPayload.User)

	w.WriteHeader(201)
	w.Header().Add("Content-Type", "application/json")
	userRegistrationResponsePayload := &User{
		Username: userRegistrationRequestPayload.User.Username,
		Email:    userRegistrationRequestPayload.User.Email,
	}
	userRegistrationResponse := UserRegistration{User: *userRegistrationResponsePayload}
	bytes, _ := json.Marshal(&userRegistrationResponse)
	_, _ = w.Write(bytes)
}
