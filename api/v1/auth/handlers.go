package auth

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"net/http"
	"strings"
)

type signupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (body *signupRequest) Bind(r *http.Request) error {
	body.Name = strings.TrimSpace(body.Name)
	body.Email = strings.TrimSpace(body.Email)
	body.Password = strings.TrimSpace(body.Password)

	return validation.ValidateStruct(body, validation.Field(&body.Name, validation.Length(3, 32), is.ASCII),
		validation.Field(&body.Email, validation.Required, is.Email),
		validation.Field(&body.Password, validation.Required, validation.Length(8, 32), is.Alphanumeric))
}

func (body *loginRequest) Bind(r *http.Request) error {
	body.Email = strings.TrimSpace(body.Email)
	body.Password = strings.TrimSpace(body.Password)

	return validation.ValidateStruct(body, validation.Field(&body.Email, validation.Required, is.Email), validation.Field(&body.Password, validation.Length(8, 32), is.Alphanumeric))
}

func (rs Resource) HandleSignUp(rw http.ResponseWriter, r *http.Request) {
	body := signupRequest{}

	fmt.Println(body)
	/*
		if err := render.Bind(r, &body); err != nil {
			log.Println(err)
			return
		}

		usr, err := rs.Store.User().New(body.Name, body.Email, body.Password)

		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(usr)


		render.Status(r, http.StatusCreated)
		render.Respond(rw, r, http.NoBody)
	*/
}

func (rs Resource) HandleLogin(rw http.ResponseWriter, r *http.Request) {
	body := loginRequest{}

	fmt.Println(body)
	/*	if err := render.Bind(r, &body); err != nil {
		log.Println(err)
		//render.Render(rw,r,http.StatusBadRequest)
		return
	}*/
}

func (rs Resource) HandleResetPassword(w http.ResponseWriter, r *http.Request) {

}
