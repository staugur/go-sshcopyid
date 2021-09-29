## go-sshcopyid

[中文](README-cn.md) | English

[![Go Reference](https://pkg.go.dev/badge/tcw.im/sshcopyid.svg)](https://pkg.go.dev/tcw.im/sshcopyid)

Golang implements a simplified version of the `ssh-copy-id` command effect,
aiming at **no interactive** synchronization key, automatic password input,
afterwards, you can log in without password using `ssh` for automated tools.

### FAQ

- If the private key has a phrase(set by environment variable `SSHCOPYID_PASSPHRASE`),
  it will be synchronized multiple times,
  and the authorized_keys of the target will accumulate the contents of the public key;
  if there is no phrase, it will prompt that it already exists.

- The target server is required to have the same port, user,
  and passwd(The `passwd` can be set through the environment variable `SSHCOPYID_PASSWD`).

### Tip

The main code is `sshcopyid.go`, `cmd.go` can be considered as an example,
it can also be packaged as an executable command `sshcopyid`.
