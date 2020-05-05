#!/usr/bin/env sh

package_name=env2js

rm -rf ./dist/*
for i in "darwin amd64" "linux amd64" "linux 386" "windows amd64" "windows 386"
do
    set -- $i
    GOOS=$1
    GOARCH=$2
    GO_BUILD_FLAGS='-s -w'

    output_dir=$GOOS'-'$GOARCH
    output_name=$package_name

    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    docker run --rm -e GOOS=$GOOS -e GOARCH=$GOARCH -w /app \
        -v $(pwd)/dist:/dist -v $(dirname $(pwd)):/app \
        golang:1.13-alpine \
        go build -ldflags "${GO_BUILD_FLAGS}" -o "/dist/${output_dir}/${output_name}" /app/main.go
done
