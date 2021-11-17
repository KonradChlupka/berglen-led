#!/bin/bash

cd "$(dirname "$0")" || exit

# Build da thing
docker buildx build --platform linux/arm/v6 -t rpi-ws281x-builder-armv6 --file Dockerfile --output out .
