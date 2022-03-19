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

func (f *Filters) filter(artist Artists) (bool, error) {
	isFilter := IsFilter{}

	// vérification par date de création
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
	if artist.CreationDate >= from && artist.CreationDate <= to {
		isFilter.CreatDate = true
	} else {
		return false, nil
	}

	// vérification du premier album
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
	frstAlbum, err := strconv.Atoi(artist.FirstAlbum[6:]) // on convertit la chaîne de caractère en int
	if err != nil {
		log.Println(err)
		return false, err
	}

	if frstAlbum >= from && frstAlbum <= to {
		isFilter.FrstAlbum = true
	} else {
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

	members := len(artist.Members)
	if members >= from && members <= to {
		isFilter.NumOfmem = true
	} else {
		return false, nil
	}

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

	if isFilter.CreatDate == true && isFilter.FrstAlbum == true && isFilter.NumOfmem == true && isFilter.Location == true {
		return true, nil
	}

	return false, nil
}
