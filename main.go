package main

import (
	"fmt"
	"log"
	"os/exec"
)


func main() {
	out, err := exec.Command("date").Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(out))
}
