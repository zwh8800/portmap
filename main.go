package main

import (
	"flag"
	"log"

	"github.com/NebulousLabs/go-upnp"
)

func main() {
	clear := flag.Bool("c", false, "clear")
	port := flag.Int("p", 0, "port")

	flag.Parse()

	if *port <= 0 || *port > 65535 {
		log.Panic("Port out of range")
	}

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

	if *clear {
		if err := igd.Clear(uint16(*port)); err != nil {
			log.Panic("igd.Forward():", err)
		}
	} else {
		if err := igd.Forward(uint16(*port), "shadowsocks"); err != nil {
			log.Panic("igd.Forward():", err)
		}
	}
}
