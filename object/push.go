package object

import (
	"context"
	"polar/models"
	"polar/structures"

	"github.com/mitchellh/mapstructure"
)

func Push(ctx context.Context, req structures.Request) error {
	var obj models.Object
	err := mapstructure.Decode(req.Content, &obj)
	if err != nil {
		return err
	}

	//TODO: push

	return nil
}
