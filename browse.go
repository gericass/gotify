package gotify

import (
	"encoding/json"

	"github.com/gericass/gotify/extensions"
	"github.com/gericass/gotify/models"
)

// GetBrowseFeaturedPlaylists : the method for GET https://api.spotify.com/v1/browse/featured-playlists
func (t *Tokens) GetBrowseFeaturedPlaylists() (*models.BrowseFeaturedPlaylists, error) {
	/**
	https://developer.spotify.com/web-api/get-list-featured-playlists/
	*/

	endpoint := "https://api.spotify.com/v1/browse/featured-playlists"

	res, err := extensions.Request(endpoint, t.AccessToken)
	if err != nil {
		return nil, err
	}

	browseFeaturedPlaylists := new(models.BrowseFeaturedPlaylists)

	err = json.Unmarshal(res, browseFeaturedPlaylists)
	if err != nil {
		return nil, err
	}
	return browseFeaturedPlaylists, nil
}

// GetBrowseNewReleases : the method for GET https://api.spotify.com/v1/browse/new-releases
func (t *Tokens) GetBrowseNewReleases() (*models.BrowseNewReleases, error) {
	/**
	https://developer.spotify.com/web-api/get-list-new-releases/
	*/

	endpoint := "https://api.spotify.com/v1/browse/new-releases"

	res, err := extensions.Request(endpoint, t.AccessToken)
	if err != nil {
		return nil, err
	}

	browseNewReleases := new(models.BrowseNewReleases)

	err = json.Unmarshal(res, browseNewReleases)
	if err != nil {
		return nil, err
	}
	return browseNewReleases, nil
}

// GetBrowseCategories : the method for GET https://api.spotify.com/v1/browse/categories
func (t *Tokens) GetBrowseCategories() (*models.BrowseCategories, error) {
	/**
	https://developer.spotify.com/web-api/get-list-new-releases/
	*/

	endpoint := "https://api.spotify.com/v1/browse/categories"

	res, err := extensions.Request(endpoint, t.AccessToken)
	if err != nil {
		return nil, err
	}

	browseCategories := new(models.BrowseCategories)

	err = json.Unmarshal(res, browseCategories)
	if err != nil {
		return nil, err
	}
	return browseCategories, nil
}

// GetBrowseCategory : the method for GET https://api.spotify.com/v1/browse/categories/{category_id}
func (t *Tokens) GetBrowseCategory(categoryID string) (*models.BrowseCategory, error) {
	/**
	https://developer.spotify.com/web-api/get-list-new-releases/
	*/

	endpoint := "https://api.spotify.com/v1/browse/categories/" + categoryID

	res, err := extensions.Request(endpoint, t.AccessToken)
	if err != nil {
		return nil, err
	}

	browseCategory := new(models.BrowseCategory)

	err = json.Unmarshal(res, browseCategory)
	if err != nil {
		return nil, err
	}
	return browseCategory, nil
}
