package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/degenerat3/meteor/core/ent/user"
	"github.com/degenerat3/meteor/pbuf"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
)

type session struct {
	User  string
	Token string
	Exp   int64
}

var sessions []session

func status(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://" + CORESERVER + ":9999/status")
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	stat := &mcs.MCS{}
	proto.Unmarshal(data, stat)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, string(stat.GetDesc()))
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	regReq := &mcs.MCS{}
	proto.Unmarshal(data, regReq)
	adminDat := regReq.GetAuthDat()
	adminPw := adminDat.GetToken()
	hasher := sha1.New()
	hasher.Write([]byte(adminPw))
	encpw := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	knownPassArr, err := DBClient.User.Query().Where(user.Username("admin")).Select(user.FieldPassword).Strings(ctx)
	knownPass := knownPassArr[0]
	if encpw != knownPass {
		val := genUnAuth()
		w.Write(val)
		return
	}
	newUser := adminDat.GetUsername()
	newPass := adminDat.GetPassword()
	hasher = sha1.New()
	hasher.Write([]byte(newPass))
	encpw = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	_, err = DBClient.User.Create().SetUsername(newUser).SetPassword(encpw).Save(ctx)
	if err != nil {
		resp := &mcs.MCS{
			Status: 500,
			Desc:   err.Error(),
		}
		rdata, _ := proto.Marshal(resp)
		w.Write(rdata)
		return
	}
	resp := &mcs.MCS{
		Status: 200,
		Desc:   "User added",
	}
	rdata, _ := proto.Marshal(resp)
	w.Write(rdata)
	return

}

func changeUserPassword(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	webdat := &mcs.MCS{}
	proto.Unmarshal(data, webdat)
	authDat := webdat.GetAuthDat()
	tok := authDat.GetToken()
	targetUser := authDat.GetUsername()
	targetPassword := authDat.GetPassword()
	authorized := validateUserToken(tok, targetUser)
	if authorized {
		hasher := sha1.New()
		hasher.Write([]byte(targetPassword))
		encpw := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
		userObj, err := DBClient.User.Query().Where(user.Username(targetUser)).Only(ctx)
		if err != nil {
			resp := &mcs.MCS{
				Status: 500,
				Desc:   err.Error(),
			}
			rdata, _ := proto.Marshal(resp)
			w.Write(rdata)
			return
		}
		_, err = userObj.Update().SetPassword(encpw).Save(ctx)
		if err != nil {
			resp := &mcs.MCS{
				Status: 500,
				Desc:   err.Error(),
			}
			rdata, _ := proto.Marshal(resp)
			w.Write(rdata)
			return
		}
		resp := &mcs.MCS{
			Status: 200,
			Desc:   "Password upated",
		}
		rdata, _ := proto.Marshal(resp)
		w.Write(rdata)
		return
	}
	resp := &mcs.MCS{
		Status: 401,
		Desc:   "User is not authorized for this operation",
	}
	rdata, _ := proto.Marshal(resp)
	w.Write(rdata)
	return
}

func refresh(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	refReq := &mcs.MCS{}
	proto.Unmarshal(data, refReq)
	validity := checkAuth(refReq)
	if validity {
		authDat := refReq.GetAuthDat()
		tok := authDat.GetToken()
		refreshSession(tok)
		w.Write([]byte("ok"))
	} else {
		w.Write([]byte("expired"))
	}
}

func userLogin(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	loginReq := &mcs.MCS{}
	proto.Unmarshal(data, loginReq)
	authData := loginReq.GetAuthDat()
	un := authData.GetUsername()
	pw := authData.GetPassword()
	hasher := sha1.New()
	hasher.Write([]byte(pw))
	encpw := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	knownPass, err := DBClient.User.Query().Where(user.Username(un)).Select(user.FieldPassword).String(ctx)
	if err != nil {
		val := genUnAuth()
		w.Write(val)
		return
	}
	if knownPass != encpw {
		val := genUnAuth()
		w.Write(val)
		return
	}
	resp := &mcs.MCS{
		Status: 200,
		Desc:   newSession(un),
	}
	rdata, _ := proto.Marshal(resp)
	w.Write(rdata)
	return
}

func forwardReq(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resp := &mcs.MCS{
			Status: 500,
		}
		rdata, _ := proto.Marshal(resp)
		w.Write(rdata)
	}
	prot := &mcs.MCS{}
	proto.Unmarshal(data, prot)
	validity := checkAuth(prot)
	if validity == false {
		resp := genInvalidTok()
		w.Write(resp)
		return
	}

	url := "http://" + CORESERVER + ":9999" + string(r.URL.Path)
	resp, err := http.Post(url, "", bytes.NewBuffer(data))
	if err != nil {
		resp := &mcs.MCS{
			Status: 500,
		}
		rdata, _ := proto.Marshal(resp)
		w.Write(rdata)
	}
	rdata, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resp := &mcs.MCS{
			Status: 500,
		}
		rdata, _ := proto.Marshal(resp)
		w.Write(rdata)
	}
	w.Write(rdata)
}

func listForward(w http.ResponseWriter, r *http.Request) {
	url := "http://" + CORESERVER + ":9999" + string(r.URL.Path)
	resp, err := http.Get(url)
	if err != nil {
		resp := &mcs.MCS{
			Status: 500,
		}
		rdata, _ := proto.Marshal(resp)
		w.Write(rdata)
	}
	rdata, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resp := &mcs.MCS{
			Status: 500,
		}
		rdata, _ := proto.Marshal(resp)
		w.Write(rdata)
	}
	w.Write(rdata)
}

func buildReq(w http.ResponseWriter, r *http.Request) {
	bytmessage, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resp := &mcs.MCS{
			Status: 500,
		}
		rdata, _ := proto.Marshal(resp)
		w.Write(rdata)
	}
	if NESTSERVER == "" {
		fmt.Println("NEST_SERVER is not defined, unable to build")
		resp := &mcs.MCS{
			Status: 500,
		}
		rdata, _ := proto.Marshal(resp)
		w.Write(rdata)
	}
	url := "http://" + NESTSERVER + "/buildreq"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(bytmessage))
	if err != nil {
		resp := &mcs.MCS{
			Status: 500,
		}
		rdata, _ := proto.Marshal(resp)
		w.Write(rdata)
	}
	ret, _ := ioutil.ReadAll(resp.Body)
	w.Write(ret)

}
