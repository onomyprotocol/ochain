package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/onomyprotocol/onomy-sdk/client"
	// "github.com/onomyprotocol/onomy-sdk/client/flags"
	// sdk "github.com/onomyprotocol/onomy-sdk/types"

	"github.com/onomyprotocol/ochain/x/ochain/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group ochain queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1

	return cmd
}
