#!/bin/bash

platforms=${PLATFORMS:-linux,darwin}
archs=${ARCHS:-amd64,arm64}

version=${VERSION:-$(cat VERSION)}

echo "Building dailimage $version for $platforms on $archs"

mkdir -p artifacts
for platform in $(echo "$platforms" | tr ',' ' '); do
    for arch in $(echo "$archs" | tr ',' ' '); do
        echo "Building for $platform/$arch"
        env GOOS="$platform" GOARCH="$arch" go build -v -o "artifacts/dailimage-$version-$platform-$arch"
    done
done
