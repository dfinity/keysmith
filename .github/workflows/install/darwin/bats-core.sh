#!/bin/bash
set -ex

# Install Bats Core.
brew unlink bats || :
brew install bats-core
