#!/bin/bash
set -ex
VERSION="0.3.0"
TARBALL="v${VERSION}.tar.gz"
pushd /tmp
wget -q "https://github.com/ztombol/bats-support/archive/${TARBALL}"
pushd /usr/local/lib
sudo tar -f /tmp/${TARBALL} -x
sudo ln -s bats-support-${VERSION} bats-support
popd
popd
BATS_SUPPORT="/usr/local/lib/bats-support"
echo "BATS_SUPPORT=${BATS_SUPPORT}" >> ${GITHUB_ENV}
