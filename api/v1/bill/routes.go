package bill

import "github.com/go-chi/chi"

func (rs Resource) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/get/{id}", rs.HandleGetBill)
	r.Get("/all", rs.HandleGetBills)
	r.Post("/create", rs.HandleCreateBill)
	r.Put("/update", rs.HandleUpdateBill)
	return r

}
