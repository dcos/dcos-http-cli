package main

import (
	"fmt"
	"os"

	"github.com/dcos/dcos-core-cli/pkg/pluginutil"
	"github.com/dcos/dcos-http-cli/pkg/cmd"
	"github.com/spf13/cobra"
)

func main() {
	if err := newCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func newCommand() *cobra.Command {
	dcosCmd := &cobra.Command{
		Use: "dcos",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cmd.SilenceUsage = true
		},
	}

	// This follows the CLI design guidelines for help formatting.
	dcosCmd.SetUsageFunc(pluginutil.Usage)
	dcosCmd.AddCommand(cmd.NewHTTPCommand())

	return dcosCmd
}
