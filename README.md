# dcos-http-cli

`dcos-http-cli` is a [DC/OS CLI plugin](https://docs.mesosphere.com/latest/cli/plugins/),
it let's you run HTTP requests against your clusters.

## Installation

Once you've [installed the DC/OS CLI](https://docs.mesosphere.com/latest/cli/install/) and it is attached to your cluster,
the plugin can be installed through the [dcos plugin add](https://docs.mesosphere.com/latest/cli/command-reference/dcos-plugin/dcos-plugin-add/) command.

### macOS

```console
$ dcos plugin add -u https://github.com/bamarni/dcos-http-cli/releases/download/0.1.0/dcos-http-cli.darwin.zip
```

### Linux

```console
$ dcos plugin add -u https://github.com/bamarni/dcos-http-cli/releases/download/0.1.0/dcos-http-cli.linux.zip
```

### Windows

```console
$ dcos plugin add -u https://github.com/bamarni/dcos-http-cli/releases/download/0.1.0/dcos-http-cli.windows.zip
```

## Usage

```console
$ dcos http /metadata/dcos-version.json
HTTP/1.1 200 OK
Transfer-Encoding: chunked
Connection: keep-alive
Content-Type: application/json
Date: Tue, 23 Apr 2019 15:20:41 GMT
Server: openresty


{"PUBLIC_IPV4": "18.184.78.101", "CLUSTER_ID": "52b4317e-7f45-4e8f-bbfc-5ee2e01f8efb"}
```

## Development

This project also acts as a reference implementation for CLI plugin developers.

It follows the [DC/OS CLI guidelines](https://github.com/dcos/dcos-cli/blob/master/design/style.md)
and provides [autocompletion support](https://github.com/dcos/dcos-cli/blob/master/design/plugin.md#add-autocompletion-to-a-plugin).
