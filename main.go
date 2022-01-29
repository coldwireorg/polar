package main

import (
	"flag"
	"os"
	"polar/database"
	"polar/models"
	"polar/server"
	"polar/utils"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	bind *string
	netw *string
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
	bind = flag.String("bind", "127.0.0.1", "address to bind polar on")
	port = flag.String("port", "1312", "port on which the node will listen")
	netw = flag.String("network", "mainnet", "network polar connect to, mainnet for production, testnet for experiment, etc")
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

	address := *bind + ":" + *port
	utils.InitNode(&address, *data, *seed, *netw)
}

func main() {
	log.Info().Msg("seeding server: " + *seed)
	log.Info().Msg("data path: " + *data)

	log.Info().Msg("Listing nodes:")
	nodes := models.Node{}.List()
	for _, node := range nodes {
		log.Info().Msg(node.Address + "/" + node.PublicKey)
	}

	server.Listen(*port)
}
