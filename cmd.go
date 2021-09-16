package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
)

const version = "0.1.0"

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
	flag.StringVar(&passwd, "passwd", "", "SSH password")
	flag.StringVar(&identity, "identity", "", "SSH identity file")
	flag.StringVar(&passphrase, "passphrase", "", "Private key passphrase")
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
			fmt.Println("passwd cannot be empty")
			os.Exit(1)
		}
		var wg sync.WaitGroup

		hosts := strings.Split(host, ",")
		for _, ip := range hosts {
			wg.Add(1)
			go func(ip string) {
				s := SSH{
					Host: ip, Port: port, User: user, Passwd: passwd,
					Identity: identity, Passphrase: passphrase,
				}
				msg, _ := s.Sync()
				fmt.Printf("Sync to %s: %s\n", ip, msg)
				wg.Done()
			}(ip)
		}
		wg.Wait()
	}
}
