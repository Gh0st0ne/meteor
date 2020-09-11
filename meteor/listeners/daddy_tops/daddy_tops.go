package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", status)
	http.HandleFunc("/status", status)
	http.HandleFunc("/register/bot", forwardReq)
	http.HandleFunc("/register/host", forwardReq)
	http.HandleFunc("/register/group", forwardReq)
	http.HandleFunc("/register/hostgroup", forwardReq)
	http.HandleFunc("/add/action/single", forwardReq)
	http.HandleFunc("/add/action/group", forwardReq)
	http.HandleFunc("/add/result", forwardReq)
	http.HandleFunc("/list/bots", listForward)
	http.HandleFunc("/list/hosts", listForward)
	http.HandleFunc("/list/groups", listForward)
	log.Fatal(http.ListenAndServe(":8888", nil))
	return
}
