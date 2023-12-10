# parampiper

parampiper is a tool for manage parameters across between systems/environments. 



[![build](https://github.com/cdalar/parampiper/actions/workflows/build.yml/badge.svg)](https://github.com/cdalar/parampiper/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/cdalar/parampiper)](https://goreportcard.com/report/github.com/cdalar/parampiper)
[![codecov](https://codecov.io/gh/cdalar/parampiper/graph/badge.svg?token=7VU7H1II09)](https://codecov.io/gh/cdalar/parampiper)
[![Github All Releases](https://img.shields.io/github/downloads/cdalar/parampiper/total.svg)]()
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/cdalar/parampiper?sort=semver)
<!-- [![Known Vulnerabilities](https://snyk.io/test/github/cdalar/parampiper/main/badge.svg)](https://snyk.io/test/github/cdalar/parampiper/main) -->

## What parampiper brings 

- single source of truth (SSOT). All you parameters in one place.
- works on a simple json file. 
- support several backends (local file, azure blob)

## Installation

### MacOS

```zsh
brew install cdalar/tap/parampiper
```

### Linux

```bash
curl -sLS https://parampiper.dalar.net/get.sh | sh 
sudo install parampiper /usr/local/bin/
```

### Windows 

- download windows binary from [releases page](https://github.com/cdalar/parampiper/releases)
- unzip and copy parampiper.exe to a location in PATH


## Usage
```bash
$ parampiper
a tool to manage parameters cross different environments

Usage:
  parampiper [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  get         Get Parameter Value by Name
  help        Help about any command
  import      Import Parameters
  init        create a .pp directory with the default configuration files
  ls          List Parameters
  out         Output Parameters
  rm          Delete Parameter
  set         Add/Update Parameter
  version     Print the version number of onctl

Flags:
  -c, --config string   Configuration file (default ".pp/parampiper.yaml")
  -h, --help            help for parampiper

Use "parampiper [command] --help" for more information about a command.
```

### Initial Configuration

Create Configuration file with default values.
```bash
$ parampiper init
parampiper environment initialized
```
Default configuration file is follows:
```yaml
local_file:
  FilePath: parampiper.json
azure_blob:
  StorageAccountName: stparampiper
  ContainerName: abc
  BlobName: parampiper.json
```