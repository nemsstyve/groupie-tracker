package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"groupie-tracker/outils"

	"github.com/julienschmidt/httprouter" //httprouter permet un bon fonctionnement du http.Handler intégré
	                                     //httprouter indique explicitement qu'une requête ne peut correspondre qu'à une route ou à aucune
	                                    //La conception du routeur encourage la création d'API RESTful sensibles et hiérarchiques
	                                   //Vous pouvez créer des serveurs de fichiers statiques efficaces
)

//on définir le temps d'actualisation de l'API
const (
	APIUpdateTime = 24 * time.Hour
)

func main() {
	go api() //On démarre notre Api

	r := httprouter.New()
	routes(r)

	const PORT = ":3080" //On définit le port et le server que l'on utilisera

	fmt.Printf("Listening on the port %v\nhttp://localhost%v/\n", PORT, PORT)
	log.Fatal(http.ListenAndServe(PORT, r))
}

func routes(r *httprouter.Router) { // Cette fonction permet de rédiriger nos routes vers nos différentes pages
	r.ServeFiles("/public/*filepath", http.Dir("public")) //on importe nos différents fichiers HTML, CSS & JS

	r.GET("/", outils.MainPage)
	r.GET("/Artist/", outils.ArtistPage)
	r.GET("/Artist/:id", outils.ArtistPage)
	r.GET("/Info", outils.InfoPage)
	r.POST("/Search", outils.SearchPage)
	r.POST("/Filter", outils.FilterPage)
}

func api() { //Cette fonction nous limite le temps d'actualisation et renvoi une erreur après que le temps est atteind
	for {
		err := outils.Parse()
		if err != nil {
			log.Println(err)
		}
		time.Sleep(APIUpdateTime)
	}
}
