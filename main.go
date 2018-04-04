package main

import (
	"fmt"
	"log"

	libvirt "github.com/libvirt/libvirt-go"
)

func main() {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer conn.Close()

	domains, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_RUNNING)
	if err != nil {
		log.Fatalf("failed to list running domains: %v", err)
	}

	for _, dom := range domains {
		interfaces, err := dom.ListAllInterfaceAddresses(0)
		if err != nil {
			log.Fatalf("failed to list intefaces: %v", err)
		}

		hostname, err := dom.GetName()
		if err != nil {
			log.Fatalf("failed to get hostname: %v", err)
		}

		fmt.Println("Host", hostname)
		fmt.Println("  HostName", interfaces[0].Addrs[0].Addr)
		fmt.Println("")
	}
}
