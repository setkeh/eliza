package main

import (
	//"encoding/json"
	"fmt"
	"github.com/stretchr/objx"
	"github.com/thoj/go-ircevent"
	"io/ioutil"
	"math/rand"
	"os"
)

func request(conf string) (objx.Map, error) {
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println("Couldn't read config file, dying...")
		panic(err)
		os.Exit(1)
	}

	data, err := objx.FromJSON(string(file))
	if err != nil {
		return nil, err
	}

	return data, nil
}

//func AddCallbacks(conn *irc.Connection, conf string) {
//	conn.AddCallback("001", func(e *irc.Event) {

//		config, err := request(conf)
//if err != nil {
//	return err
//}
//var channels = config.Get(irc[0].channels).Str()
//fmt.Sprintf(channels)
//		fmt.Sprintf(config.Get(irc[0].channels).Str())

// conn.Join(channels)
//	})
//}

func main() {
	var conf string
	rand.Seed(64)

	config, err := request(conf)
	//fmt.Sprintf(config.Get(confirc[0].channels).Str())
	conn := irc.IRC(config.Get("confirc[0].nick").Str(), config.Get("confirc[0].realname").Str())
	err = conn.Connect(config.Get("confirc[0].server").Str())

	if err != nil {
		fmt.Println("Failed to connect.")
		panic(err)
	}

	//AddCallbacks(conn, conf)
	conn.Loop()
}
