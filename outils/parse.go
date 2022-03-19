package outils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var InfoArtists AllInfo

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

func Recover(url string, InfoArtists interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, InfoArtists)
}
