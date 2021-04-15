package main

import "github.com/nyumbapoa/backend/app/helpers"

func main() {

/*	log.SetFlags(log.Lshortfile | log.LstdFlags)

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
	postgres.Migrate(database)*/

	helpers.Info{
		Name: "Walter Sawenja",
	}.SendEMail()
	//helpers.Sms("0746445198", "Jemo Karibu")
}
