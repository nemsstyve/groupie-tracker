package outils

type AllInfo struct {
	Art       []Artists
	Rel       Relation
	Search    []Artists
	Filter    []Artists
	SearchArt string
}

type Filters struct {
	CreationDateFrom string
	CreationDateTo   string
	FirstAlbumFrom   string
	FirstAlbumTo     string
	NumOfMembersFrom string
	NumOfMembersTo   string
	Locations        []string
}

type IsFilter struct {
	CreatDate bool
	FrstAlbum bool
	NumOfmem  bool
	Location  bool
}

type Artists struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Relation struct {
	Index []struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

type OneArtistOrBand struct {
	Art Artists
	Rel interface{}
}
