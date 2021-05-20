package sherlock

import (
	"errors"
	"fmt"
	"net/http"
	urlpkg "net/url"
	"regexp"
	"strings"
)

type SherLock struct {
	Usernames []string
	Client    *http.Client
}

func NewSherLock(usernames []string, client *http.Client) *SherLock {
	return &SherLock{
		Usernames: usernames,
		Client:    client,
	}
}

func (sl *SherLock) Run() (bool, error) {

	sites, err := NewSites(urlpkg.URL{}, sl.Client)
	if err != nil {
		return false, err
	}

	// TODO: Add goroutine magic
	for _, username := range sl.Usernames {
		// TODO: Complete the package
		fmt.Println(username)
		for site, data := range sites {
			req, err := prepareRequest(sl.Client, username, site, data)
			if err != nil {
				fmt.Println(fmt.Errorf("request halted %v", err))
			} else {
				response, err := sl.Client.Do(req)
				if err != nil {
					fmt.Println(fmt.Errorf("request halted %v", err))
				}
				fmt.Print(response)
			}
		}
	}

	return true, nil
}

func prepareRequest(client *http.Client, username, site string, data SiteInformation) (*http.Request, error) {
	var req http.Request
	//if err!=nil{
	//	return nil,err
	//}

	if pattern, found := data.Information["regexCheck"]; found {
		matched, err := regexp.MatchString(pattern.(string), username)
		if err != nil {
			return nil, err
		} else if !matched {
			return nil, errors.New("regex matching failed, aborting request creation")
		}
	}

	url := data.UrlUsernameFormat
	if urlProbe, found := data.Information["urlProbe"]; found {
		url = strings.Replace(urlProbe.(string), "{}", username, 1)
	}

	u, err := urlpkg.Parse(url)
	if err != nil {
		return nil, err
	}

	req.URL = u
	req.Method = http.MethodGet
	switch data.Information["errorType"].(string) {
	case "status_code":
		if isHeadReq := data.Information["request_head_only"]; isHeadReq != nil && isHeadReq.(bool) {
			req.Method = http.MethodHead
		}
	case "response_url":
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	return &req, nil
}
