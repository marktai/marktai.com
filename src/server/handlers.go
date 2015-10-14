package server

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "log"
    "desktopIP"
	"net/http"
	"posts"
	"strconv"
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


func getIP(w http.ResponseWriter, r *http.Request) {
	
	ip := desktopIP.Get()
	fmt.Fprint(w, ip.String())
}

func setIP(w http.ResponseWriter, r *http.Request) {
    stringIP := r.FormValue("IP")
    desktopIP.Set(stringIP)
    fmt.Fprint(w, "Set")
}

func clearIP(w http.ResponseWriter, r *http.Request) {
    desktopIP.Clear()
    fmt.Fprint(w, "Cleared")
}
