package conn

import (
	"fmt"
	"net"
	"os/exec"
	"time"
)

var (
	users      = []string{"ec2-user", "centos", "ubuntu"}
	timeout    = "ConnectTimeout=10"
	noPass     = "PasswordAuthentication=no"
	noKeyCheck = "StrictHostKeyChecking=no"
)

// Vssh testa conexão sem chave ou senha
func Vssh(ip string) string {
	sshPortIsOpen := isSSHPortOpen(ip)

	if sshPortIsOpen {
		sshConnStatus := haveSSHConn(ip)
		if sshConnStatus == "None" {
			result := fmt.Sprintf("%s;ssh connection ok;cant connect with any user;nok\n", ip)
			return result
		} else {
			result := fmt.Sprintf("%s\n", sshConnStatus)
			return result
		}
	} else {
		result := fmt.Sprintf("%s;ssh connection nok;nok\n", ip)
		return result
	}
}

func haveSSHConn(ip string) string {
	for _, user := range users {
		cmd := fmt.Sprintf(`ssh -o %s -o %s -o %s "%s@%s"`, timeout, noKeyCheck, noPass, user, ip)

		_, err := exec.Command("/bin/bash", "-c", cmd).Output()
		if err == nil {
			value := fmt.Sprintf("%s; %s; ok", ip, user)
			return value
		}
	}
	return "None"
}

func isSSHPortOpen(ip string) bool {
	timeout := time.Second * 2

	_, err := net.DialTimeout("tcp", net.JoinHostPort(ip, "22"), timeout)

	if err == nil {
		return true
	} else {
		return false
	}
}
