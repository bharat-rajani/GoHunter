package factory

import (
	"github.com/bharat-rajani/GoHunter/internal/client"
	"github.com/bharat-rajani/GoHunter/internal/config"
	"github.com/bharat-rajani/GoHunter/pkg/iostreams"
	"net/http"
	"os"
)

func NewHttpClient(io *iostreams.IOStreams, configuration *config.Configuration) *http.Client {

	var opts []client.ClientOption
	if verbose := os.Getenv("DEBUG"); verbose != "" {
		logTraffic := true
		opts = append(opts, client.VerboseLog(io.ErrOut, logTraffic, io.IsStderrTTY()))
	}

	// create opts based on configuration options
	//opts = append(opts,
	//	//client.AddHeader("User-Agent", fmt.Sprintf("GitHub CLI %s", appVersion)),
	//)

	return client.NewHTTPClient(opts...)
}
