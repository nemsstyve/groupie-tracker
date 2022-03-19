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
)

//on définir le temps d'actualisation de l'API
const (
	APIUpdateTime = 24 * time.Hour
)

func main() {
	go api() //On démarre notre Api
	// on définit la variable httprouter
	r := httprouter.New()
	routes(r)
	//On définit le port et le server que l'on utilisera
	const PORT = ":3080"
	// on lance notre programme sur notre localhost via le port 3080
	fmt.Printf("Listening on the port %v\nhttp://localhost%v/\n", PORT, PORT)
	log.Fatal(http.ListenAndServe(PORT, r))
}

// Cette fonction permet de rédiriger nos routes vers nos différentes pages
func routes(r *httprouter.Router) {
	//on importe nos différents fichiers HTML, CSS & JS,
	//le paramètre r correspondent à n’importe quoi jusqu’à la fin du chemin, y compris l’index du répertoire
	r.ServeFiles("/public/*filepath", http.Dir("public"))

	//Le routeur fait correspondre les demandes entrantes par la méthode de requête et le chemin d’accès
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
		time.Sleep(APIUpdateTime) // Sleep interrompt l'actualisation du serveur en cours pendant au moins la durée définit. Une durée négative ou nulle entraîne le retour immédiat de Sleep.
	}
}
