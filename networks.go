package tmdb

import "fmt"

// NetworkDetails type is a struct for details JSON response.
type NetworkDetails struct {
	Headquarters  string `json:"headquarters"`
	Homepage      string `json:"homepage"`
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	OriginCountry string `json:"origin_country"`
}

// NetworkAlternativeNames type is a struct for alternative names JSON response.
type NetworkAlternativeNames struct {
	ID      int64 `json:"id"`
	Results []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"results"`
}

// NetworkImages type is a struct for images JSON response.
type NetworkImages struct {
	ID    int64 `json:"id"`
	Logos []struct {
		AspectRatio float64 `json:"aspect_ratio"`
		FilePath    string  `json:"file_path"`
		Height      int     `json:"height"`
		ID          string  `json:"id"`
		FileType    string  `json:"file_type"`
		VoteAverage float64 `json:"vote_average"`
		VoteCount   int64   `json:"vote_count"`
		Width       int     `json:"width"`
	} `json:"logos"`
}

// GetNetworkDetails get the details of a network.
//
// https://developers.themoviedb.org/3/networks/get-network-details
func (c *Client) GetNetworkDetails(id int) (*NetworkDetails, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d?api_key=%s",
		baseURL, networkURL, id, c.apiKey,
	)
	n := NetworkDetails{}
	err := c.get(tmdbURL, &n)
	if err != nil {
		return nil, err
	}
	return &n, nil
}

// GetNetworkAlternativeNames get the alternative names of a network.
//
// https://developers.themoviedb.org/3/networks/get-network-alternative-names
func (c *Client) GetNetworkAlternativeNames(
	id int,
) (*NetworkAlternativeNames, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d/alternative_names?api_key=%s",
		baseURL, networkURL, id, c.apiKey,
	)
	n := NetworkAlternativeNames{}
	err := c.get(tmdbURL, &n)
	if err != nil {
		return nil, err
	}
	return &n, nil
}

// GetNetworkImages get the TV network logos by id.
//
// There are two image formats that are supported for networks,
// PNG's and SVG's. You can see which type the original file is
// by looking at the file_type field. We prefer SVG's as they are
// resolution independent and as such, the width and height are
// only there to reflect the original asset that was uploaded.
// An SVG can be scaled properly beyond those dimensions if you
// call them as a PNG.
//
//
//
// https://developers.themoviedb.org/3/networks/get-network-images
func (c *Client) GetNetworkImages(id int) (*NetworkImages, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d/images?api_key=%s",
		baseURL, networkURL, id, c.apiKey,
	)
	n := NetworkImages{}
	err := c.get(tmdbURL, &n)
	if err != nil {
		return nil, err
	}
	return &n, nil
}
