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
$ dcos http /dcos-metadata/dcos-version.json
HTTP/1.1 200 OK
Content-Length: 193
Accept-Ranges: bytes
Connection: keep-alive
Content-Type: application/json
Date: Tue, 23 Apr 2019 15:37:33 GMT
Etag: "5cbd73c7-c1"
Last-Modified: Mon, 22 Apr 2019 07:56:55 GMT
Server: openresty


{
  "version": "1.13.0-alpha",
  "dcos-image-commit": "5263b0cc09c1bf250e826cab64b902180298fa4b",
  "bootstrap-id": "1c5b7331b17e5c21e46c85ff1486389c0de1504e",
  "dcos-variant": "enterprise"
}
```

## Development

This project also acts as a reference implementation for CLI plugin developers.

It follows the [DC/OS CLI guidelines](https://github.com/dcos/dcos-cli/blob/master/design/style.md)
and provides [autocompletion support](https://github.com/dcos/dcos-cli/blob/master/design/plugin.md#add-autocompletion-to-a-plugin).
