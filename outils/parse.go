package outils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var InfoArtists AllInfo

// on lance l'analyse de nos API (artists et relation)
func Parse() error { 
	const (
		artists  string = "https://groupietrackers.herokuapp.com/api/artists"
		relation string = "https://groupietrackers.herokuapp.com/api/relation"
	)
	err := Recover(artists, &InfoArtists.Art)
	if err != nil {
		return err
	}
	err = Recover(relation, &InfoArtists.Rel)
	if err != nil {
		return err
	}
	return nil
}

// cette fonction nous permet de récupérer les informations de chaque artistes
func Recover(url string, InfoArtists interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close() // là on reporte la reponse de notre corps
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, InfoArtists)  //Unmarshal analyse les données encodées en JSON et stocke le résultat dans la valeur pointée par v. Si v est nil ou n'est pas un pointeur, Unmarshal renvoie une InvalidUnmarshalError.
	                                         // Unmarshal utilise l'inverse des encodages utilisés par Marshal, allouant des cartes, des tranches et des pointeurs selon les besoins, avec les règles supplémentaires.
}
