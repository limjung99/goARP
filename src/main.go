package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func defaultInterface() string {
	switch runtime.GOOS {
	case "darwin":
		return "en0"
	case "windows":
		return "Ethernet"
	case "linux":
		return "ens33"
	}
	return "eth0"
}

var (
	flagInterface      = flag.String("i", defaultInterface(), "NetworkInterface")
	flagHelp           = flag.Bool("h", false, "Help and describe what program it is.")
	flagListInterfaces = flag.Bool("l", false, "LIst available interfaces and exit.")
)

func main() {
	flag.Parse()

	if *flagHelp {
		fmt.Println("Program for ARP spoof in your local net.")
		fmt.Println("You can spoof other devices on local net.")
		fmt.Println("You could punished if you use this program,,,")
		fmt.Println("Be careful!")
		fmt.Println("	-i : You can define your network interface to spoof local network.")
		fmt.Println("	-l : list up your all network interface hardware and it's mac address.")
	}

	if *flagListInterfaces {
		ifaces, err := net.Interfaces()

		if err != nil {
			log.Fatal("Failed to show-up interfaces :", err)
		}

		for idx, iface := range ifaces {
			if iface.HardwareAddr == nil {
				continue
			}
			fmt.Printf("%d interface \n", idx)
			fmt.Printf("	MAC Address : %s \n", iface.HardwareAddr)
			fmt.Printf("	Interface Name : %s \n", iface.Name)
		}

		os.Exit(0)
	}

	/*
		TODO
			1. List up all host machines name and write it to standard output
			2. User can select an NIC from information
			3. Get an gateway IP and MAC address ( but how ? )
			4. Spoof gateway
			5. If possible , relay packets in local network.
	*/

	handle , err := pcap.OpenLive(*flagInterface,1600,true,pcap.BlockForever)
	if err != nil {
		panic(err)
	}

	// lazy resource return
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle,handle.LinkType())

	for packet := range packetSource.Packets() {
		fmt.Println("packet:",packet)
	}


}
