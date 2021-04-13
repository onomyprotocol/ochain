package main

import (
	"os"

	svrcmd "github.com/onomyprotocol/cosmos-sdk/server/cmd"
	"github.com/onomyprotocol/ochain/app"
	"github.com/onomyprotocol/ochain/cmd/ochaind/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
