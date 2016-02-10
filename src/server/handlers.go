package server

import (
	"base3"
	"fmt"
	"github.com/gorilla/mux"
	"ipCircBuffer"
	"net"
	"net/http"
	"nginxParser"
	"posts"
	"strconv"
	"time"
)

func getPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	post, err := posts.GetPost(vars["Title"])
	WriteOutputError(w, map[string]*posts.Post{"Post": post}, err)
}

func getPostList(w http.ResponseWriter, r *http.Request) {
	WriteJson(w, map[string][]string{"Posts": posts.GetPostList()})
}

func getParagraph(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	post, err := posts.GetPost(vars["Title"])
	if err != nil {
		WriteError(w, err, 400)
		return
	}
	i, err := strconv.Atoi(vars["id"])
	if err != nil {
		WriteError(w, err, 400)
		return
	}

	if i >= len(post.Content) {
		WriteErrorString(w, fmt.Sprintf("%s does not have paragraph %d", vars["Title"], i), 400)
		return
	}

	paragraph := post.Content[i]
	WriteJson(w, map[string]string{"Paragraph": paragraph})
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	post, err := posts.GetPost(vars["Title"])
	if err != nil {
		WriteError(w, err, 400)
		return
	}
	info := (*post).Info()
	WriteJson(w, map[string]posts.PostInfo{"Info": info})
}

func getDesktopIP(w http.ResponseWriter, r *http.Request) {

	ip := ipCircBuffer.DesktopIP.Get()
	if ip == nil {
		ip = net.ParseIP("0.0.0.0")
	}
	fmt.Fprint(w, ip.String())
}

func postDesktopIP(w http.ResponseWriter, r *http.Request) {
	stringIP := r.FormValue("IP")
	if stringIP == "-1" {
		ipCircBuffer.DesktopIP.Revert()
		fmt.Fprint(w, "Reverted")
		return
	}
	ipCircBuffer.DesktopIP.Set(stringIP)
	fmt.Fprint(w, stringIP)
}

func clearDesktopIP(w http.ResponseWriter, r *http.Request) {
	ipCircBuffer.RaspberryIP.Clear()
	fmt.Fprint(w, "Cleared")
}

func getRaspberryIP(w http.ResponseWriter, r *http.Request) {

	ip := ipCircBuffer.RaspberryIP.Get()
	if ip == nil {
		ip = net.ParseIP("0.0.0.0")
	}
	fmt.Fprint(w, ip.String())
}

func postRaspberryIP(w http.ResponseWriter, r *http.Request) {
	stringIP := r.FormValue("IP")
	if stringIP == "-1" {
		ipCircBuffer.RaspberryIP.Revert()
		fmt.Fprint(w, "Reverted")
		return
	}
	ipCircBuffer.RaspberryIP.Set(stringIP)
	fmt.Fprint(w, stringIP)
}

func clearRaspberryIP(w http.ResponseWriter, r *http.Request) {
	ipCircBuffer.RaspberryIP.Clear()
	fmt.Fprint(w, "Cleared")
}

func get24HourRequests(w http.ResponseWriter, r *http.Request) {
	requests, err := nginxParser.GetRequests(24*time.Hour, "/var/log/nginx/access.log")
	WriteOutputError(w, map[string]int{"Requests": requests}, err)
}

func base3Decode(w http.ResponseWriter, r *http.Request) {
	encodedString := r.FormValue("encoded")
	encoded, err := base3.ParseString(encodedString)
	if err != nil {
		WriteError(w, err, 400)
		return
	}
	WriteJson(w, map[string]uint{"Decoded": encoded.Decode()})
}

func base3Encode(w http.ResponseWriter, r *http.Request) {
	valueString := r.FormValue("value")
	value, err := strconv.Atoi(valueString)
	if err != nil {
		WriteError(w, err, 400)
		return
	}
	WriteJson(w, map[string]string{"Encoded": base3.New(uint(value)).String()})
}
