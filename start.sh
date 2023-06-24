#!/bin/sh

# script exits immediately if error occurred
set -e

# take all parameter and run it
echo "start the app"
exec "$@"
