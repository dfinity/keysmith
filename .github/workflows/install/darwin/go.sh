#!/bin/bash
set -ex
ARCH="$(uname -m)"
VERSION="1.16"
TARBALL="go${VERSION}.darwin-${ARCH}.tar.gz"
pushd /tmp
curl -L -O -s "https://storage.googleapis.com/golang/${TARBALL}"
pushd /usr/local
sudo tar -f /tmp/${TARBALL} -x
popd
popd
GOROOT="/usr/local/go"
GOPATH="${HOME}/go"
echo "GOROOT=${GOROOT}" >> ${GITHUB_ENV}
echo "GOPATH=${GOPATH}" >> ${GITHUB_ENV}
echo "${GOROOT}/bin" >> ${GITHUB_PATH}
echo "${GOPATH}/bin" >> ${GITHUB_PATH}
mkdir ${GOPATH}
