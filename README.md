# dcos-http-cli

`dcos-http-cli` is a [DC/OS CLI plugin](https://docs.mesosphere.com/latest/cli/plugins/),
it let's you run HTTP requests against your clusters.

## Installation

Once you've [installed the DC/OS CLI](https://docs.mesosphere.com/latest/cli/install/) and it is attached to your cluster,
the plugin can be installed through the [dcos plugin add](https://docs.mesosphere.com/latest/cli/command-reference/dcos-plugin/dcos-plugin-add/) command.

### macOS

```console
$ dcos plugin add -u https://github.com/dcos/dcos-http-cli/releases/download/0.1.0/dcos-http-cli.darwin.zip
```

### Linux

```console
$ dcos plugin add -u https://github.com/dcos/dcos-http-cli/releases/download/0.1.0/dcos-http-cli.linux.zip
```

### Windows

```console
$ dcos plugin add -u https://github.com/dcos/dcos-http-cli/releases/download/0.1.0/dcos-http-cli.windows.zip
```

## Usage

```console
$ dcos http /dcos-metadata/dcos-version.json
{
  "version": "1.13.0-alpha",
  "dcos-image-commit": "5263b0cc09c1bf250e826cab64b902180298fa4b",
  "bootstrap-id": "1c5b7331b17e5c21e46c85ff1486389c0de1504e",
  "dcos-variant": "enterprise"
}
```

Run `dcos http` for command usage information.

## Development

This project also acts as a reference implementation for CLI plugin developers.

It follows the [DC/OS CLI guidelines](https://github.com/dcos/dcos-cli/blob/master/design/style.md)
and provides [autocompletion support](https://github.com/dcos/dcos-cli/blob/master/design/plugin.md#add-autocompletion-to-a-plugin).


### Running the plugin

In order to run the plugin from sources, you must first have the DC/OS CLI installed and attached to a cluster.

Then you can build the plugin and add it to your CLI:

```console
$ make install
```

It can now be invoked through `dcos http [...]`
