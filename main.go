package main

import (
	"encoding/json"
	"fmt"
	"github.com/thoj/go-ircevent"
	"io/ioutil"
	"math/rand"
	//"os"
)

type Config struct {
	Nick       string
	Trigger    string
	Realname   string
	Server     string
	Channel    string
	Machine    string
	Maintainer string
	Host       string
	Port       string
	Temp       string
	Pwm_duty   string
	Start_temp string
	Start_duty string
}

func AddCallbacks(conn *irc.Connection, config *Config) {
	conn.AddCallback("001", func(e *irc.Event) {

		pass, err := ioutil.ReadFile("./services.password")

		if err != nil {
			fmt.Println("Could not read Services password")
		}

		var channel = config.Channel
		conn.Join(channel)
		conn.Privmsg("nickserv", "identify "+string(pass))
	})
}

func main() {

	rand.Seed(64)

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println("Couldn't read config file, dying...")
		panic(err)
	}

	config := &Config{}

	err = json.Unmarshal([]byte(file), config)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(config.Nick)
	//fmt.Sprintf(config.Get(confirc[0].channels).Str())
	conn := irc.IRC(config.Nick, config.Realname)
	err = conn.Connect(config.Server)

	if err != nil {
		fmt.Println("Failed to connect.", err)
	}

	AddCallbacks(conn, config)
	conn.Loop()
}
