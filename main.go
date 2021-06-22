package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

var options int
var name string
var dropDelay float64
var unixINPUT int64

var config map[string]interface{}

func main() {

	webhookvar, _ := ioutil.ReadFile("config.json")
	json.Unmarshal(webhookvar, &config)

	e, err := ioutil.ReadFile("accounts.txt")
	if err != nil {
		panic(err)
	}

	test := strings.Split(string(e), ":")

	b, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	c := color.New(color.FgHiRed).Add(color.Bold).Add(color.Underline)
	b3 := color.New(color.FgHiBlue).Add(color.Bold)
	Gold := color.New(color.FgHiYellow).Add(color.Bold)

	c.Println(
		`
███╗░░░███╗███████╗██████╗░██╗░░░██╗░██████╗░█████╗░
████╗░████║██╔════╝██╔══██╗██║░░░██║██╔════╝██╔══██╗
██╔████╔██║█████╗░░██║░░██║██║░░░██║╚█████╗░███████║
██║╚██╔╝██║██╔══╝░░██║░░██║██║░░░██║░╚═══██╗██╔══██║
██║░╚═╝░██║███████╗██████╔╝╚██████╔╝██████╔╝██║░░██║
╚═╝░░░░░╚═╝╚══════╝╚═════╝░░╚═════╝░╚═════╝░╚═╝░░╚═╝`)

	b3.Println(string(b))

	url1 := ("https://api.minecraftservices.com/minecraft/profile")

	b3.Print("\nChoose an option:\n")
	fmt.Print(`
  [1]: SFA Sniper
  [2]: MFA Sniper
  [3]: GC
  [4]: Microsoft
  `)
	b3.Print("\n>: ")
	fmt.Scanln(&options)

	email := test[0]
	password := test[1]

	if options == 1 {

		/*
			dropTime, err := getDropTime(name)
			if err != nil {
				fmt.Println(err)
				return
			}
		*/

		bearer := mojangLogin2(email, password)

		if bearer == "" {
			fmt.Println("\n[INFO] Account isnt useable!")
			os.Exit(0)
		}

		b, err := ioutil.ReadFile("SFALogo.txt")
		if err != nil {
			panic(err)
		}

		fmt.Print(string(b), "\n")

		fmt.Println("\n[INFO] Name to Snipe:")
		b3.Print(">: ")
		fmt.Scanln(&name)
		fmt.Println("[INFO] Delay:")
		b3.Print(">: ")
		fmt.Scanln(&dropDelay)
		fmt.Println("[INFO] Unix Timestamp [TEMP]:")
		b3.Print(">: ")
		fmt.Scanln(&unixINPUT)
		/*
			dropTime_UNIX := unixINPUT

			testee := time.Unix(dropTime_UNIX, 0)

			drop_piece := time.Until(testee) + time.Duration(dropDelay)*time.Millisecond

			fmt.Println("\n[+] Dropping at @:", testee, drop_piece, "\n")

			time.Sleep(drop_piece)
		*/

		dropTime_UNIX := unixINPUT

		snipe_time := dropTime_UNIX - (int64(dropDelay / 1000))

		fmt.Println("\n[+] Dropping at @:", time.Unix(dropTime_UNIX, 0), "\n")

		for time.Now().Unix() < snipe_time {
			time.Sleep(1 * time.Millisecond)
		}

		Gold.Println("•:•.•:•.•:•:☾۞☽•:•.•:•.•:•:\n")

		go socketSending(bearer, name, dropDelay)

		time.Sleep(1 * time.Second)

		Gold.Println("\n•:•.•:•.•:•:☾۞☽•:•.•:•.•:•:")

		time.Sleep(1 * time.Second)

	} else if options == 2 {

		if (len(test)) != 5 {
			fmt.Println("[INFO] Incorrect length for MFA check accounts.txt")
			os.Exit(0)
		}

		bearer := mojangLogin1()

		if bearer == "" {
			fmt.Println("\n[INFO] Account isnt useable!")
			os.Exit(0)
		}

		b, err := ioutil.ReadFile("SFALogo.txt")
		if err != nil {
			panic(err)
		}

		fmt.Print(string(b), "\n")

		fmt.Println("\n[INFO] Name to Snipe:")
		b3.Print(">: ")
		fmt.Scanln(&name)
		fmt.Println("[INFO] Delay:")
		b3.Print(">: ")
		fmt.Scanln(&dropDelay)
		fmt.Println("[INFO] Unix Timestamp [TEMP]:")
		b3.Print(">: ")
		fmt.Scanln(&unixINPUT)

		dropTime_UNIX := unixINPUT

		snipe_time := dropTime_UNIX - (int64(dropDelay / 1000))

		fmt.Println("\n[+] Dropping at @:", time.Unix(dropTime_UNIX, 0), "\n")

		for time.Now().Unix() < snipe_time {
			time.Sleep(1 * time.Millisecond)
		}

		Gold.Println("•:•.•:•.•:•:☾۞☽•:•.•:•.•:•:\n")

		go socketSending(bearer, name, dropDelay)

		time.Sleep(1 * time.Second)

		Gold.Println("\n•:•.•:•.•:•:☾۞☽•:•.•:•.•:•:")

		time.Sleep(1 * time.Second)

	} else if options == 3 {

		b, err := ioutil.ReadFile("GCLogo.txt")
		if err != nil {
			panic(err)
		}

		fmt.Println("\n", string(b))

		var bearerGC string

		fmt.Println("\n[INFO] Bearer:")
		b3.Print(">: ")
		fmt.Scanln(&bearerGC)
		fmt.Println("\n[INFO] Name to Snipe:")
		b3.Print(">: ")
		fmt.Scanln(&name)
		fmt.Println("[INFO] Delay:")
		b3.Print(">: ")
		fmt.Scanln(&dropDelay)
		fmt.Println("[INFO] Unix Timestamp [TEMP]:")
		b3.Print(">: ")
		fmt.Scanln(&unixINPUT)

		dropTime_UNIX := unixINPUT

		snipe_time := dropTime_UNIX - (int64(dropDelay / 1000))

		fmt.Println("\n[+] Dropping at @:", time.Unix(dropTime_UNIX, 0), "\n")

		for time.Now().Unix() < snipe_time {
			time.Sleep(1 * time.Millisecond)
		}

		Gold.Println("•:•.•:•.•:•:☾۞☽•:•.•:•.•:•:\n")

		go testingGC(name, bearerGC, dropDelay)

		time.Sleep(1 * time.Second)

		Gold.Println("\n•:•.•:•.•:•:☾۞☽•:•.•:•.•:•:")

		time.Sleep(2 * time.Second)

	} else if options == 4 {
		b, err := ioutil.ReadFile("MCLogo.txt")
		if err != nil {
			panic(err)
		}

		fmt.Print(string(b), "\n")

		var bearerMS string

		fmt.Println("\n[INFO] Enter Bearer:")
		b3.Print(">: ")
		fmt.Scanln(&bearerMS)
		fmt.Println("\n[INFO] Name to Snipe:")
		b3.Print(">: ")
		fmt.Scanln(&name)
		fmt.Println("[INFO] Delay:")
		b3.Print(">: ")
		fmt.Scanln(&dropDelay)
		fmt.Println("[INFO] Unix Timestamp [TEMP]:")
		b3.Print(">: ")
		fmt.Scanln(&unixINPUT)

		dropTime_UNIX := unixINPUT

		snipe_time := dropTime_UNIX - (int64(dropDelay / 1000))

		fmt.Println("\n[+] Dropping at @:", time.Unix(dropTime_UNIX, 0), "\n")

		for time.Now().Unix() < snipe_time {
			time.Sleep(1 * time.Millisecond)
		}

		Gold.Println("•:•.•:•.•:•:☾۞☽•:•.•:•.•:•:\n")

		go socketSendingMS(url1, bearerMS, name, dropDelay)

		Gold.Println("\n•:•.•:•.•:•:☾۞☽•:•.•:•.•:•:")

		time.Sleep(1 * time.Second)

	}

}
