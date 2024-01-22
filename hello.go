package main

/*
	First golang programming
	for studying network and golang
*/

import (
	"flag"
	"fmt"
	"runtime"
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
	flagInterface  = flag.String("i",defaultInterface(),"NetworkInterface")
	flagHelp 	= flag.String("h","help","Help and describe what program it is.")
)

	


func main() {
	flag.Parse()
	
	if flagHelp != nil {
		fmt.Println("Program for ARP spoof in your local net.")
		fmt.Println("")
		fmt.Println("You can spoof other devices on local net.")
		fmt.Println("")
		fmt.Println("You could punished if you use this program,,,")
		fmt.Println("Be careful")
		fmt.Println("You can define NIC via -i flag")
	}

	/*
		TODO
	*/
}
