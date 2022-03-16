package controller

import (
	"strconv"
	"strings"
)

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
		if strings.Contains(v.Name, search) {
			if info.isDuplicate(v.ID) {
				continue
			}
			info.Search = append(info.Search, v)
		}
		if strings.Contains(v.FirstAlbum, search) {
			if info.isDuplicate(v.ID) {
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

func (info *AllInfo) isDuplicate(id int) bool {
	for _, v := range info.Search {
		if v.ID == id {
			return true
		}
	}
	return false
}
