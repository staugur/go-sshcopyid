## go-sshcopyid

[中文](README-cn.md) | English

[![Go Reference](https://pkg.go.dev/badge/tcw.im/sshcopyid.svg)](https://pkg.go.dev/tcw.im/sshcopyid)

Golang implements a simplified version of the `ssh-copy-id` command effect,
aiming at **no interactive** synchronization key, automatic password input,
afterwards, you can log in without password using `ssh` for automated tools.

### FAQ

- If the private key has a phrase, it will be synchronized multiple times, and the authorized_keys of the target will accumulate the contents of the public key; if there is no phrase, it will prompt that it already exists.

- The target server is required to have the same port, user, and passwd.

### Tip

The main code is `sshcopyid.go`, `cmd.go` can be considered as an example, or it can be packaged into a command line.
