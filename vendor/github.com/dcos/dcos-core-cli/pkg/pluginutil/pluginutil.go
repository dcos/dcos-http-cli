package pluginutil

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/dcos/client-go/dcos"
	"github.com/dcos/dcos-cli/pkg/httpclient"
	"github.com/dcos/dcos-cli/pkg/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// HTTPClient returns an HTTP client from a plugin runtime.
func HTTPClient(baseURL string, opts ...httpclient.Option) *httpclient.Client {
	if baseURL == "" {
		baseURL, _ = os.LookupEnv("DCOS_URL")
	}
	var baseOpts []httpclient.Option

	if acsToken, _ := os.LookupEnv("DCOS_ACS_TOKEN"); acsToken != "" {
		baseOpts = append(baseOpts, httpclient.ACSToken(acsToken))
	}

	if verbosity, _ := os.LookupEnv("DCOS_VERBOSITY"); verbosity != "" {
		baseOpts = append(baseOpts, httpclient.Logger(Logger()))
	}

	tlsInsecure, _ := os.LookupEnv("DCOS_TLS_INSECURE")
	if tlsInsecure == "1" {
		baseOpts = append(baseOpts, httpclient.TLS(&tls.Config{InsecureSkipVerify: true})) //nolint: gosec
	} else {
		tlsCAPath, _ := os.LookupEnv("DCOS_TLS_CA_PATH")
		if tlsCAPath != "" {
			rootCAsPEM, err := ioutil.ReadFile(tlsCAPath)
			if err == nil {
				certPool := x509.NewCertPool()
				if certPool.AppendCertsFromPEM(rootCAsPEM) {
					baseOpts = append(baseOpts, httpclient.TLS(&tls.Config{RootCAs: certPool}))
				}
			}
		}
	}
	return httpclient.New(strings.TrimRight(baseURL, "/"), append(baseOpts, opts...)...)
}

// NewHTTPClient returns an HTTP client from a plugin runtime.
// It is created with the `client-go` package and will eventually
// replace the httpClient of the CLI.
func NewHTTPClient() *http.Client {
	dcosConfig := dcos.NewConfig(nil)
	dcosConfig = SetConfigFromEnv(dcosConfig)
	client := dcos.NewHTTPClient(dcosConfig)
	client.Transport.(*dcos.DefaultTransport).Logger = Logger()
	return client
}

// SetConfigFromEnv updates the fields of a DC/OS Config depending
// on the environment where it is being run.
func SetConfigFromEnv(dcosConfig *dcos.Config) *dcos.Config {
	if acsToken, _ := os.LookupEnv("DCOS_ACS_TOKEN"); acsToken != "" {
		dcosConfig.SetACSToken(acsToken)
	}

	var tls dcos.TLS
	tlsInsecure, _ := os.LookupEnv("DCOS_TLS_INSECURE")
	if tlsInsecure == "1" {
		tls.Insecure = true
	} else {
		tls.RootCAsPath, _ = os.LookupEnv("DCOS_TLS_CA_PATH")
		if tls.RootCAsPath != "" {
			rootCAsPEM, err := ioutil.ReadFile(tls.RootCAsPath)
			if err == nil {
				certPool := x509.NewCertPool()
				if certPool.AppendCertsFromPEM(rootCAsPEM) {
					tls.RootCAs = certPool
				}
			}
		}
	}
	dcosConfig.SetTLS(tls)
	return dcosConfig
}

// Logger returns a logger for a given plugin runtime.
func Logger() *logrus.Logger {
	logger := &logrus.Logger{
		Out:       os.Stderr,
		Formatter: &log.Formatter{},
		Hooks:     make(logrus.LevelHooks),
	}

	verbosity, _ := os.LookupEnv("DCOS_VERBOSITY")
	switch verbosity {
	case "0":
		logger.SetLevel(logrus.WarnLevel)
	case "1":
		logger.SetLevel(logrus.InfoLevel)
	case "2":
		logger.SetLevel(logrus.DebugLevel)
		os.Setenv("DCOS_DEBUG", "1")
	}

	return logger
}

// Usage defines a help template following the UX guidelines
func Usage(command *cobra.Command) error {
	tpl := template.New("top")
	template.Must(tpl.Parse(`Usage:{{if .HasAvailableSubCommands}}
    {{.CommandPath}} [command]{{else if .Runnable}}
    {{.UseLine}}{{end}}{{if .HasExample}}

Examples:
    {{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
    {{.Name}}
        {{.Short}}{{end}}{{end}}{{end}}
`))

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
				fmt.Printf("    -%s, --%s\n", f.Shorthand, f.Name)
			} else if f.Shorthand != "" {
				fmt.Printf("    -%s\n", f.Shorthand)
			} else {
				fmt.Printf("    --%s\n", f.Name)
			}
			fmt.Println("        " + f.Usage)
		})
	}
	if command.HasAvailableSubCommands() {
		fmt.Println("\n" + `Use "` + command.CommandPath() + ` [command] --help" for more information about a command.`)
	}
	return nil
}
