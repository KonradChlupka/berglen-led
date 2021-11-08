#!/bin/sh
# This script will just run forever,
# keeping the "main" binary running,
# restarting it if any file changes are detected.
while true; do
  ./main &
  PID=$!
  inotifywait ./main
  kill $PID
done
