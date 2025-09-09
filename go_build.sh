#!/bin/bash

platforms=${PLATFORMS:-linux,darwin}
archs=${ARCHS:-amd64,arm64}

version=${VERSION:-$(cat VERSION)}

echo "Building dailimage $version for $platforms on $archs"

mkdir -p artifacts
for platform in $(echo "$platforms" | tr ',' ' '); do
    for arch in $(echo "$archs" | tr ',' ' '); do
        echo "Building for $platform/$arch"
        pkg_name="dailimage-$version-$platform-$arch"
        env CGO_ENABLED=0 GOOS="$platform" GOARCH="$arch" \
            go build -v -o "artifacts/$pkg_name" -ldflags="-w -s"
    done
done
