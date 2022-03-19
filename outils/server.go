package outils

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// MainPage - page avec tous les artistes/groupes
func MainPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer recoverHandle(rw)

	path := filepath.Join("public", "html", "index.html")

	tmpl, err := template.ParseFiles(path)
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

	path := filepath.Join("public", "html", "artist.html")

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

func SearchPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer recoverHandle(rw)
	InfoArtists.Search = nil

	path := filepath.Join("public", "html", "search.html")

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		errorHandler(rw, http.StatusInternalServerError)
		return
	}

	search := r.FormValue("search-choice")
	InfoArtists.SearchArtists(search)
	InfoArtists.SearchArt = search

	err = tmpl.Execute(rw, InfoArtists)
	if err != nil {
		errorHandler(rw, http.StatusInternalServerError)
		return
	}
}

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
	filter.CreationDateFrom = r.FormValue("Creation-date-from")
	filter.CreationDateTo = r.FormValue("Creation-date-to")
	filter.FirstAlbumFrom = r.FormValue("First-album-from")
	filter.FirstAlbumTo = r.FormValue("First-album-to")
	filter.NumOfMembersFrom = r.FormValue("Number-of-members-from")
	filter.NumOfMembersTo = r.FormValue("Number-of-members-to")

	fmt.Println("\n***********************************")
	fmt.Printf("%v - %v\t| Creation date\n", filter.CreationDateFrom, filter.CreationDateTo)
	fmt.Println("-----------------------------------")
	fmt.Printf("%v - %v\t| First album\n", filter.FirstAlbumFrom, filter.FirstAlbumTo)
	fmt.Println("-----------------------------------")
	fmt.Printf("%v - %v\t\t| Number of members\n", filter.NumOfMembersFrom, filter.NumOfMembersTo)

	for _, value := range r.Form["location"] {
		filter.Locations = append(filter.Locations, value)
	}

	fmt.Println(filter.Locations)

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

func recoverHandle(rw http.ResponseWriter) {
	if err := recover(); err != nil {
		errorHandler(rw, http.StatusInternalServerError)
		return
	}
}

func errorHandler(rw http.ResponseWriter, status int) {
	http.Error(rw, http.StatusText(status), status)
	return
}
