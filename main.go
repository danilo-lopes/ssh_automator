package main

import (
	"flag"
	"fmt"
	"ssh_automator/cmd"
	"ssh_automator/pkg/readfile"
)

var (
	maxThreads = 4
)

func main() {
	file := flag.String("f", "example.txt", "Arquivo com os endereços IP")
	option := flag.String("o", "test | patch", "Operacao que o programa ira fazer")
	script := flag.String("s", "script.sh", "Caminho do script caso queira plicar patch")
	flag.Parse()

	ips := readfile.File(*file)

	switch *option {
	case "test":
		cmd.SSHTester(ips, maxThreads)
	case "patch":
		cmd.SSHPatcher(ips, *script, maxThreads)
	default:
		fmt.Println("Unkwon Option")
	}
}
