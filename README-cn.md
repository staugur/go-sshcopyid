## go-sshcopyid

中文 | [English](README.md)

[![Go Reference](https://pkg.go.dev/badge/tcw.im/sshcopyid.svg)](https://pkg.go.dev/tcw.im/sshcopyid)

golang实现简化版的 `ssh-copy-id` 命令效果，旨在 **无交互式** 同步密钥，自动输入密码，
之后就可用 `ssh` 免密登录，用于自动化工具。

### FAQ

- 私钥有短语的话，多次同步，目标端 authorized_keys 会累加公钥内容；无短语则提示已存在。

- 要求目标服务器有相同的port、user、passwd（可通过环境变量 `SSHCOPYID_PASSWD` 设置密码）

### 提示

主代码是 `sshcopyid.go` ，`cmd.go` 可以认为是示例，也可以用它打包成命令行。
