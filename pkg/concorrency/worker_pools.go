package concorrency

import (
	"ssh_automator/pkg/command"
	"ssh_automator/pkg/conn"
)

func Vssh(tasks chan string, results chan string) {
	for ip := range tasks {
		results <- conn.Vssh(ip)
	}
}

func Patch(tasks chan string, results chan string, script string) {
	for ip := range tasks {
		results <- command.Patch(ip, script)
	}
}
