package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/opbteam/spammessage/data"
	"github.com/opbteam/spammessage/spammer"
	"github.com/opbteam/spammessage/util"
)

func main() {
	fmt.Printf(color.HiBlueString("\n\n░█████╗░██████╗░███████╗███╗░░██╗  ░██████╗██████╗░░█████╗░███╗░░░███╗\n██╔══██╗██╔══██╗██╔════╝████╗░██║  ██╔════╝██╔══██╗██╔══██╗████╗░████║\n██║░░██║██████╔╝█████╗░░██╔██╗██║  ╚█████╗░██████╔╝███████║██╔████╔██║\n██║░░██║██╔═══╝░██╔══╝░░██║╚████║  ░╚═══██╗██╔═══╝░██╔══██║██║╚██╔╝██║\n╚█████╔╝██║░░░░░███████╗██║░╚███║  ██████╔╝██║░░░░░██║░░██║██║░╚═╝░██║\n░╚════╝░╚═╝░░░░░╚══════╝╚═╝░░╚══╝  ╚═════╝░╚═╝░░░░░╚═╝░░╚═╝╚═╝░░░░░╚═╝\n 		%v %v"), color.HiWhiteString("Author:"), color.HiRedString("Phuongaz\n\n"))
	log := log.Default()
	util.InitColor()
	if err := data.InitializeToken(log); err != nil {
		log.Fatal(err)
	}

	log.Println("XBL: Token initialized")
	log.Printf("<Command Type> start <host:port> <message>")

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		text := s.Text()
		if len(text) == 0 {
			fmt.Printf("<Command Type> start <host:port> <message>")
			continue
		}
		arr := strings.Split(text, " ")
		if arr[0] == "start" {
			log.Println("XBL: Starting")
			spam := &spammer.MessageData{
				Message: strings.Join(arr[2:], " "),
				Address: arr[1],
				Log:     log,
			}
			spam.Run()
		}
	}
}
