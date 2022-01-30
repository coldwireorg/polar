package rpc

import (
	"context"
	"polar/structures"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/jhttp"
	"github.com/rs/zerolog/log"
)

func Call(endpoint string, method string, req structures.Request) interface{} {
	ch := jhttp.NewChannel(endpoint, nil)
	client := jrpc2.NewClient(ch, nil)

	var res interface{}
	err := client.CallResult(context.Background(), method, req, &res)
	if err != nil {
		log.Err(err)
	}

	return res
}
