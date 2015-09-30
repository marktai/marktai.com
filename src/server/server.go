package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"posts"
	"time"
)

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test")
}

func getPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	post, err := posts.GetPost(vars["Title"])
	var jsonOut []byte
	if err != nil {
		jsonOut, err = json.Marshal(err)
	} else {
		jsonOut, err = json.Marshal(post)
	}
	if err != nil {
		log.Panic(err)
		return
	}
	w.Write(jsonOut)
}

func getPostList(w http.ResponseWriter, r *http.Request) {
	posts := posts.GetPostList()

	jsonOut, err := json.Marshal(posts)
	if err != nil {
		log.Panic(err)
		return
	}
	w.Write(jsonOut)
}

func Run(port uint16) {
	//start := time.Now()
	err := posts.Init("posts")
	if err != nil {
		log.Println(err)
	}
	//log.Println("Took %s", time.Now().Sub(start))
	//log.Println(post)
	r := mux.NewRouter()
	r.HandleFunc("/healthz", healthzHandler)
	r.HandleFunc("/posts", getPostList)
	r.HandleFunc("/posts/{Title}", getPost)
	for {
		log.Printf("Running at 0.0.0.0:%d\n", port)
		log.Println(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), r))
		time.Sleep(1 * time.Second)
	}
}
