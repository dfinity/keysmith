#!/bin/bash
set -ex
VERSION="0.3.0"
TARBALL="v${VERSION}.tar.gz"
pushd /tmp
curl -L -O -s "https://github.com/ztombol/bats-support/archive/${TARBALL}"
pushd /usr/local/lib
tar -f /tmp/${TARBALL} -x
ln -s bats-support-${VERSION} bats-support
popd
popd
BATS_SUPPORT="/usr/local/lib/bats-support"
echo "BATS_SUPPORT=${BATS_SUPPORT}" >> ${GITHUB_ENV}
