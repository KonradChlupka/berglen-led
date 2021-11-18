#!/bin/bash

cd "$(dirname "$0")" || exit

# Deploy da thing
scp -r ./out/led-lights pi:~
