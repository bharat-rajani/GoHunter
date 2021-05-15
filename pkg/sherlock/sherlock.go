package sherlock

import (
	"fmt"
	"net/http"
)

type SherLock struct {
	Usernames []string
	Client *http.Client
}

func (sl *SherLock) Run() bool {

	for _, username := range sl.Usernames {
		// TODO: Complete the package
		fmt.Println(username)
	}
	return true
}
