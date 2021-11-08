#!/bin/bash

cd "$(dirname "$0")" || exit

scp -r ./src pi:~/lights
