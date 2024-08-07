package main

import (
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {

if err := godotenv.Load(); err != nil {
	panic(err)
}


	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgStore/migrations", 
		"--config",
		"./internal/store/pgStore/migrations/tern.conf",
	)

	if err := cmd.Run(); err != nil {
		println(err)
	}
}
