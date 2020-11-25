package building

import "github.com/go-chi/chi"


//
func (rs Resource) Router() *chi.Mux {
	r := chi.NewRouter()

	// Get routes for querying records
	r.Get("/get", rs.HandleGetBuildings)
	r.Get("/get/{id}", rs.HandleGetBuilding)
	r.Get("/get/{id}/images", rs.HandleGetBuildingImages)
	r.Get("/get/{id}/types", rs.HandleDeleteBuildingType)

	// Post Routs for storing to the server
	r.Post("/create", rs.HandleCreateBuilding)
	r.Post("/create-image", rs.HandleCreateBuildingImage)
	r.Post("/create-type", rs.HandleCreateBuildingType)

	// Verifying a  building  is certified
	r.Post("/Verify/{id}", rs.HandleVerifyBuilding)

	// Update Routes for updating building record
	r.Put("/update/{id}", rs.HandleUpdateBuilding)
	r.Put("/update/{id}/image", rs.HandleUpdateBuildingImage)
	r.Put("/update/{id}/type", rs.HandleUpdateBuildingType)

	//
	r.Delete("/delete/{id}", rs.HandleDeleteBuilding)
	r.Delete("/delete/{id}/image", rs.HandleDeleteBuildingImage)
	r.Delete("/delete/{id}/type", rs.HandleDeleteBuildingType)

	return r

}
