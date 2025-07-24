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

	var mailstring string

	vdlist := ctrls.Controller[0].ResponseData.VDLIST
	for i := 0; i < len(vdlist); i++ {
		if (vdlist[i].State != "Optl" && vdlist[i].State != "GHS") {
			mailstring = mailstring + "\nUh oh! Looks like something is wrong with a Virtual Drive...." + 
			"\n-------------------------------- Details --------------------------------" +
			"\n- DG/VD = " + vdlist[i].DGVD +
			"\n- DG/VD = " + vdlist[i].DGVD +
			"\n- TYPE = " + vdlist[i].TYPE +
			"\n- NAME = " + vdlist[i].Name +
			"\n- STATE = " + vdlist[i].State +
			"\n----------------------------------------------------------------"
		}
	}

	pdlist := ctrls.Controller[0].ResponseData.PDLIST
	for i := 0; i < len(pdlist); i++ {
		if (pdlist[i].State != "Onln" && pdlist[i].State != "DHS" && vdlist[i].State != "GHS") {
			mailstring = mailstring + "\nUh oh! Looks like something is wrong with a Physical Drive...." +
			"\n-------------------------------- Details --------------------------------" +
			"\n- EID:Slot = " + pdlist[i].EIDSlt +
			"\n- Model = " + pdlist[i].Model +
			"\n- STATE = " + pdlist[i].State +
			"\n----------------------------------------------------------------"
		}
	}

	fmt.Println(mailstring)
}
