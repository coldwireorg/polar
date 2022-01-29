package server

import (
	"net"
	"net/http"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/handler"
	"github.com/creachadair/jrpc2/jhttp"
	"github.com/rs/zerolog/log"
)

/**
 * Listener package embed a RPC server listening on the port 1312 by default.
 */

func Listen(port string) {
	f := handler.Map{
		"Push":     handler.New(Push),
		"Seed":     handler.New(Seed),
		"Register": handler.New(Register),
	}

	bridge := jhttp.NewBridge(f, &jhttp.BridgeOptions{
		Server: &jrpc2.ServerOptions{},
	})
	defer bridge.Close()

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Info().Err(err)
	}
	defer listener.Close()

	log.Info().Msg("Listening on port: " + port)

	http.Serve(listener, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bridge.ServeHTTP(w, r)
	}))
}
