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
		if (vdlist[i].State != "Optl") {
			fmt.Println("Uh oh! Looks like something is wrong with a Virtual Drive....")
			fmt.Println("-------------------------------- Details --------------------------------")
			fmt.Println("- DG/VD = " + vdlist[i].DGVD)
			fmt.Println("- TYPE = " + vdlist[i].TYPE)
			fmt.Println("- NAME = " + vdlist[i].Name)
			fmt.Println("- STATE = " + vdlist[i].State)
			fmt.Println("----------------------------------------------------------------")
		}
	}

	pdlist := ctrls.Controller[0].ResponseData.PDLIST
	for i := 0; i < len(pdlist); i++ {
		if (pdlist[i].State != "Onln") {
			fmt.Println("Uh oh! Looks like something is wrong with a Physical Drive....")
			fmt.Println("-------------------------------- Details --------------------------------")
			fmt.Println("- EID:Slot = " + pdlist[i].EIDSlt)
			fmt.Println("- Model = " + pdlist[i].Model)
			fmt.Println("- STATE = " + pdlist[i].State)
			fmt.Println("----------------------------------------------------------------")
		}
	}
}
