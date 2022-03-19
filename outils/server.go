package outils

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"  //import le fichier et le chemins à utiliser pour notre travail
	"strconv"
	"text/template"  //import les données textes et templates 

	"github.com/julienschmidt/httprouter"   //cet import nous permet de créer des serveurs de fichiers statiques efficaces
)

// MainPage - page avec tous les artistes/groupes
func MainPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer recoverHandle(rw)
	// on définit les chemins d'accès à nos fichiers html
	path := filepath.Join("public", "html", "index.html")

	tmpl, err := template.ParseFiles(path)
	// là on definit si erreur sur les fichiers, alors il retourne la ligne d'erreur sur le terminal
	if err != nil { 
		errorHandler(rw, http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(rw, InfoArtists)
	if err != nil {
		errorHandler(rw, http.StatusInternalServerError)
		return
	}
}

// ArtistPage - Artist/Bands Information Page
func ArtistPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer recoverHandle(rw)

	if r.URL.Path == "/Artist/" {
		http.Redirect(rw, r, "/", http.StatusSeeOther)
		return
	}
	path := filepath.Join("public", "html", "artist.html")  // chemin d'accès qui joint n'importe quel nombre d'éléments de chemin en un seul chemin, 
	                                                       //en les séparant avec un séparateur spécifique au système d'exploitation. Les éléments vides sont ignorés. Le résultat est Nettoyé.
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		errorHandler(rw, http.StatusInternalServerError)
		return
	}
	id, err := strconv.Atoi(r.URL.Path[8:])
	if err != nil {
		errorHandler(rw, http.StatusBadRequest)
		return
	}
	if id < 1 || id > len(InfoArtists.Art) {
		errorHandler(rw, http.StatusNotFound)
		return
	}
	id = id - 1
	artistInfo := &OneArtistOrBand{
		Art: InfoArtists.Art[id],
		Rel: InfoArtists.Rel.Index[id],
	}

	err = tmpl.Execute(rw, artistInfo)
	if err != nil {
		errorHandler(rw, http.StatusInternalServerError)
		return
	}
}

// SearchPage - page de recherche avec toutes les informations de l'artiste
func SearchPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer recoverHandle(rw)
	InfoArtists.Search = nil
	path := filepath.Join("public", "html", "search.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		errorHandler(rw, http.StatusInternalServerError)
		return
	}

	// La variable search renvoie la première valeur du composant nommé de la requête. 
	// Les paramètres de corps POST et PUT ont priorité sur les valeurs de chaîne de requête d'URL.
	search := r.FormValue("search-choice")
	InfoArtists.SearchArtists(search)
	InfoArtists.SearchArt = search
	err = tmpl.Execute(rw, InfoArtists)
	if err != nil {
		errorHandler(rw, http.StatusInternalServerError)
		return
	}
}

//FilterPage - initialise les filtres en foctions des informations des artistes telque la date de création, le numéro de membre, date des first Album
func FilterPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer recoverHandle(rw)
	InfoArtists.Filter = nil
	path := filepath.Join("public", "html", "filter.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		log.Println(err)
		errorHandler(rw, http.StatusInternalServerError)
		return
	}

	filter := &Filters{}
	filter.CreationDateFrom = r.FormValue("Creation-date-from") 	     // filtre de la date de création de la première année
	filter.CreationDateTo = r.FormValue("Creation-date-to")    		    // filtre de la date de création de la dernière année
	filter.FirstAlbumFrom = r.FormValue("First-album-from")   		   // filtre du premier album de 
	filter.FirstAlbumTo = r.FormValue("First-album-to")      		  // filtre du premier album à
	filter.NumOfMembersFrom = r.FormValue("Number-of-members-from")  // filtre du numéro de membres de 
	filter.NumOfMembersTo = r.FormValue("Number-of-members-to")     // filtre du numéro de membres de

	// on retourne sur la page /Filtre, le resultat du filtre éffectué
	fmt.Println("\n***********************************")
	fmt.Printf("%v - %v\t| Creation date\n", filter.CreationDateFrom, filter.CreationDateTo)
	fmt.Println("-----------------------------------")
	fmt.Printf("%v - %v\t| First album\n", filter.FirstAlbumFrom, filter.FirstAlbumTo)
	fmt.Println("-----------------------------------")
	fmt.Printf("%v - %v\t\t| Number of members\n", filter.NumOfMembersFrom, filter.NumOfMembersTo)

	for _, value := range r.Form["location"] { // on initialise le filtre en fonction de la localisation
		filter.Locations = append(filter.Locations, value)
	}

	// on retourne sur la page /Filtre le resultat du filtre localisation
	fmt.Println(filter.Locations)

	// Et si erreur il me retourne NOT FOUND
	err = InfoArtists.FilterArtist(filter)
	if err != nil {
		log.Println(err)
		errorHandler(rw, http.StatusBadRequest)
		return
	}

	err = tmpl.Execute(rw, InfoArtists)
	if err != nil {
		log.Println(err)
		errorHandler(rw, http.StatusInternalServerError)
		return
	}
}

// La page InfoPage nous retourne les Instructions données et les objectifs attendus pour notre travail 
func InfoPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer recoverHandle(rw)
	path := filepath.Join("public", "html", "info.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		log.Println(err)
		errorHandler(rw, http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(rw, InfoArtists)
	if err != nil {
		log.Println(err)
		errorHandler(rw, http.StatusInternalServerError)
		return
	}
}

// on recupère notre handle
func recoverHandle(rw http.ResponseWriter) {
	if err := recover(); err != nil {
		errorHandler(rw, http.StatusInternalServerError)
		return
	}
}

// on renvoie les erreurs du handle
func errorHandler(rw http.ResponseWriter, status int) {
	http.Error(rw, http.StatusText(status), status)
	return
}
