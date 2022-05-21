package command

import (
	"fmt"
	"os/exec"
)

var (
	users      = []string{"ec2-user", "centos", "ubuntu"}
	timeout    = "ConnectTimeout=10"
	noPass     = "PasswordAuthentication=no"
	noKeyCheck = "StrictHostKeyChecking=no"
)

// Patch ira aplicar um script passado por argumento no servidor remoto
func Patch(ip, script string) string {
	for _, user := range users {
		cmd := fmt.Sprintf(`ssh -o %s -o %s -o %s "%s@%s" "bash -s" < "%s"`, timeout, noKeyCheck, noPass, user, ip, script)

		stdout, err := exec.Command("/bin/bash", "-c", cmd).Output()
		if err == nil {
			result := string(stdout)
			return result
		} else {
			msg := fmt.Sprintf("Error On Apply Patch on %s. Please Check Connectivity", ip)
			return msg
		}
	}

	return "Nothing to Do"
}
