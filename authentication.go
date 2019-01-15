package tmdb

import (
	"fmt"
)

// RequestToken type is a struct for request token JSON response.
type RequestToken struct {
	Success      bool   `json:"success"`
	ExpiresAt    string `json:"expires_at"`
	RequestToken string `json:"request_token"`
}

// AccessToken type is a struct for access token JSON request.
type AccessToken struct {
	AccessToken string `json:"access_token"`
}

// Session type is a struct for session JSON response.
type Session struct {
	Success   bool   `json:"success"`
	SessionID string `json:"session_id"`
}

// SessionWithLogin type is a struct for session with login JSON response.
type SessionWithLogin struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	RequestToken string `json:"request_token"`
}

// CreateRequestToken creates a temporary request token
// that can be used to validate a TMDb user login.
//
// Query String: api_key.
//
// https://developers.themoviedb.org/3/authentication/create-request-token
func (c *Client) CreateRequestToken() (*RequestToken, error) {
	tmdbURL := fmt.Sprintf("%s%stoken/new?api_key=%s", baseURL, authenticationURL, c.APIKey)
	a := RequestToken{}
	err := c.get(tmdbURL, &a)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

// CreateSession creates a new session id.
//
// You can use this method to create a fully valid session ID
// once a user has validated the request token.
//
// Query String: api_key.
//
// https://developers.themoviedb.org/3/authentication/create-session
func (c *Client) CreateSession(rt string) (*Session, error) {
	tmdbURL := fmt.Sprintf("%s%ssession/new?api_key=%s", baseURL, authenticationURL, c.APIKey)
	requestToken := RequestToken{RequestToken: rt}
	a := Session{}
	err := c.post(tmdbURL, &requestToken, &a)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

// CreateSessionWithLogin creates a new session id using login.
//
// This method allows an application to validate a request token
// by entering a username and password. Not all applications have
// access to a web view so this can be used as a substitute.
// If you decide to use this method please use HTTPS.
//
// Query String: api_key.
//
// https://developers.themoviedb.org/3/authentication/validate-request-token
// func (c *Client) CreateSessionWithLogin(u, p, rt string) (*RequestToken, error) {
// 	tmdbURL := fmt.Sprintf("%s%stoken/validate_with_login?api_key=%s", baseURL, authenticationURL, c.APIKey)
// 	loginSession := SessionWithLogin{
// 		Username:     u,
// 		Password:     p,
// 		RequestToken: rt,
// 	}
// 	a := RequestToken{}
// 	err := c.post(tmdbURL, &loginSession, &a)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &a, nil
// }

// CreateSessionFromV4 creates a new session id.
//
// Use this method to create a v3 session ID if you already have
// a valid v4 access token. The v4 token needs to be authenticated by the user.
// Your standard "read token" will not validate to create a session ID.
//
// Query String: api_key.
//
// https://developers.themoviedb.org/3/authentication/create-session-from-v4-access-token
// func (c *Client) CreateSessionFromV4(at string) (*Session, error) {
// 	tmdbURL := fmt.Sprintf("%s%ssession/convert/4?api_key=%s", baseURL, authenticationURL, c.APIKey)
// 	accessToken := AccessToken{AccessToken: at}
// 	a := Session{}
// 	err := c.post(tmdbURL, &accessToken, &a)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &a, nil
// }
