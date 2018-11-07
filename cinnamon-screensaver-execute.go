package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/godbus/dbus"
)

type Config struct {
	Locked   []string  `json:"locked"`
	Unlocked []string  `json:"unlocked"`
}

func main() {
	// read configuration file
	home := os.Getenv("HOME")
	dat, err := ioutil.ReadFile(home + "/.cinnamon-screensaver-execute.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// parse json
	config := Config{}
	json.Unmarshal(dat, &config)
	fmt.Println("Execute on locked:")
	fmt.Println("  ", config.Locked)
	fmt.Println("Execute on unlocked:")
	fmt.Println("  ", config.Unlocked, "\n")


	conn, err := dbus.SessionBus()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to session bus:", err)
		os.Exit(1)
	}

	call := conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',interface=org.cinnamon.ScreenSaver")
	if call.Err != nil {
		fmt.Fprintln(os.Stderr, "Failed to add match:", call.Err)
		os.Exit(1)
	}
	c := make(chan *dbus.Message, 10)
	conn.Eavesdrop(c)
	fmt.Println("Listening for screensaver")
	for v := range c {
		if v.Body[0] == true {
			if len(config.Locked) > 0 {
				exec.Command(config.Locked[0], config.Locked[1:]...).Run()
			}
		} else {
			if len(config.Unlocked) > 0 {
				exec.Command(config.Unlocked[0], config.Unlocked[1:]...).Run()
			}
		}
	}
}
