package main

import (
	"fmt"
	"github.com/nyumbapoa/backend/app"
	"github.com/nyumbapoa/backend/app/registry"
	"github.com/nyumbapoa/backend/app/routing"
	"github.com/nyumbapoa/backend/app/storage/postgres"
	"github.com/nyumbapoa/backend/configs"
	"log"
	"os"
)

func main() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// read yaml configs file. Dont pass path to read
	// from default path
	yamlConfig := configs.ReadYaml("")
	config := app.GetConfig(*yamlConfig)

	database, err := postgres.NewDatabase(config)

	if err != nil {
		log.Printf("database err %s", err)
		os.Exit(1)
	}

	// run migrations; update tables
	postgres.Migrate(database)

	domain := registry.NewDomain(config, database)

	// create the fiber server.
	server := routing.Router(domain, config) // add endpoints

	// listen and serve
	port := fmt.Sprintf(":%v", 6700)
	log.Fatal(server.Listen(port))

	//helpers.Info{
	//	Name: "Walter Sawenja",
	//}.SendEMail()
	//helpers.Sms("0746445198", "Jemo Karibu")
}
