# `ecsu`

**This is just a toy project; consider it useless and insecure unless proven otherwise.**

## Installation

Run `make` and copy the resulting `ecsu` binary to the target machine, then `chown root:root` and `chmod u+s` it.

> **Warning**
> treat the generated `ecsu-keygen` binary as a key file; anybody with enough permissions to read or run it will also be able to unlock `ecsu`

## Usage
```console
$ ecsu-keygen
wJI7yT9ZZWTiqyYoIOlmkv3WUqDLss9wHnF7ePA8KWqXbU3KfKRDZpWRXK5l/46O7oIAtAtJy89KPKX6wKtvCA==
```

```console
$ ecsu id
time-based code: wJI7yT9ZZWTiqyYoIOlmkv3WUqDLss9wHnF7ePA8KWqXbU3KfKRDZpWRXK5l/46O7oIAtAtJy89KPKX6wKtvCA==
uid=0(root) gid=0(root) groups=0(root)
```

## Codes

The time-based codes used by this tool are just the Base64-encoded result of calling [`crypto/ed25519.Sign`](https://pkg.go.dev/crypto/ed25519#Sign) with the current timestep number (i.e. Unix time in 30 second increments)
