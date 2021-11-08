#!/bin/bash

cd "$(dirname "$0")" || exit

scp -r ./src pi:~/lights
ssh pi "cd lights/src; g++ main.cpp ws2812-rpi.cpp"
