package discovery

import (
	"context"
	"polar/models"
	"polar/structures"
)

func Seed(ctx context.Context, req structures.Request) ([]models.Node, error) {
	nodes := models.Node{}.List()
	return nodes, nil
}
