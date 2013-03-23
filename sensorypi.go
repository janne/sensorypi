package main

import (
	"fmt"
	"github.com/janne/tempered"
	"log"
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
	t, err := tempered.New()
	if err != nil {
		log.Fatal(err)
	}
	defer t.Close()

	sensing, err := t.Devices[0].Sense()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f°C %.1f%%RH\n", sensing.TempC, sensing.RelHum)
}
