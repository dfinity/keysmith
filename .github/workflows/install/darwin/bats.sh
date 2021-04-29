#!/bin/bash
set -ex

# Install Bats.
brew unlink bats
brew install bats-core
