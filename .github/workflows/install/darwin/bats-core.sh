#!/bin/bash
set -ex

# Install Bats Core.
brew uninstall --force bats
brew cleanup -s bats
brew cleanup --prune-prefix 
brew unlink bats
brew install bats-core
