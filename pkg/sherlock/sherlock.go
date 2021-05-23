package sherlock

import (
	"errors"
	"fmt"
	"net/http"
	urlpkg "net/url"
	"regexp"
	"strings"
	"sync"
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

type SherlockRequests struct {
	Req chan *http.Request
	Err chan error
}

func NewSherlockRequests() *SherlockRequests {
	return &SherlockRequests{
		Req: make(chan *http.Request, 20),
		Err: make(chan error, 20),
	}
}

func (sl *SherLock) Run() (bool, error) {

	sites, err := NewSites(urlpkg.URL{}, sl.Client)
	if err != nil {
		return false, err
	}

	var wg sync.WaitGroup
	for _, username := range sl.Usernames {
		wg.Add(1)
		go sherLockIt(&wg,sl.Client,sites,username)
	}
	wg.Wait()
	return false, err
}

func sherLockIt(wg *sync.WaitGroup, client *http.Client, sites Sites, username string){
	defer wg.Done()
	for site, data := range sites {
		if req := prepareRequest(username, site, data); req != nil {
			wg.Add(1)
			go doRequest(wg, req, client)
		}
	}
}

func doRequest(wg *sync.WaitGroup, req *http.Request, client *http.Client) {
	defer wg.Done()
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(fmt.Errorf("request halted %v", err))
	} else if response.StatusCode == http.StatusOK {
		fmt.Printf("-------->%v\n", response.Status)
	}
}

func prepareRequest(username, site string, data SiteInformation) *http.Request {

	//defer wg.Done()
	var req http.Request

	if pattern, found := data.Information["regexCheck"]; found {
		matched, err := regexp.MatchString(pattern.(string), username)
		if err != nil {
			//reqChan.Req <- nil
			//reqChan.Err <- err
			fmt.Println(err.Error())
			return nil
		} else if !matched {
			//reqChan.Req <- nil
			//reqChan.Err <- errors.New("regex matching failed, aborting request creation")
			fmt.Println(pattern,matched,errors.New("regex matching failed, aborting request creation"))
			return nil
		}
	}

	url := data.UrlUsernameFormat
	url = strings.Replace(url, "{}", username, 1)
	if urlProbe, found := data.Information["urlProbe"]; found {
		url = strings.Replace(urlProbe.(string), "{}", username, 1)
		fmt.Println(urlProbe,url)
	}

	u, err := urlpkg.Parse(url)
	if err != nil {
		//reqChan.Req <- nil
		//reqChan.Err <- err
		return nil
	}

	req.URL = u
	req.Method = http.MethodGet
	switch data.Information["errorType"].(string) {
	case "status_code":
		if isHeadReq := data.Information["request_head_only"]; isHeadReq != nil && isHeadReq.(bool) {
			req.Method = http.MethodHead
		}
		//case "response_url":
		//	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		//		return http.ErrUseLastResponse
		//	}
	}

	return &req
}
