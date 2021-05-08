# Keysmith

Hierarchical Deterministic Key Derivation for the Internet Computer

[![Build Status](https://github.com/dfinity/keysmith/workflows/build/badge.svg)](https://github.com/dfinity/keysmith/actions?query=workflow%3Abuild)

## Introduction

Keysmith lets you derive cryptographic keys and identifiers for the Internet Computer. Among these identifiers includes an account identifier, which indicates the source or destination of an ICP token transfer. Keysmith does not sign or send messages to the Internet Computer. Hence, Keysmith does not facilitate ICP token transfer, but rather only ICP token custody. For use cases other than custody, such as payments, consider using Keysmith in conjunction with other software, such as the [DFINITY Canister SDK](https://github.com/dfinity/keysmith#integration-with-the-dfinity-canister-sdk).

## Download

Download the latest tarball [here](https://github.com/dfinity/keysmith/releases).

## Verify

If you want to verify the authenticity of the tarball, then please also download the supplementary `SHA256.SIG` and `SHA256.SUM` files, as well as my public key, which you can find [here](https://sovereign.io/public.key).

Verify the SHA256 checksum of the tarball.

```text
grep "$(openssl dgst -sha256 keysmith-*.tar.gz)" SHA256.SUM
```

Verify the signature on the tarball.

```text
openssl dgst -sha256 -verify public.key -signature SHA256.SIG SHA256.SUM
```

The command above should display the following output.

```text
Verified OK
```

## Install

Extract the executable from the tarball.

```text
tar -f keysmith-*.tar.gz -x 
```

Add the executable to your `PATH`.

```text
sudo install -d /usr/local/bin
sudo install keysmith /usr/local/bin
```

## Usage

Below is list of commands and their behavior.

- `account` prints your account identifier.
- `generate` generates your mnemonic seed.
- `legacy-address` prints your legacy address.
- `principal` prints your principal identifier.
- `private-key` derives your private key.
- `public-key` prints your public key.
- `version` prints the version number.
- `x-public-key` prints your extended public key.

## Integration with the DFINITY Canister SDK

```bash
keysmith generate # Generate your mnemonic seed.
keysmith private-key # Derive your private key.
echo {} > dfx.json # Create an empty dfx project directory.
dfx identity import alternate identity.pem # Import your private key.
dfx identity use alternate # Use your private key to sign messages.
dfx ledger account-id # Print your account identifier.
dfx ledger --network=https://ic0.app balance # Check your balance.
dfx ledger --network=https://ic0.app transfer \ # Send me some tokens.
    --amount=1.23456789 \ # ICP is divisible by 10^(-8) just like Bitcoin.
    --memo=244837814094590 \ # Include a numeric memo.
    --to=89e99f79ec4d81f77a6c8cb243e536e7b3244d7294fb803bcd77b3dd4e32ae36
```
