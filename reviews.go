package tmdb

import "fmt"

// ReviewDetails type is a struct for details JSON response.
type ReviewDetails struct {
	ID         string `json:"id"`
	Author     string `json:"author"`
	Content    string `json:"content"`
	Iso639_1   string `json:"iso_639_1"`
	MediaID    int64  `json:"media_id"`
	MediaTitle string `json:"media_title"`
	MediaType  string `json:"media_type"`
	URL        string `json:"url"`
}

// GetReviewDetails get review details by id.
//
// https://developers.themoviedb.org/3/movies/get-movie-details
func (c *Client) GetReviewDetails(
	id string,
) (*ReviewDetails, error) {
	tmdbURL := fmt.Sprintf(
		"%s/review/%s?api_key=%s",
		baseURL, id, c.apiKey,
	)
	t := ReviewDetails{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
