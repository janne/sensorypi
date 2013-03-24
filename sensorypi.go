package main

import (
	"fmt"
	"github.com/janne/tempered"
	"log"
)

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
