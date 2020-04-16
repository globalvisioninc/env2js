# env2js

This small utility extracts environment variables for dynamic injection in a JavaScript file.

## Usage
```
./env2js --out ./config.js [--prefix PREFIX] [--env .env.local] [--env .env]
```
`--out path` [REQUIRED] Outputs JS config file at path

`--prefix PREFIX` [OPTIONAL] Required prefix for environment variables to be included into the JS config output. Default value is `REACT_APP_`.

`--env file` [OPTIONAL] File to be parsed by [godotenv](https://github.com/joho/godotenv) to populate environment. You can specify this flag multiple times.

## Build
With GO version 1.13 and later installed, run the following:
```
go build -o ./dist
```
It will generate an executable named `env2js` in the dist folder.
