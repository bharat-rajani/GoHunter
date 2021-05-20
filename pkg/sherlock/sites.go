package sherlock

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type SiteInformation struct {
	Name              string                 `json:"name"`
	UrlHome           string                 `json:"urlMain"`
	UrlUsernameFormat string                 `json:"url"`
	UsernameClaimed   string                 `json:"username_claimed"`
	UsernameUnclaimed string                 `json:"username_unclaimed"`
	// Information contains json object in order to retain unknowns.
	Information       map[string]interface{} `json:"-"`
}

func NewSiteInformation(siteName string, data map[string]interface{}) (*SiteInformation){

	siteInfo :=  &SiteInformation{
		Name: siteName,
		UrlHome: data["urlMain"].(string),
		UrlUsernameFormat: data["url"].(string),
		UsernameClaimed: data["username_claimed"].(string),
		UsernameUnclaimed: data["username_unclaimed"].(string),
	}

	delete(data,"urlMain")
	delete(data,"url")
	delete(data,"username_claimed")
	delete(data,"username_unclaimed")

	siteInfo.Information = data
	return siteInfo
}


type Sites map[string]SiteInformation

func NewSites(urlPath url.URL, client *http.Client) (Sites, error) {

	dataFilePath := "https://raw.githubusercontent.com/sherlock-project/sherlock/master/sherlock/resources/data.json"
	sites := make(Sites)
	switch urlPath {
	case url.URL{}:
		resp, err := client.Get(dataFilePath)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		var siteData map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&siteData)
		if err != nil {
			return nil, err
		}
		for key, val := range siteData{
			sites[key] = *NewSiteInformation(key, val.(map[string]interface{}))
		}

		return sites, nil
	}
	return nil, nil
}
