package cmd

import (
	"fmt"
	"ssh_automator/pkg/concorrency"
)

func SSHTester(ips []string, maxThreads int) {
	tasks := make(chan string, len(ips))
	results := make(chan string, len(ips))

	// Quantity of conccorent ip checking
	for i := 0; i < maxThreads; i++ {
		go concorrency.Vssh(tasks, results)
	}

	// Sending tasks to the channel
	for _, ip := range ips {
		tasks <- ip
	}
	close(tasks)

	// Getting the results from the channel
	for i := 0; i < len(ips); i++ {
		result := <-results
		fmt.Println(result)
	}
}

func SSHPatcher(ips []string, script string, maxThreads int) {
	tasks := make(chan string, len(ips))
	results := make(chan string, len(ips))

	// Creating the queues
	for i := 0; i < maxThreads; i++ {
		go concorrency.Patch(tasks, results, script)
	}

	// Sending tasks to the queue
	for _, ip := range ips {
		tasks <- ip
	}
	close(tasks)

	// Consuming the queue
	for i := 0; i < len(ips); i++ {
		result := <-results
		fmt.Println(result)
	}
}
