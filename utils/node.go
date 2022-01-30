package utils

import (
	"math/rand"
	"polar/models"
	"polar/rpc"
	"polar/structures"

	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
)

func randomSeed() models.Node {
	nodes := models.Node{}.List()
	randomIndex := rand.Intn(len(nodes))

	return nodes[randomIndex]
}

func seedServers(seed string) {
	r := rpc.Call(seed, "Seed", structures.Request{})

	n := []models.Node{}
	err := mapstructure.Decode(r, &n)
	if err != nil {
		log.Err(err)
	}

	for _, node := range n {
		n := models.Node{
			PublicKey: node.PublicKey,
		}.FindWithPbKey()

		if n.PublicKey == "" {
			log.Info().Msg("New node discovered: " + node.Address + "/" + node.PublicKey)

			rpc.Call("http://"+node.Address, "Register", structures.Request{
				Content: models.Node{}.List()[0],
			})
			node.Add()
		}
	}
}

func InitNode(address *string, path string, seed string, netw string) {
	// Get node's free space
	freeSpace, err := GetFreeSpace()
	if err != nil {
		log.Fatal().Err(err)
	}

	// if the node list is empty
	nNodes := models.Node{}.Number()
	if nNodes <= 1 {
		if nNodes == 0 {
			models.Node{
				Address:   *address,
				Network:   netw,
				PublicKey: GetPbKey(path),
				Status:    1,
				Storage:   freeSpace,
			}.Add()
		}

		seedServers(seed)
	} else {
		seedServers(randomSeed().Address)
	}
}
