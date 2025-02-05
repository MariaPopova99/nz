package nz

import "fmt"

type caracterResponse struct {
	Data []AssetData `json:"data"`
}

type AssetData struct {
	ID              int      `json:"_id"`
	Films           []string `json:"films"`
	ShortFilms      []string `json:"shortFilms"`
	TVShows         []string `json:"tvShows"`
	VideoGames      []string `json:"videoGames"`
	ParkAttractions []string `json:"parkAttractions"`
	Name            string   `json:"name"`
	ImageURL        string   `json:"imageUrl"`
}

func (d AssetData) GetInfo() string {
	var film string
	switch {
	case len(d.Films) > 0:
		film = d.Films[0]
	case len(d.ShortFilms) > 0:
		film = d.ShortFilms[0]
	case len(d.TVShows) > 0:
		film = d.TVShows[0]
	case len(d.VideoGames) > 0:
		film = d.VideoGames[0]
	case len(d.ParkAttractions) > 0:
		film = d.ParkAttractions[0]
	}

	return fmt.Sprintf("The character %s played at %s. Check the pic during the link %s",
		d.Name, film, d.ImageURL)
}
