package main

import (
	"bytes"
	"errors"
	"fmt"
	"net/http/httputil"
	"os"
	"os/user"

	"github.com/dcos/dcos-cli/pkg/cli"
	"github.com/spf13/afero"
	"github.com/spf13/pflag"
)

func main() {
	ctx := cli.NewContext(&cli.Environment{
		Args:       os.Args,
		Input:      os.Stdin,
		Out:        os.Stdout,
		ErrOut:     os.Stderr,
		EnvLookup:  os.LookupEnv,
		UserLookup: user.Current,
		Fs:         afero.NewOsFs(),
	})

	args := ctx.Args()
	if len(args) == 2 {
		args = append(args, "-h")
	} else if len(args) < 2 {
		checkErr(ctx, errors.New("invalid usage"))
	}

	if args[2] == "--info" {
		fmt.Fprintln(ctx.Out(), "HTTP requests for your cluster!")
		return
	}

	var method, data string
	flags := pflag.NewFlagSet("http", pflag.ContinueOnError)
	flags.StringVarP(&method, "request", "X", "GET", "Specify request command to use")
	flags.StringVarP(&data, "data", "d", "", "HTTP POST data")
	err := flags.Parse(args[2:])
	checkErr(ctx, err)
	args = flags.Args()
	if len(args) != 1 {
		checkErr(ctx, errors.New("invalid usage"))
	}

	cluster, err := ctx.Cluster()
	checkErr(ctx, err)

	httpClient := ctx.HTTPClient(cluster)
	req, err := httpClient.NewRequest(method, args[0], bytes.NewBufferString(data))
	checkErr(ctx, err)

	resp, err := httpClient.Do(req)
	checkErr(ctx, err)

	respDump, err := httputil.DumpResponse(resp, true)
	checkErr(ctx, err)

	fmt.Fprintln(ctx.Out(), string(respDump))

}

func checkErr(ctx *cli.Context, err error) {
	if err != nil {
		fmt.Fprintln(ctx.ErrOut(), err)
		os.Exit(1)
	}
}
