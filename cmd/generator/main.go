package main

import (
	"flag"
	"fmt"
	"github.com/Oppodelldog/cmdsrv/clientgen"
	"github.com/Oppodelldog/jsonrpc-remote-cli/usecases"
	"os"
)

func main() {
	var (
		interactors               = usecases.Interactors()
		clientFolder, endpointUri = processFlags()
	)

	wd, err := os.Getwd()
	assertNoErr(err)

	fmt.Printf("generating client in wd '%s'/'%s'\n", wd, clientFolder)

	assertNoErr(clientgen.SourceCode(clientFolder, endpointUri, interactors))
}

func processFlags() (string, string) {
	var (
		flags        = flag.NewFlagSet("flags", flag.ContinueOnError)
		clientFolder = flags.String("client-folder", ".", "--client-folder=/tmp/client-code")
		endpointUri  = flags.String("endpoint-uri", ".", "--endpoint-uri=http://localhost:8080/rpc")
	)
	assertNoErr(flags.Parse(os.Args[1:]))

	return *clientFolder, *endpointUri
}

func assertNoErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
