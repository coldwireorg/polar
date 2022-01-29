package server

import (
	"context"
	"polar/models"
	"polar/structures"

	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
)

func Register(ctx context.Context, req structures.Request) error {
	var node models.Node
	err := mapstructure.Decode(req.Content, &node)
	if err != nil {
		return err
	}

	node.Add()

	log.Info().Msg("New node discovered: " + node.Address + "/" + node.PublicKey)

	return nil
}
