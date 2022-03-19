package outils

import (
	"log"
	"strconv"
)

//On initialise du Filtre
func (info *AllInfo) FilterArtist(f *Filters) error {
	var have bool
	var err error
	for _, v := range info.Art {
		have, err = f.filter(v)
		if err != nil {
			log.Println(err)
			return err
		}
		if have && err == nil {
			info.Filter = append(info.Filter, v)
		}
	}
	return nil
}

// cette fonction traite nos différents filtres selon le booléens et les erreurs
func (f *Filters) filter(artist Artists) (bool, error) {
	isFilter := IsFilter{}

	// on initie les deux variables de notre filtre creation de date
	from, err := strconv.Atoi(f.CreationDateFrom)
	if err != nil {
		log.Println(err)
		return false, err
	}
	to, err := strconv.Atoi(f.CreationDateTo)
	if err != nil {
		log.Println(err)
		return false, err
	}
	// là on pose la condition du filtre création de date et si cette condition n'est pas vrai alors il nous retourne not found
	if artist.CreationDate >= from && artist.CreationDate <= to {
		isFilter.CreatDate = true
	} else {
		return false, nil
	}

	// on initie les deux variables de notre filtre FirstAlbum
	from, err = strconv.Atoi(f.FirstAlbumFrom)
	if err != nil {
		log.Println(err)
		return false, err
	}
	to, err = strconv.Atoi(f.FirstAlbumTo)
	if err != nil {
		log.Println(err)
		return false, err
	}
	frstAlbum, err := strconv.Atoi(artist.FirstAlbum[6:]) // on convertit chaque album qui est une chaîne de caractère en int
	if err != nil {
		log.Println(err)
		return false, err
	}

	// là on pose la condition du filtre de l'Album
	if frstAlbum >= from && frstAlbum <= to {
		isFilter.FrstAlbum = true
	} else { // si cette condition n'est pas vrai alors il nous retourne not found
		return false, nil
	}

	// On vérifie par le nombre de membres
	from, err = strconv.Atoi(f.NumOfMembersFrom)
	if err != nil {
		log.Println(err)
		return false, err
	}
	to, err = strconv.Atoi(f.NumOfMembersTo)
	if err != nil {
		log.Println(err)
		return false, err
	}

	// là on pose la condition du filtre de l'Id du membre 
	members := len(artist.Members) // là, on definit une variable membre qui sera égale à la fonction intégrée len renvoie la longueur de v, selon son type :
	if members >= from && members <= to { 
		isFilter.NumOfmem = true // le filtre de la condition posée affiche si la conditon est respectée, sinon n'affiche rien
	} else {
		return false, nil
	}

	// là on pose la condition du filtre de localisation
	if len(f.Locations) == 0 {
		isFilter.Location = true
	} else {
		for _, v1 := range f.Locations {
			var checker bool
			for v2, _ := range InfoArtists.Rel.Index[artist.ID-1].DatesLocations {
				if v1 == v2 {
					checker = true
					break
				}
			}
			if !checker {
				return false, nil
			}
			isFilter.Location = true
		}
	}

	// là on définit la condition exact de notre filtre qui nous sera retourné sur noter localhost, sinon il ne nous retourne nul
	if isFilter.CreatDate == true && isFilter.FrstAlbum == true && isFilter.NumOfmem == true && isFilter.Location == true {
		return true, nil
	}
	return false, nil
}
