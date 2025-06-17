package main

import (
	"fmt"
	"log"
	"os/exec"
)


func main() {
	out, err := exec.Command("/opt/MegaRAID/storcli/storcli64 /c0 show all J").Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(out))
}
