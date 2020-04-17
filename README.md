# env2js

This small utility extracts environment variables for dynamic injection in a JavaScript file. The main intent is to be able to use runtime configuration for containerized React app instead of standard build-time injection when building with CRA (create-react-app).

## Installation
With GO version 1.13 and later installed, run the following:
```
go get -u -ldflags "-s -w" github.com/globalvisioninc/env2js
```
It will install an executable named `env2js` in $GOPATH/bin or $HOME/go/bin if the GOPATH environment variable is not set (see [go command reference](https://golang.org/cmd/go/#hdr-Compile_and_install_packages_and_dependencies) for details).

## Usage
```
./env2js --out ./config.js [--prefix PREFIX] [--env .env.local] [--env .env]
```
`--out path` [REQUIRED] Outputs JS config file at path

`--prefix PREFIX` [OPTIONAL] Required prefix for environment variables to be included into the JS config output. Default value is `REACT_APP_`.

`--env file` [OPTIONAL] File to be parsed by [godotenv](https://github.com/joho/godotenv) to populate environment. You can specify this flag multiple times.
