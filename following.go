package gotify

import (
	"encoding/json"
	"fmt"

	"github.com/gericass/gotify/extensions"
	"github.com/gericass/gotify/models"
)

// GetFollowingArtists : the method for GET https://api.spotify.com/v1/me/following?type=artist
func (t *Tokens) GetFollowingArtists() (*models.FollowingArtists, error) {
	/**
	https://developer.spotify.com/web-api/get-followed-artists/
	*/

	endpoint := "https://api.spotify.com/v1/me/following?type=artist"

	res, err := extensions.GetRequest(endpoint, t.AccessToken)
	if err != nil {
		return nil, err
	}

	followingArtists := new(models.FollowingArtists)

	err = json.Unmarshal(res, followingArtists)
	if err != nil {
		return nil, err
	}
	return followingArtists, nil
}

// FollowArtistsOrUsers : the method for PUT https://api.spotify.com/v1/me/following
func (t *Tokens) FollowArtistsOrUsers(followType string, IDs []string) error {
	/**
	https://developer.spotify.com/web-api/follow-artists-users/
	*/

	endpoint := "https://api.spotify.com/v1/me/following?type=" + followType + "&ids="

	for i, v := range IDs {
		if i == 0 {
			endpoint += v
		} else {
			endpoint += "," + v
		}
	}

	res, err := extensions.PutRequest(endpoint, t.AccessToken)
	if err != nil {
		return err
	}
	if res != 204 {
		return fmt.Errorf("%d", 204)
	}
	return nil
}

// UnfollowArtistsOrUsers : the method for DELETE https://api.spotify.com/v1/me/following
func (t *Tokens) UnfollowArtistsOrUsers(unfollowType string, IDs []string) error {
	/**
	https://developer.spotify.com/web-api/follow-artists-users/
	*/

	endpoint := "https://api.spotify.com/v1/me/following?type=" + unfollowType + "&ids="

	for i, v := range IDs {
		if i == 0 {
			endpoint += v
		} else {
			endpoint += "," + v
		}
	}

	res, err := extensions.DeleteRequest(endpoint, t.AccessToken)
	if err != nil {
		return err
	}
	if res != 204 {
		return fmt.Errorf("%d", 204)
	}
	return nil
}

// CurrentFollowsArtistsOrUsers : the method for GET https://api.spotify.com/v1/me/following/contains
func (t *Tokens) CurrentFollowsArtistsOrUsers(followType string, IDs []string) (*models.CurrentFollowsArtistsOrUsers, error) {
	/**
	https://developer.spotify.com/web-api/check-current-user-follows/
	*/

	endpoint := "https://api.spotify.com/v1/me/following/contains?type=" + followType + "&ids="

	for i, v := range IDs {
		if i == 0 {
			endpoint += v
		} else {
			endpoint += "," + v
		}
	}

	res, err := extensions.GetRequest(endpoint, t.AccessToken)
	if err != nil {
		return nil, err
	}
	currentFollowArtistsOrUsers := new(models.CurrentFollowsArtistsOrUsers)

	err = json.Unmarshal(res, currentFollowArtistsOrUsers)
	if err != nil {
		return nil, err
	}
	return currentFollowArtistsOrUsers, nil
}

// FollowPlaylist : the method for PUT https://api.spotify.com/v1/users/{owner_id}/playlists/{playlist_id}/followers
func (t *Tokens) FollowPlaylist(userID string, playlistID string) error {
	/**
	https://developer.spotify.com/web-api/follow-playlist/
	*/

	endpoint := "https://api.spotify.com/v1/users/" + userID + "/playlists/" + playlistID + "/followers"
	res, err := extensions.PutRequest(endpoint, t.AccessToken)
	if err != nil {
		return err
	}
	if res != 200 {
		return fmt.Errorf("%d", res)
	}
	return nil
}
