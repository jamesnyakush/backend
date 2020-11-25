package feedback

import "github.com/go-chi/chi"

func (rs Resource) Router() *chi.Mux {
	r := chi.NewRouter()

	// Getting all Feedback
	r.Get("/get", rs.HandleGetFeedback)

	// Create new Feedback
	r.Post("/create", rs.HandleCreateFeedback)

	// Edit Feedback
	r.Put("/update/{id]", rs.HandleUpdateFeedback)

	// Delete Feedback
	r.Delete("/delete/{id]", rs.HandleDeleteFeedback)

	return r
}
