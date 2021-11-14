package main

import (
	"github.com/Oppodelldog/cmdsrv/rpcserver"
	"github.com/Oppodelldog/jsonrpc-remote-cli/usecases"
	"github.com/swaggest/openapi-go/openapi3"
	"log"
)

func main() {
	const (
		title       = "JSON-RPC Example"
		version     = "v0.1.0"
		description = "This app showcases a trivial JSON-RPC API."
		addr        = ":8080"
	)

	log.Printf("starting %s (%s) on %s", title, version, addr)

	var err = rpcserver.Run(
		rpcserver.WithServerAddr(addr),
		rpcserver.WithInteractors(usecases.Interactors()...),
		rpcserver.WithInfo(new(openapi3.Info).
			WithTitle(title).
			WithVersion(version).
			WithDescription(description)),
	)
	if err != nil {
		log.Fatal(err)
	}
}
