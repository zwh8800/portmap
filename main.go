package main

import (
	"log"

	"github.com/NebulousLabs/go-upnp"
)

func main() {
	igd, err := upnp.Discover()
	if err != nil {
		log.Panic("upnp.Discover():", err)
	}
	log.Println("Localtion:", igd.Location())
	extIP, err := igd.ExternalIP()
	if err != nil {
		log.Panic("igd.ExternalIP():", err)
	}
	log.Println("ExternalIP:", extIP)
}
