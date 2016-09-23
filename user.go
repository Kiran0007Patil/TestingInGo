package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = []User{}

func signUpHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		failureResponse(w, err.Error())
		return
	}

	type SignupRequestBody struct {
		Email    *string `json:"email"`
		Username *string `json:"username"`
		Password *string `json:"password"`
	}

	var userBody SignupRequestBody

	err = json.Unmarshal(body, &userBody)
	if err != nil {
		failureResponse(w, err.Error())
		return
	}

	if userBody.Email == nil || *userBody.Email == `` {
		failureResponse(w, `email parameter is not specified or empty`)
		return
	}

	if userBody.Username == nil || *userBody.Username == `` {
		failureResponse(w, `username parameter is not specified or empty`)
		return
	}

	if userBody.Password == nil || *userBody.Password == `` {
		failureResponse(w, `password parameter is not specified or empty`)
		return
	}

	if len(*userBody.Password) < 6 {
		failureResponse(w, `Number of characters in password should not be less than 6 `)
		return
	}

	err = checkUniquness(*userBody.Username)

	if err != nil {
		failureResponse(w, `Specified username is not unique! `)
		return
	}

	newUser := User{
		Id:       len(users) + 1,
		Username: *userBody.Username,
		Email:    *userBody.Email,
		Password: *userBody.Password,
	}

	users = append(users, newUser)

	type Output struct {
		Message string `json:"message"`
	}
	var output Output
	output.Message = "Successful User Creation"

	j, err := json.Marshal(output)
	if err != nil {
		failureResponse(w, `Internal Server Error! `)
		return
	}

	successResponse(w, j)

}

func checkUniquness(username string) error {
	for _, u := range users {
		if u.Username == username {
			return errors.New("Username is already used")
		}
	}

	return nil
}

func userListHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	users, err := json.Marshal(users)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(users)
}
