# Migrate to [quill](https://github.com/dfinity/quill)

To provide a unified user experience, we recommand [`quill`](https://github.com/dfinity/quill) which provide more comprehensive support for ledger and governance on the Internet Computer.

We will not add new features to `keysmith` or release new versions. 

Please refer to this [migration guide](quill_migration.md) to get the `quill` equivalents of `keysmith` commands.

Note that some sophasicated functionalities are not available in quill yet. If your workflow relies on those `keysmith` commands, you can keep using it.

# Keysmith

Hierarchical Deterministic Key Derivation for the Internet Computer

[![Build Status](https://github.com/dfinity/keysmith/workflows/build/badge.svg)](https://github.com/dfinity/keysmith/actions?query=workflow%3Abuild)

## Disclaimer

YOU EXPRESSLY ACKNOWLEDGE AND AGREE THAT USE OF THIS SOFTWARE IS AT YOUR SOLE RISK. AUTHORS OF THIS SOFTWARE SHALL NOT BE LIABLE FOR DAMAGES OF ANY TYPE, WHETHER DIRECT OR INDIRECT.

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
- `generate` generates your mnemonic seed and writes it to a file.
- `legacy-address` prints your legacy address.
- `principal` prints your principal identifier.
- `private-key` derives your private key and writes it to a file.
- `public-key` prints your public key.
- `shortlist` prints the available commands.
- `version` prints the version number.
- `x-private-key` derives your extended private key and writes it to a file.
- `x-public-key` prints your extended public key.

## Integration with the DFINITY Canister SDK

The [DFINITY Canister SDK](https://sdk.dfinity.org) can sign and send messages to the Internet Computer. Versions `0.7.0-beta.6` and greater provide a convenient `ledger` command that facilitates ICP token transfer. Consider the workflow below.

```bash
# Generate your mnemonic seed.
keysmith generate
# Derive your private key.
keysmith private-key
# Create an empty project.
echo {} > dfx.json
# Import your private key.
dfx identity import alternate identity.pem
# Use your private key to sign messages.
dfx identity use alternate
# Print your account identifier.
dfx ledger account-id
# Check your balance.
dfx ledger --network=https://ic0.app balance
# Send me some tokens.
dfx ledger --network=https://ic0.app transfer \
    --amount=1.23456789 \
    --memo=0 \
    --to=89e99f79ec4d81f77a6c8cb243e536e7b3244d7294fb803bcd77b3dd4e32ae36
```
