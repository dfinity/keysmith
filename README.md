# Keysmith

Hierarchical Deterministic Key Derivation for the Internet Computer

[![Build Status](https://github.com/dfinity/keysmith/workflows/build/badge.svg)](https://github.com/dfinity/keysmith/actions?query=workflow%3Abuild)

## Install

Download the release binary [here](https://github.com/dfinity/keysmith/releases).

Below is list of supported operating systems and architectures.

- darwin/amd64
- darwin/arm64
- linux/amd64
- linux/arm64
- windows/amd64

If you want to verify the authenticity of the release binary, then please also download the `SHA256.SIG` and `SHA256.SUM` files, as well as my [public key](https://sovereign.io/public.key).

Verify the SHA256 checksum of the release binary.

```text
grep "$(openssl dgst -sha256 keysmith-*.tar.gz)" SHA256.SUM
```

Verify the signature on the release binary.

```text
openssl dgst -verify public.key -signature SHA256.SIG SHA256.SUM
```

The command above should display the following output.

```text
Verified OK
```

## Usage

```text
usage: keysmith <command> [<args>]

Available Commands:
    account             Print your account identifier.
    generate            Generate your mnemonic seed.
    legacy-address      Print your legacy address.
    principal           Print your principal identifier.
    private-key         Write your private key to a file.
    public-key          Print your public key.
    version             Print the version number.
    x-public-key        Print your extended public key.
```
