package cmd

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/dcos/dcos-cli/pkg/httpclient"
	"github.com/dcos/dcos-core-cli/pkg/pluginutil"
	"github.com/spf13/cobra"
)

// NewHTTPCommand creates the `http` command with all the available subcommands.
func NewHTTPCommand() *cobra.Command {
	var method, data string
	var headers []string

	cmd := &cobra.Command{
		Use:   "http <path>",
		Short: "Make HTTP requests against your cluster",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cmd.SilenceUsage = true
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			httpClient := pluginutil.HTTPClient("")

			var httpOptions []httpclient.Option
			if data != "" {
				requestFlag := cmd.Flags().Lookup("request")

				if requestFlag != nil && !requestFlag.Changed {
					method = "POST"
					httpOptions = append(httpOptions, httpclient.Header("Content-Type", "application/x-www-form-urlencoded"))
				}
			}

			for _, header := range headers {
				headerParts := strings.SplitN(header, ":", 2)
				if len(headerParts) < 2 {
					return fmt.Errorf("invalid header '%s'", header)
				}
				key, value := headerParts[0], strings.TrimSpace(headerParts[1])
				httpOptions = append(httpOptions, httpclient.Header(key, value))
			}

			req, err := httpClient.NewRequest(method, args[0], bytes.NewBufferString(data), httpOptions...)
			if err != nil {
				return err
			}

			resp, err := httpClient.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			return dumpResponse(resp)
		},
	}

	cmd.Flags().StringVarP(&data, "data", "d", "", "HTTP POST data")
	cmd.Flags().StringVarP(&method, "request", "X", "GET", "Specify request command to use")
	cmd.Flags().StringSliceVarP(&headers, "header", "H", nil, "Pass custom header(s) to server")

	return cmd
}

func dumpResponse(resp *http.Response) error {
	respDump, err := httputil.DumpResponse(resp, false)
	if err != nil {
		return err
	}
	fmt.Println(string(respDump))
	_, err = io.Copy(os.Stdout, resp.Body)
	return err
}
