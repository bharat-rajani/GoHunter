package sherlock

import "net/url"

type SiteInformation struct {
}

type Sites struct {
	SiteInfoMap map[string]SiteInformation
}

func NewSites(urlPath url.URL) {

	if urlPath == (url.URL{}) {

	}
}
