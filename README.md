# vrsncode

```
Vrsncode is a tool to fetch and update Android version code with respect to
the latest version code for an app deployed to play store

Usage:
  vrsncode [command]

Available Commands:
  fetch       A brief description of your command
  help        Help about any command
  increment   A brief description of your command

Flags:
  -c, --config string    config file (default is .vrsncode)
  -h, --help             help for vrsncode
  -k, --key string       key file (e.g service account json file)
  -p, --package string   package name

Use "vrsncode [command] --help" for more information about a command.
```

## Usage

`vrsncode -k service-account.json -p example.com.demo`

## Installation
`go get github.com/nathanmkaya/vrsncode`
