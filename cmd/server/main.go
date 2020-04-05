package main

import (
	"log"
	"os/exec"
	"strconv"

	"github.com/davidcorbin/coreutils-date-http-api/internal"
)

func main() {
	// Check that program is running as root
	cmd := exec.Command("id", "-u")
	output, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	// output has trailing \n
	// need to remove the \n
	// otherwise it will cause error for strconv.Atoi
	// log.Println(output[:len(output)-1])

	// 0 = root, 501 = non-root user
	i, err := strconv.Atoi(string(output[:len(output)-1]))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Running as user:", i)

	if i != 0 {
		log.Fatal("This program must be run as root!")
	}

	internal.StartHTTPServer()
}
