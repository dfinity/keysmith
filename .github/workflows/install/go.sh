#!/bin/bash
set -ex
VERSION="1.16"
TARBALL="go${VERSION}.linux-amd64.tar.gz"
pushd /tmp
wget -q "https://storage.googleapis.com/golang/${TARBALL}"
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
