package client

import (
	"context"
	"polar/structures"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/jhttp"
)

func Call(endpoint string, method string, req structures.Request) interface{} {
	ch := jhttp.NewChannel(endpoint, nil)
	client := jrpc2.NewClient(ch, nil)

	var res interface{}
	client.CallResult(context.Background(), method, req, &res)

	return res
}
