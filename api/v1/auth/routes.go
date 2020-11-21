package auth

import "github.com/go-chi/chi"

func (rs Resource) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/login", rs.HandleLogin)
	r.Post("/signup", rs.HandleSignUp)
	r.Post("/reset", rs.HandleResetPassword)
	return r

}
