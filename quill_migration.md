# Migrate to [quill](https://github.com/dfinity/quill)

You can find `quill` equivalents of `keysmith` commands.


## Generate mnemonic seed and private key (PEM)

### Before

```bash
# Generate mnemonic seed at [seed.txt]
keysmith generate -o seed.txt
# Derive PEM from [seed.txt] at [identity.pem]
keysmith private-key -f seed.txt -o identity.pem
```

### After

```bash
# Generate [seed.txt] and [identity.pem]
# The [Principal id] and [Account id] will also be shown on the console
quill generate --seed-file seed.txt --pem-file identity.pem
# If you already have the seed phrases generated, you can derive PEM with
quill generate --phrase '<SEED PHRASE>' --seed-file seed.txt --pem-file identity.pem
```

## Get principal identifier and/or account id

### Before

```bash
# Print your principal identifier
keysmith principal -f seed.txt
# Print your account identifier
keysmith account -f seed.txt
```

### After

```bash
# Print principal identifier and account id
quill --seed-file seed.txt public-ids
# or
quill --pem-file identity.pem public-ids
```