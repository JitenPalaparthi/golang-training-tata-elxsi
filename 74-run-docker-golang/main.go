package main

import (
	"fmt"
	"os/exec"
)

func main() {

	cmd := exec.Command("docker", "run", "-d", "--name", "pg", "-p", "5432:5432", "-e", "POSTGRES_USER=postgres", "-e", "POSTGRES_PASSWORD=postgres", "-e", "POSTGRES_DB=usersdb", "postgres:16")

	if bytes, err := cmd.CombinedOutput(); err != nil {
		println(err.Error())
	} else {
		fmt.Println(string(bytes))
	}

	// 3b3b57310aef3e389e3515e869c55488cce3b9075fcc1c864a98301eee657745

	//bytes, err := cmd.Output()
}
