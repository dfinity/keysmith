#!/bin/bash
set -ex

# Define version and tarball.
VERSION="0.3.0"
TARBALL="v${VERSION}.tar.gz"

# Install Bats Support.
pushd /tmp
curl -L -O -s "https://github.com/ztombol/bats-support/archive/${TARBALL}"
pushd /usr/local/lib
tar -f /tmp/${TARBALL} -x
ln -s bats-support-${VERSION} bats-support
popd
popd

# Configure workspace.
BATS_SUPPORT="/usr/local/lib/bats-support"
echo "BATS_SUPPORT=${BATS_SUPPORT}" >> ${GITHUB_ENV}
