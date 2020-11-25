package house

import "github.com/go-chi/chi"

func (rs Resource) Router() *chi.Mux {
	r := chi.NewRouter()

	//
	r.Get("/get", rs.HandleGetHouses)
	r.Get("/get/{id}", rs.HandleGetHouse)
	r.Get("/get-image", rs.HandleGetHouseImage)
	r.Get("/get-type", rs.HandleGetHouseType)

	//
	r.Post("/create", rs.HandleCreateHouse)
	r.Post("/create-image", rs.HandleCreateHouseImage)
	r.Post("/create-type", rs.HandleCreateHouseType)

	//
	r.Put("/update", rs.HandleUpdateHouse)
	r.Put("/update/{id}/image", rs.HandleUpdateHouseImage)
	r.Put("/update/{id}/type", rs.HandleUpdateHouseType)

	//
	r.Delete("/delete/{id}", rs.HandleDeleteHouse)
	r.Delete("/delete/{id}/image", rs.HandleDeleteHouseType)
	r.Delete("/delete/{id}/type", rs.HandleDeleteHouseImage)

	return r
}
