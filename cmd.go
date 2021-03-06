package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
)

const version = "0.1.1"

var (
	v bool

	host       string // ip,ip,hostname
	port       uint
	user       string
	passwd     string
	identity   string
	passphrase string
)

func init() {
	flag.BoolVar(&v, "v", false, "show version and exit")
	flag.StringVar(&host, "host", "", "Target server hostname or IP, multiple separated by comma")
	flag.UintVar(&port, "port", 22, "Target server SSH port")
	flag.StringVar(&user, "user", "root", "SSH user")
	flag.StringVar(&passwd, "passwd", "", "SSH password(env)")
	flag.StringVar(&identity, "identity", "", "SSH identity file")
	flag.StringVar(&passphrase, "passphrase", "", "Private key passphrase(env)")
}

func main() {
	flag.Parse()
	if v {
		fmt.Println(version)
	} else {
		if host == "" {
			fmt.Println("host cannot be empty")
			os.Exit(1)
		}
		if passwd == "" {
			passwd = os.Getenv("SSHCOPYID_PASSWD")
		}
		if passwd == "" {
			fmt.Println("passwd cannot be empty")
			os.Exit(1)
		}
		if passphrase == "" {
			passphrase = os.Getenv("SSHCOPYID_PASSPHRASE")
		}
		var wg sync.WaitGroup

		hosts := strings.Split(host, ",")
		isok := true
		for _, ip := range hosts {
			wg.Add(1)
			go func(ip string) {
				s := SSH{
					Host: ip, Port: port, User: user, Passwd: passwd,
					Identity: identity, Passphrase: passphrase,
				}
				msg, err := s.Sync()
				if err != nil {
					isok = false
				}
				fmt.Printf("Sync to %s: %s\n", ip, msg)
				wg.Done()
			}(ip)
		}
		wg.Wait()
		if !isok {
			os.Exit(1)
		}
	}
}
