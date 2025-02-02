package tmdb

import (
	"fmt"
)

// TVEpisodeDetails type is a struct for details JSON response.
type TVEpisodeDetails struct {
	AirDate string `json:"air_date"`
	Crew    []struct {
		ID          int64  `json:"id"`
		CreditID    string `json:"credit_id"`
		Name        string `json:"name"`
		Department  string `json:"department"`
		Job         string `json:"job"`
		Gender      int    `json:"gender"`
		ProfilePath string `json:"profile_path"`
	} `json:"crew"`
	EpisodeNumber int `json:"episode_number"`
	GuestStars    []struct {
		ID          int64  `json:"id"`
		Name        string `json:"name"`
		CreditID    string `json:"credit_id"`
		Character   string `json:"character"`
		Order       int    `json:"order"`
		Gender      int    `json:"gender"`
		ProfilePath string `json:"profile_path"`
	} `json:"guest_stars"`
	Name           string  `json:"name"`
	Overview       string  `json:"overview"`
	ID             int64   `json:"id"`
	ProductionCode string  `json:"production_code"`
	SeasonNumber   int     `json:"season_number"`
	StillPath      string  `json:"still_path"`
	VoteAverage    float32 `json:"vote_average"`
	VoteCount      int64   `json:"vote_count"`
	*TVEpisodeCreditsAppend
	*TVEpisodeExternalIDsAppend
	*TVEpisodeImagesAppend
	*TVEpisodeTranslationsAppend
	*TVEpisodeVideosAppend
}

// TVEpisodeChanges type is a struct for changes JSON response.
type TVEpisodeChanges struct {
	Changes []struct {
		Key   string `json:"key"`
		Items []struct {
			ID     string `json:"id"`
			Action string `json:"action"`
			Time   string `json:"time"`
			Value  string `json:"value"`
		} `json:"items"`
	} `json:"changes"`
}

// TVEpisodeCredits type is a struct for credits JSON response.
type TVEpisodeCredits struct {
	Cast []struct {
		Character   string `json:"character"`
		CreditID    string `json:"credit_id"`
		Gender      int    `json:"gender"`
		ID          int64  `json:"id"`
		Name        string `json:"name"`
		Order       int    `json:"order"`
		ProfilePath string `json:"profile_path"`
	} `json:"cast"`
	Crew []struct {
		ID          int64  `json:"id"`
		CreditID    string `json:"credit_id"`
		Name        string `json:"name"`
		Department  string `json:"department"`
		Job         string `json:"job"`
		Gender      int    `json:"gender"`
		ProfilePath string `json:"profile_path"`
	} `json:"crew"`
	GuestStars []struct {
		ID          int64  `json:"id"`
		Name        string `json:"name"`
		CreditID    string `json:"credit_id"`
		Character   string `json:"character"`
		Order       int    `json:"order"`
		Gender      int    `json:"gender"`
		ProfilePath string `json:"profile_path"`
	} `json:"guest_stars"`
	ID int64 `json:"id,omitempty"`
}

// TVEpisodeCreditsAppend type is a struct
// for credits in append to response.
type TVEpisodeCreditsAppend struct {
	Credits *TVEpisodeCredits `json:"credits,omitempty"`
}

// TVEpisodeExternalIDs type is a struct for external ids JSON response.
type TVEpisodeExternalIDs struct {
	ID          int64  `json:"id,omitempty"`
	IMDbID      string `json:"imdb_id"`
	FreebaseMID string `json:"freebase_mid"`
	FreebaseID  string `json:"freebase_id"`
	TVDBID      int64  `json:"tvdb_id"`
	TVRageID    int64  `json:"tvrage_id"`
}

// TVEpisodeExternalIDsAppend type is a struct
// for external ids in append to response.
type TVEpisodeExternalIDsAppend struct {
	ExternalIDs *TVEpisodeExternalIDs `json:"external_ids,omitempty"`
}

// TVEpisodeImages type is a struct for images JSON response.
type TVEpisodeImages struct {
	ID     int64 `json:"id,omitempty"`
	Stills []struct {
		AspectRatio float32     `json:"aspect_ratio"`
		FilePath    string      `json:"file_path"`
		Height      int         `json:"height"`
		Iso6391     interface{} `json:"iso_639_1"`
		VoteAverage float32     `json:"vote_average"`
		VoteCount   int64       `json:"vote_count"`
		Width       int         `json:"width"`
	} `json:"stills"`
}

// TVEpisodeImagesAppend type is a struct
// for images in append to response.
type TVEpisodeImagesAppend struct {
	Images *TVEpisodeImages `json:"images,omitempty"`
}

// TVEpisodeTranslations type is a struct for translations JSON response.
type TVEpisodeTranslations struct {
	ID           int64 `json:"id,omitempty"`
	Translations []struct {
		Iso3166_1   string `json:"iso_3166_1"`
		Iso639_1    string `json:"iso_639_1"`
		Name        string `json:"name"`
		EnglishName string `json:"english_name"`
		Data        struct {
			Name     string `json:"name"`
			Overview string `json:"overview"`
		} `json:"data"`
	} `json:"translations"`
}

// TVEpisodeTranslationsAppend type is a struct
// for translations in append to response.
type TVEpisodeTranslationsAppend struct {
	Translations *TVEpisodeTranslations `json:"translations,omitempty"`
}

// TVEpisodeRate type is a struct for rate JSON response.
type TVEpisodeRate struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}

// TVEpisodeVideos type is a struct for videos JSON response.
type TVEpisodeVideos struct {
	ID      int64 `json:"id,omitempty"`
	Results []struct {
		ID        string `json:"id"`
		Iso639_1  string `json:"iso_639_1"`
		Iso3166_1 string `json:"iso_3166_1"`
		Key       string `json:"key"`
		Name      string `json:"name"`
		Site      string `json:"site"`
		Size      int    `json:"size"`
		Type      string `json:"type"`
	} `json:"results"`
}

// TVEpisodeVideosAppend type is a struct
// for videos in append to response.
type TVEpisodeVideosAppend struct {
	Videos *TVEpisodeVideos `json:"videos,omitempty"`
}

// GetTVEpisodeDetails get the TV episode details by id.
//
// Supports append_to_response.
//
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-details
func (c *Client) GetTVEpisodeDetails(
	id, s, e int, o map[string]string,
) (*TVEpisodeDetails, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d%s%d?api_key=%s%s",
		baseURL, tvURL, id, tvSeasonURL, s, tvEpisodeURL, e, c.apiKey, options,
	)
	t := TVEpisodeDetails{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVEpisodeChanges get the changes for a TV episode.
// By default only the last 24 hours are returned.
//
// You can query up to 14 days in a single query by using
// the start_date and end_date query parameters.
//
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-changes
func (c *Client) GetTVEpisodeChanges(
	id int, o map[string]string,
) (*TVEpisodeChanges, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%sepisode/%d/changes?api_key=%s%s",
		baseURL, tvURL, id, c.apiKey, options,
	)
	t := TVEpisodeChanges{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVEpisodeCredits get the credits (cast, crew and guest stars) for a TV episode.
//
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-credits
func (c *Client) GetTVEpisodeCredits(
	id, s, e int,
) (*TVEpisodeCredits, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d%s%d/credits?api_key=%s",
		baseURL, tvURL, id, tvSeasonURL, s, tvEpisodeURL, e, c.apiKey,
	)
	t := TVEpisodeCredits{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVEpisodeExternalIDs get the external ids for a TV episode.
// We currently support the following external sources.
//
// Media Databases: IMDb ID, TVDB ID, Freebase MID*, Freebase ID* TVRage ID*.
//
// *Defunct or no longer available as a service.
//
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-external-ids
func (c *Client) GetTVEpisodeExternalIDs(
	id, s, e int,
) (*TVEpisodeExternalIDs, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d%s%d/external_ids?api_key=%s",
		baseURL, tvURL, id, tvSeasonURL, s, tvEpisodeURL, e, c.apiKey,
	)
	t := TVEpisodeExternalIDs{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVEpisodeImages get the images that belong to a TV episode.
//
// Querying images with a language parameter will filter the results.
// If you want to include a fallback language (especially useful for backdrops)
// you can use the include_image_language parameter.
// This should be a comma separated value like so: include_image_language=en,null.
//
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-images
func (c *Client) GetTVEpisodeImages(
	id, s, e int,
) (*TVEpisodeImages, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d%s%d/images?api_key=%s",
		baseURL, tvURL, id, tvSeasonURL, s, tvEpisodeURL, e, c.apiKey,
	)
	t := TVEpisodeImages{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVEpisodeTranslations get the translation data for an episode.
//
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-translations
func (c *Client) GetTVEpisodeTranslations(
	id, s, e int,
) (*TVEpisodeTranslations, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d%s%d/translations?api_key=%s",
		baseURL, tvURL, id, tvSeasonURL, s, tvEpisodeURL, e, c.apiKey,
	)
	t := TVEpisodeTranslations{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVEpisodeVideos get the videos that have been added to a TV episode.
//
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-videos
func (c *Client) GetTVEpisodeVideos(
	id, s, e int, o map[string]string,
) (*TVEpisodeVideos, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d%s%d/videos?api_key=%s%s",
		baseURL, tvURL, id, tvSeasonURL, s, tvEpisodeURL, e, c.apiKey, options,
	)
	t := TVEpisodeVideos{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// TODO: Account States, Rate TV Episode, Delete Rating.
