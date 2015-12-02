package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"posts"
	"time"
)

func Run(port uint16) {
	//start := time.Now()
	err := posts.Init("posts")
	if err != nil {
		log.Println(err)
	}
	//log.Println("Took %s", time.Now().Sub(start))
	//log.Println(post)
	r := mux.NewRouter()
	r.HandleFunc("/posts", getPostList).Methods("GET")
	r.HandleFunc("/posts/{Title}", getPost).Methods("GET")
	r.HandleFunc("/posts/{Title}/paragraph/{id:[0-9]+}", getParagraph).Methods("GET")
	r.HandleFunc("/posts/{Title}/info", getInfo).Methods("GET")
	r.HandleFunc("/desktopIP", getIP).Methods("GET")
	r.HandleFunc("/desktopIP", postIP).Methods("POST")
	r.HandleFunc("/desktopIP", clearIP).Methods("DELETE")

	r.HandleFunc("/renderImage", renderImage).Methods("POST")
	r.HandleFunc("/renderImage", renderImageGet).Methods("GET")

	for {
		log.Printf("Running at 0.0.0.0:%d\n", port)
		log.Println(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), r))
		time.Sleep(1 * time.Second)
	}
}
