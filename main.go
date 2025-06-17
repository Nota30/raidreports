package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)


func main() {
	out, err := exec.Command("/opt/MegaRAID/storcli/storcli64", "/c0", "show", "all", "J").Output()
	if err != nil {
		log.Fatal(err)
	}

	var ctrls Controllers
	err = json.Unmarshal(out, &ctrls)
	if err != nil {
		log.Fatal(err)
	}

	vdlist := ctrls.Controller[0].ResponseData.VDLIST
	for i := 0; i < len(vdlist); i++ {
		fmt.Println(vdlist[i].DGVD)
		fmt.Println(vdlist[i].TYPE)
		fmt.Println(vdlist[i].Name)
		fmt.Println(vdlist[i].State)
		fmt.Println("--------------------------------")
	}
}
