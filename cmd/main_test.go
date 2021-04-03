package main

import (
	"os"
	"os/exec"
	"testing"
)

func Test(t *testing.T)  {
	//path, err := exec.LookPath("go")
	//if err!= nil {
	//	panic(err)
	//}
	//fmt.Println(path)
	cmd := exec.Command("go","mod","init","a")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return 
	}
}
