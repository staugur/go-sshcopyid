package main

import (
	"fmt"
	"strings"
	"time"

	expect "github.com/google/goexpect"
)

type SSH struct {
	Host     string
	Port     uint
	User     string
	Passwd   string
	Identity string

	Passphrase string
}

func (s SSH) Sync() (string, error) {
	cmd := fmt.Sprintf("ssh-copy-id -p %d", s.Port)
	if s.Identity != "" {
		cmd += fmt.Sprintf(" -i %s", s.Identity)
	}
	cmd += fmt.Sprintf(" %s@%s", s.User, s.Host)
	timeout := 10 * time.Second
	ge, _, err := expect.Spawn(cmd, timeout)
	if err != nil {
		return "Unconnected", err
	}
	defer ge.Close()

	caser := []expect.Caser{
		&expect.BCase{R: "Enter passphrase", S: fmt.Sprintln(s.Passphrase)},
		&expect.BCase{R: "yes/no", S: "yes\n"},
		&expect.BCase{R: "password:", S: fmt.Sprintln(s.Passwd)},
	}
	msg := ""
	for {
		if rst, _, _, err := ge.ExpectSwitchCase(caser, timeout); err != nil {
			if strings.Contains(rst, "added") {
				msg = "Succeeded"
			} else if strings.Contains(rst, "exist") {
				msg = "Existed"
			} else {
				msg = "Failed"
				return msg, fmt.Errorf(msg)
			}
			break
		}
	}
	return msg, nil
}
