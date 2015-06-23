package main

import (
	"encoding/json"
	"fmt"
	//"github.com/thoj/go-ircevent"
	"io/ioutil"
	"math/rand"
	//"os"
)

type Config struct {
	Confirc struct {
		Nick     string
		Trigger  string
		Realname string
		Server   string
		Channels struct {
			Ldc  string
			Tskp string
		}
		AutoReconnect string
		AutoRejoin    string
		FloodDelay    string
		Secure        string
		Debug         string
		Vhost         string
	}
	General struct {
		Machine    string
		Maintainer string
	}
	Docker struct {
		Host string
		Port string
	}
	Odt struct {
		Temp       string
		Pwm_duty   string
		Start_temp string
		Start_duty string
	}
}

/*func AddCallbacks(conn *irc.Connection, config *Config) {
	conn.AddCallback("001", func(e *irc.Event) {

		fmt.Println(config)
		//		var channels = config.Channels

		//conn.Join(channels)
	})
}*/

func main() {
	var conf = "./config.json"
	config := &Config{}
	rand.Seed(64)

	file, err := ioutil.ReadFile(conf)

	if err != nil {
		fmt.Println("Couldn't read config file, dying...")
		panic(err)
	}

	json.Unmarshal([]byte(file), &config)

	fmt.Println(config)
	//fmt.Sprintf(config.Get(confirc[0].channels).Str())
	//	conn := irc.IRC(config.Get("confirc[0].nick").Str(), config.Get("confirc[0].realname").Str())
	//	err = conn.Connect(config.Get("confirc[0].server").Str())

	//	if err != nil {
	//		fmt.Println("Failed to connect.", err)
	//	}

	//AddCallbacks( /*conn, */ config)
	//conn.Loop()
}
