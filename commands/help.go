package commands

import "fmt"

//
// Nmap 7.95 ( https://nmap.org )
// Usage: nmap [Scan Type(s)] [Options] {target specification}
// TARGET SPECIFICATION:
//   Can pass hostnames, IP addresses, networks, etc.
//   Ex: scanme.nmap.org, microsoft.com/24, 192.168.0.1; 10.0.0-255.1-254
//   -iL <inputfilename>: Input from list of hosts/networks
//   -iR <num hosts>: Choose random targets
//   --exclude <host1[,host2][,host3],...>: Exclude hosts/networks
//   --excludefile <exclude_file>: Exclude list from file

func Help() {
	fmt.Printf(`
	Morphsploit - Available Commands:

  echo       : Prints the provided arguments to the screen.
  exit       : Exits the Morphsploit shell.
  help       : Shows this help message.
  set        : Sets options or variables for modules.
  showOptions: Displays the current options set.
  use        : Loads and uses a module.

	Usage:
  <command> [arguments]

	Example: 
	use <module name>
	set <the value you want to set>
	`,
	)

}
