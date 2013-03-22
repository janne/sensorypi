package main

import (
	"log"
	"os"
	"os/exec"
	"os/user"
)

func init() {
	if usr, err := user.Current(); err != nil {
		panic(err)
	} else if usr.Uid != "0" {
		log.Fatal("sensorypi needs to be run as root")
	}
}

func main() {
	cmd := exec.Command("read-repeat", "0001:0005:01")
	cmd.Stdout = os.Stdout

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	cmd.Wait()
}
