package outils

import (
	"strconv"
	"strings"
)

// on initialise le début des différentes recherches pour chaque artiste
func (info *AllInfo) SearchArtists(search string) {
	date, err := strconv.Atoi(search)
	if err == nil {
		for _, v := range info.Art {
			if date == v.CreationDate {
				info.Search = append(info.Search, v)
				continue
			}
		}
		return
	}

	for _, v := range info.Art {
		if strings.Contains(v.Name, search) {  // on pose la condition de la recherche du nom en lui disant si le nom recherché est présent dans notre struct
			if info.isDuplicate(v.ID) {       // on lui dit même si l'info est dupliqué avec l'id, il continue la recherche
				continue
			}
			info.Search = append(info.Search, v)
		}
		if strings.Contains(v.FirstAlbum, search) {   // on pose la condition de la recherche du FirstAlbum en lui disant si la date de l'album recherchée est présente dans notre struct
			if info.isDuplicate(v.ID) {              // on lui dit même si l'info est dupliqué avec l'id, il continue la recherche
				continue
			}
			info.Search = append(info.Search, v)
		}
		for _, v1 := range v.Members {
			if strings.Contains(v1, search) {
				if info.isDuplicate(v.ID) {
					continue
				}
				info.Search = append(info.Search, v)
				break
			}
		}
		if info.isDuplicate(v.ID) {
			continue
		}
		for i, datelocation := range info.Rel.Index {
			if info.isDuplicate(datelocation.ID) {
				break
			}
			for location := range datelocation.DatesLocations {
				if strings.Contains(location, search) {
					info.Search = append(info.Search, info.Art[i])
					break
				}
			}
		}
	}
}

// cette fonction initialise toutes les recherches selon l' ID même si celle-ci sont dupliquées 
func (info *AllInfo) isDuplicate(id int) bool {
	for _, v := range info.Search {
		if v.ID == id {
			return true
		}
	}
	return false
}
