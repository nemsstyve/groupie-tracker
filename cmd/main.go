package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/nemsstyve/groupie-tracker/controller"
	"github.com/julienschmidt/httprouter"
)

//
const (
	APIUpdateTime = 72 * time.Hour
)

func main() {
	go api()

	r := httprouter.New()
	routes(r)

	const PORT = ":8080"

	fmt.Printf("Listening on the port %v\nhttp://localhost%v/\n", PORT, PORT)
	log.Fatal(http.ListenAndServe(PORT, r))
}

func routes(r *httprouter.Router) {
	r.ServeFiles("/public/*filepath", http.Dir("public"))

	r.GET("/", controller.MainPage)
	r.GET("/Artist/", controller.ArtistPage)
	r.GET("/Artist/:id", controller.ArtistPage)
	r.GET("/Info", controller.InfoPage)
	r.POST("/Search", controller.SearchPage)
	r.POST("/Filter", controller.FilterPage)
}

func api() {
	for {
		err := controller.Parse()
		if err != nil {
			log.Println(err)
		}
		time.Sleep(APIUpdateTime)
	}
}
