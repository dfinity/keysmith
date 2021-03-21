# Keysmith

> A Supplementary CLI for DFINITY Seed Donors and Early Contributors

[![Build Status](https://github.com/enzoh/keysmith/workflows/build/badge.svg)](https://github.com/enzoh/keysmith/actions?query=workflow%3Abuild)

## Install

```
go get github.com/enzoh/keysmith
```

## Usage

Write your mnemonic seed phrase to a file.
```text
echo 'verb bottom twelve symptom plastic believe beach cargo inherit viable dice loop' > seed.txt
```

Generate a PEM file from your seed file.
```
keysmith private-key -f seed.txt -o identity.pem
```

Inspect the contents of your pem file.
```
cat identity.pem
```

Verify the contents of PEM file match what you would expect.
```
-----BEGIN EC PARAMETERS-----
BgUrgQQACg==
-----END EC PARAMETERS-----
-----BEGIN EC PRIVATE KEY-----
MHQCAQEEIAgy7nZEcVHkQ4Z1Kdqby8SwyAiyKDQmtbEHTIM+WNeBoAcGBSuBBAAK
oUQDQgAEgO87rJ1ozzdMvJyZQ+GABDqUxGLvgnAnTlcInV3NuhuPv4O3VGzMGzeB
N3d26cRxD99TPtm8uo2OuzKhSiq6EQ==
-----END EC PRIVATE KEY-----
```
