# Keysmith

Hierarchical Deterministic Key Derivation for the Internet Computer

[![Build Status](https://github.com/dfinity/keysmith/workflows/build/badge.svg)](https://github.com/dfinity/keysmith/actions?query=workflow%3Abuild)

## Install

Download the latest release binary [here](https://github.com/dfinity/keysmith/releases).

Below is list of operating systems and architectures that we currently support.

- `darwin/amd64`
- `darwin/arm64`
- `linux/amd64`
- `linux/arm64`
- `windows/amd64`

If you want to verify the authenticity of the release binary, then please also download the `SHA256.SIG` and `SHA256.SUM` files, as well as my public key, which you can find [here](https://sovereign.io/public.key).

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

Below is list of commands and their behavior.

- `account` prints your account identifier
- `generate` generate your mnemonic seed
- `legacy-address` prints your legacy address
- `principal` prints your principal identifier
- `private-key` write your private key to a file
- `public-key` prints your public key
- `version` prints the version number
- `x-public-key` prints your extended public key
