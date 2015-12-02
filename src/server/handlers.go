package server

import (
	"desktopIP"
	"fmt"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"posts"
	"render"
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
	if ip == nil {
		ip = net.ParseIP("0.0.0.0")
	}
	fmt.Fprint(w, ip.String())
}

func postIP(w http.ResponseWriter, r *http.Request) {
	stringIP := r.FormValue("IP")
	if stringIP == "-1" {
		desktopIP.Revert()
		fmt.Fprint(w, "Reverted")
		return
	}
	desktopIP.Set(stringIP)
	fmt.Fprint(w, stringIP)
}

func clearIP(w http.ResponseWriter, r *http.Request) {
	desktopIP.Clear()
	fmt.Fprint(w, "Cleared")
}

func renderImage(w http.ResponseWriter, r *http.Request) {
	urlSlice, ok := r.Header["Url"]
	if !ok || urlSlice == nil || len(urlSlice) == 0 {
		WriteErrorString(w, "No URL header provided", 400)
		return
	}
	inputUrl := urlSlice[0]

	nameSlice, ok := r.Header["Name"]
	if !ok || nameSlice == nil || len(nameSlice) == 0 {
		// nameSlice = []string{inputUrl}
		WriteErrorString(w, "No Name header provided", 400)
		return
	}
	name := nameSlice[0]

	filePath := "/home/ubuntu/repos/marktai.com/upload/" + name

	stdOut, stdErr, err := render.Image(inputUrl, filePath)
	if err != nil {
		w.WriteHeader(500)
		WriteJson(w, map[string]interface{}{"StdOut": stdOut, "StdErr": stdErr, "Error": err.Error()})
		return
	}

	WriteJson(w, map[string]interface{}{"StdOut": stdOut, "StdErr": stdErr})

}

func renderImageGet(w http.ResponseWriter, r *http.Request) {
	inputUrl := r.FormValue("URL")
	if inputUrl == "" {
		WriteErrorString(w, "No URL query provided", 400)
		return
	}

	name := r.FormValue("Name")
	if name == "" {
		WriteErrorString(w, "No Name query provided", 400)
		return
	}

	filePath := "/home/ubuntu/repos/marktai.com/upload/" + name

	stdOut, stdErr, err := render.Image(inputUrl, filePath)
	if err != nil {
		w.WriteHeader(500)
		WriteJson(w, map[string]interface{}{"StdOut": stdOut, "StdErr": stdErr, "Error": err.Error()})
		return
	}

	http.Redirect(w, r, "http://www.marktai.com/upload/"+name, 302)

	// w.WriteHeader(302)
	// w.Header().Add("Location", "http://www.marktai.com/upload/"+name)
	// WriteJson(w, map[string]interface{}{"StdOut": stdOut, "StdErr": stdErr})

}
