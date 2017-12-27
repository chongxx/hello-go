package main

import (
	"encoding/json"
	"fmt"

	"github.com/bitly/go-simplejson"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s.Servers[0].ServerName)

	bitlyJson, _ := simplejson.NewJson([]byte(`{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`))
	arr, _ := bitlyJson.Get("servers").Array()
	fmt.Println(arr[0])
}
