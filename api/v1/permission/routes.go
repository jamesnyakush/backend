package permission

import "github.com/go-chi/chi"

func (rs Resource) Router() *chi.Mux  {
	r := chi.NewRouter()

	return r
}