package main

import (
	"flag"
	"os"
	"polar/client"
	"polar/database"
	"polar/models"
	"polar/server"
	"polar/structures"
	"polar/utils"

	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	port *string
	seed *string
	data *string
)

func init() {
	// Print banner, because it's nice :)
	utils.Banner("alpha-01")

	// Setup logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Declare CLI flags
	port = flag.String("port", "1312", "port on which the node will listen")
	seed = flag.String("seed", "https://polar.coldnet.org:1312", "address to an instance of polar to download the network and discorver new nodes")
	data = flag.String("data", "/opt/polar/", "path to the folder where the local database will be stored")
	flag.Parse()

	// Check if storage dir exist and if signing keys are there
	utils.CheckPath(*data)
	utils.ChecKeys(*data)

	// Init database
	database.Connect(*data)
	database.DB.AutoMigrate(&models.Node{})
	database.DB.AutoMigrate(&models.Object{})

	if (models.Node{}.Number() == 0) {
		models.Node{
			Address:   "0.0.0.0",
			PublicKey: utils.GetPbKey(*data),
			Status:    1,
			Storage:   400000,
		}.Add()

		var nodeList []models.Node
		r := client.Call(*seed, "Seed", structures.Request{})
		mapstructure.Decode(r, nodeList)

		for _, node := range nodeList {
			log.Print(node.Address)
			node.Add()
		}
	}
}

func main() {
	log.Info().Msg("seeding server: " + *seed)
	log.Info().Msg("data path: " + *data)

	log.Print(models.Node{}.List())

	server.Listen(*port)
}
