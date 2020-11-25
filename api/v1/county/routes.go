package county

import "github.com/go-chi/chi"

func (rs Resource) Router() *chi.Mux {

	r := chi.NewRouter()

	// fetching all counties
	r.Get("/get", rs.HandleGetCounties)

	// Creating a new county
	r.Post("/create", rs.HandleCreateCounty)

	// updating a county
	r.Put("/update/{id}", rs.HandleUpdateCounty)

	// Deleting a county
	r.Delete("/delete/{id}", rs.HandleDeleteCounty)

	return r
}
