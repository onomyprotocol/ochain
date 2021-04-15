package main

import (
	"os"

	"github.com/onomyprotocol/ochain/app"
	"github.com/onomyprotocol/ochain/cmd/ochaind/cmd"
	svrcmd "github.com/onomyprotocol/onomy-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
