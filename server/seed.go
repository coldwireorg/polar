package server

import (
	"context"
	"polar/models"
	"polar/structures"

	"github.com/rs/zerolog/log"
)

func Seed(ctx context.Context, req structures.Request) ([]models.Node, error) {
	nodes := models.Node{}.List()

	log.Print(nodes)

	return nodes, nil
}
