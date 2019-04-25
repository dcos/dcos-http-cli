package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/bamarni/dcos-http-cli/pkg/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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
	tpl := template.New("top")
	template.Must(tpl.Parse(`Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if .HasExample}}

Examples:
  {{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{.Name}}
	  {{.Short}}{{end}}{{end}}{{end}}
`))

	dcosCmd.SetUsageFunc(func(command *cobra.Command) error {
		err := tpl.Execute(os.Stdout, command)
		if err != nil {
			return err
		}

		if command.HasAvailableLocalFlags() {
			fmt.Println("\nOptions:")
			command.LocalFlags().VisitAll(func(f *pflag.Flag) {
				if f.Hidden {
					return
				}
				if f.Shorthand != "" && f.Name != "" {
					fmt.Printf("  -%s, --%s\n", f.Shorthand, f.Name)
				} else if f.Shorthand != "" {
					fmt.Printf("  -%s\n", f.Shorthand)
				} else {
					fmt.Printf("  --%s\n", f.Name)
				}
				fmt.Println("      " + f.Usage)
			})
		}
		if command.HasAvailableSubCommands() {
			fmt.Println(`Use "` + command.CommandPath() + ` [command] --help" for more information about a command.`)
		}
		return nil
	})

	dcosCmd.AddCommand(cmd.NewHTTPCommand())

	return dcosCmd
}
