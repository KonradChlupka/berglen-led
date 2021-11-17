#!/bin/bash

# Build da thing
docker buildx build --platform linux/arm/v6 -t rpi-ws281x-builder-armv6 --file Dockerfile --output out .
