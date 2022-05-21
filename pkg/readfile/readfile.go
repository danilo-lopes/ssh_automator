package readfile

import (
	"bufio"
	"log"
	"os"
)

// File retorna uma lista de string com endereços IP's dado um path
func File(filePath string) []string {
	var ips []string

	fileContent, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer fileContent.Close()

	scan := bufio.NewScanner(fileContent)
	for scan.Scan() {
		ips = append(ips, scan.Text())
	}

	return ips
}
