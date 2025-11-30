#!/bin/bash
set -e

# This script sets up the development environment for the ccwc project.

# Download necessary test data
echo "Downloading test data..."
mkdir -p tests
curl -L -o tests/test.txt 'https://www.dropbox.com/scl/fi/d4zs6aoq6hr3oew2b6a9v/test.txt?rlkey=20c9d257pxd5emjjzd1gcbn03&e=3&dl=1'
echo "Setup complete."