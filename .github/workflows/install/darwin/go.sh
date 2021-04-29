#!/bin/bash
set -ex

# Define operating system and architecture.
GOOS=darwin
case "$(uname -m)" in
    x86_64)
        GOARCH="amd64";;
    arm64)
        GOARCH="arm64";;
    *)
        echo "Error: Unsupported architecture!"
        exit 1
esac

# Define version and tarball.
VERSION="1.16"
TARBALL="go${VERSION}.${GOOS}-${GOARCH}.tar.gz"

# Install Go compiler.
pushd /tmp
curl -L -O -s "https://storage.googleapis.com/golang/${TARBALL}"
pushd /usr/local
sudo tar -f /tmp/${TARBALL} -x
popd
popd

# Configure workspace.
GOROOT="/usr/local/go"
GOPATH="${HOME}/go"
echo "GOOS=${GOOS}" >> ${GITHUB_ENV}
echo "GOARCH=${GOARCH}" >> ${GITHUB_ENV}
echo "GOROOT=${GOROOT}" >> ${GITHUB_ENV}
echo "GOPATH=${GOPATH}" >> ${GITHUB_ENV}
echo "${GOROOT}/bin" >> ${GITHUB_PATH}
echo "${GOPATH}/bin" >> ${GITHUB_PATH}
mkdir ${GOPATH}
