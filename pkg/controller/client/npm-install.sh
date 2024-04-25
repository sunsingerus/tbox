#!/bin/bash

# node_modules are located higher is the source tree and would be:
# 1. installed in cur dir
# 2. moved higher in the source tree afterwards

# Install locally
rm -rf node_modules
npm install

# Move higher
rm -rf ../../node_modules
mv node_modules ../..
