package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"ipCircBuffer"
	"log"
	"net/http"
	"posts"
	"time"
)

func Run(port uint16) {
	//start := time.Now()
	err := posts.Init("posts")

	ipCircBuffer.Init()

	if err != nil {
		log.Println(err)
	}
	//log.Println("Took %s", time.Now().Sub(start))
	//log.Println(post)
	r := mux.NewRouter()
	r.HandleFunc("/posts", getPostList)
	r.HandleFunc("/posts/{Title}", getPost)
	r.HandleFunc("/posts/{Title}/paragraph/{id:[0-9]+}", getParagraph).Methods("GET")
	r.HandleFunc("/posts/{Title}/info", getInfo).Methods("GET")
	r.HandleFunc("/desktopIP", getDesktopIP).Methods("GET")
	r.HandleFunc("/desktopIP", postDesktopIP).Methods("POST")
	r.HandleFunc("/desktopIP", clearDesktopIP).Methods("DELETE")
	r.HandleFunc("/raspberryIP", getRaspberryIP).Methods("GET")
	r.HandleFunc("/raspberryIP", postRaspberryIP).Methods("POST")
	r.HandleFunc("/raspberryIP", clearRaspberryIP).Methods("DELETE")
	r.HandleFunc("/requestCount", get24HourRequests).Methods("GET")

	r.HandleFunc("/base3/decode", base3Decode).Methods("GET")
	r.HandleFunc("/base3/encode", base3Encode).Methods("GET")

	for {
		log.Printf("Running at 0.0.0.0:%d\n", port)
		log.Println(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), r))
		time.Sleep(1 * time.Second)
	}
}
